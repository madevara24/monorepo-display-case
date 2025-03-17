const fs = require('fs');
const path = require('path');
const http = require('http');
const { google } = require('googleapis');
const { Command } = require('commander');
require('dotenv').config();

const program = new Command();

async function loadData(sheetName = process.env.GOOGLE_SHEET_NAME) {
    try {
        // Load credentials
        const credentials = JSON.parse(
            fs.readFileSync(process.env.GOOGLE_APPLICATION_CREDENTIALS, 'utf8')
        );

        // Create auth client
        const auth = new google.auth.GoogleAuth({
            credentials,
            scopes: ['https://www.googleapis.com/auth/spreadsheets.readonly']
        });

        const sheets = google.sheets({ version: 'v4', auth });

        // Get the spreadsheet data
        const response = await sheets.spreadsheets.values.get({
            spreadsheetId: process.env.GOOGLE_SHEET_ID,
            range: sheetName
        });

        const rows = response.data.values;
        if (!rows || rows.length === 0) {
            throw new Error('No data found in spreadsheet');
        }

        // Convert to JSON with header row as keys
        const headers = rows[0];
        const jsonData = rows.slice(1).map(row => {
            const obj = {};
            headers.forEach((header, index) => {
                obj[header] = row[index] || ''; // Use empty string as default value
            });
            return obj;
        });

        return jsonData;
    } catch (error) {
        console.error('Error loading data from Google Sheets:', error);
        throw error;
    }
}

async function postDataSequentially(data, version = 'v1') {
    const backendUrl = process.env.BACKEND_URL || 'http://localhost:8080';
    for (const item of data) {
        const postData = JSON.stringify(item);
        const path = version === 'v1' ? '/api/v1/store' : '/api/v2/store';

        const options = {
            hostname: new URL(backendUrl).hostname,
            port: new URL(backendUrl).port,
            path,
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Content-Length': Buffer.byteLength(postData),
                'X-API-Key': process.env.API_KEY
            }
        };

        await new Promise((resolve, reject) => {
            const req = http.request(options, (res) => {
                let data = '';
                res.on('data', (chunk) => {
                    data += chunk;
                });
                res.on('end', () => {
                    console.log(`Response: ${data}`);
                    resolve();
                });
            });

            req.on('error', (e) => {
                console.error(`Problem with request: ${e.message}`);
                reject(e);
            });

            req.write(postData);
            req.end();
        });

        await new Promise(resolve => setTimeout(resolve, 100));
    }
}

function checkHealth() {
    return new Promise((resolve, reject) => {
        const backendUrl = process.env.BACKEND_URL || 'http://localhost:8080';
        http.get(`${backendUrl}/health`, (resp) => {
            resolve(resp.statusCode);
        }).on("error", (err) => {
            reject("Error: " + err.message);
        });
    });
}

async function backfillV1() {
    try {
        const data = await loadData('Portofolio Model');
        console.log('Loaded data for v1 backfill:', data.length, 'records');
        const healthStatusCode = await checkHealth();
        console.log('Health Check Status Code:', healthStatusCode);

        if (healthStatusCode === 200) {
            // await postDataSequentially(data, 'v1');
            console.log('V1 backfill completed successfully');
        } else {
            console.error('Health check failed. Backend might be unavailable.');
        }
    } catch (error) {
        console.error('Error during v1 backfill:', error);
    }
}

async function backfillV2() {
    try {
        const rawData = await loadData('Portofolio Model v2');
        const data = rawData.map(item => ({
            metadata: {
                company: item.company,
                role: item.role,
                project: item.project,
                category: item.category,
                year: item.year,
                content: item.content,
                granularity: 'sentence'
            },
            content: item.content
        }));

        console.log('Loaded data for v2 backfill:', data.length, 'records');
        const healthStatusCode = await checkHealth();
        console.log('Health Check Status Code:', healthStatusCode);

        if (healthStatusCode === 200) {
            await postDataSequentially(data, 'v2');
            console.log('V2 backfill completed successfully');
        } else {
            console.error('Health check failed. Backend might be unavailable.');
        }
    } catch (error) {
        console.error('Error during v2 backfill:', error);
    }
}

async function saveDataToJson(data, filename) {
    try {
        // Create data directory if it doesn't exist
        const dataDir = path.join(__dirname, 'data');
        if (!fs.existsSync(dataDir)) {
            fs.mkdirSync(dataDir, { recursive: true });
        }

        const filePath = path.join(dataDir, filename);
        fs.writeFileSync(filePath, JSON.stringify(data, null, 2));
        console.log(`Data saved successfully to ${filePath}`);
    } catch (error) {
        console.error('Error saving data to JSON:', error);
        throw error;
    }
}

async function fetchAndSaveData() {
    try {
        const sheetName = process.env.GOOGLE_SHEET_NAME;
        const data = await loadData(sheetName);
        const timestamp = new Date().toISOString().replace(/[:.]/g, '-');
        const filename = `portfolio-data-${timestamp}.json`;
        await saveDataToJson(data, filename);
    } catch (error) {
        console.error('Error during data fetch and save:', error);
    }
}

program
    .name('backfill-tool')
    .description('CLI tool for running data backfills')
    .version('1.0.0');

program.command('v1')
    .description('Run backfill using v1 implementation')
    .action(backfillV1);

program.command('v2')
    .description('Run backfill using v2 implementation')
    .action(backfillV2);

program.command('fetch')
    .description('Fetch data from Google Sheets and save to JSON file')
    .action(fetchAndSaveData);

program.parse(process.argv);
