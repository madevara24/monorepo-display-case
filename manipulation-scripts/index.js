const fs = require('fs');
const path = require('path');
const http = require('http');
const xlsx = require('xlsx');
require('dotenv').config();

function loadData() {
    const filePath = path.join(__dirname, 'data.xlsx');
    const workbook = xlsx.readFile(filePath);
    const sheetName = workbook.SheetNames[0];
    const worksheet = workbook.Sheets[sheetName];
    const jsonData = xlsx.utils.sheet_to_json(worksheet, { defval: '' });

    return jsonData;
}

async function postDataSequentially(data) {
    const backendUrl = process.env.BACKEND_URL || 'http://localhost:8080';
    for (const item of data) {
        const postData = JSON.stringify(item);

        const options = {
            hostname: new URL(backendUrl).hostname,
            port: new URL(backendUrl).port,
            path: '/store',
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Content-Length': Buffer.byteLength(postData)
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

async function main() {
    try {
        const data = loadData();
        console.log(data);
        const healthStatusCode = await checkHealth();
        console.log('Health Check Status Code:', healthStatusCode);

        await postDataSequentially(data);
    } catch (error) {
        console.error(error);
    }
}

main();
