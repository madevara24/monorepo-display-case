# Frontend Configuration Guide for Kubernetes

## Environment Variables in Kubernetes

The Kubernetes deployment of this frontend application uses environment variables to configure the backend URL. This approach is simpler than a full runtime configuration file but still allows you to change the backend URL without rebuilding the container.

## Required Changes to Your Frontend Code

Your frontend container needs to be built to read environment variables at runtime. Here are two common approaches:

### Approach 1: Environment Variable Substitution at Startup

If you're using NGINX to serve your frontend, you can modify your Dockerfile to replace placeholders with environment variables at startup:

1. Add this to your `nginx.conf` or `default.conf`:
```nginx
server {
    # ... other config ...
    
    location / {
        # ... other config ...
        sub_filter_once off;
        sub_filter_types application/javascript;
        sub_filter 'BACKEND_URL_PLACEHOLDER' '$BACKEND_URL';
    }
}
```

2. In your frontend JavaScript, use a placeholder that will be replaced at runtime:
```javascript
const apiUrl = 'BACKEND_URL_PLACEHOLDER';
```

### Approach 2: Window Environment Variables

This approach is simpler but requires your container entrypoint to generate a config file:

1. Create an entrypoint script in your Dockerfile:
```bash
#!/bin/sh
# Generate config with environment variables
echo "window.ENV = { API_URL: '$BACKEND_URL' };" > /usr/share/nginx/html/config.js
# Start nginx
nginx -g 'daemon off;'
```

2. Include a script tag in your HTML:
```html
<script src="/config.js"></script>
```

3. In your code:
```javascript
const apiUrl = window.ENV.API_URL || 'http://localhost:8080';
```

## Kubernetes Configuration

In our Helm chart, we're using option 2, setting the `BACKEND_URL` environment variable based on your Helm values:

```yaml
env:
  - name: BACKEND_URL
    value: {{ if eq .Values.frontend.api.mode "internal" }}{{ .Values.frontend.api.internalUrl }}{{ else }}{{ .Values.frontend.api.externalUrl }}{{ end }}
```

## Benefits of This Approach

- Simpler configuration than a full ConfigMap
- Still allows for environment-specific URLs without rebuilding
- Clear separation of code and configuration
- Works well when you have just a few configuration values

## Testing Locally

When developing locally:
- The `.env` file or hardcoded defaults will still work
- You can set environment variables when starting your dev server to test different configurations 