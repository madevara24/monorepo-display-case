/**
 * API utility to handle different API endpoints
 */

// Base API URL from environment (either from .env or deployed environment)
const BASE_URL = import.meta.env.VITE_API_URL || '';
// Environment setting, defaults to production for safety if not specified
const APP_ENV = import.meta.env.VITE_APP_ENV || 'production';
// Check if we're in production mode
const IS_PRODUCTION = APP_ENV === 'production';

/**
 * Get full URL for an API endpoint
 * @param path The API path from API_PATHS
 * @returns The complete URL
 */
export const getApiUrl = (path: string): string => {
  // For development - use localhost
  if (!IS_PRODUCTION) {
    return `http://localhost:8080${path}`;
  }
  // For production - use relative paths (will use same host as frontend)
  return path;
};

// API paths - consistent across environments
export const API_PATHS = {
  // Backend API endpoints
  ASK: '/api/v2/ask',
  STORE: '/api/v2/store',
  
  // Root endpoints
  HEALTH: '/health',
};

// For debugging
console.log('API Configuration:', {
  environment: APP_ENV,
  isProduction: IS_PRODUCTION,
  baseUrl: BASE_URL,
  healthPath: API_PATHS.HEALTH,
  askPath: API_PATHS.ASK
});

// Export a default configuration object for fetch
export const defaultFetchConfig = {
  headers: {
    'Content-Type': 'application/json',
  },
  credentials: 'include' as RequestCredentials,
};

export default {
  getApiUrl,
  API_PATHS,
  defaultFetchConfig,
}; 