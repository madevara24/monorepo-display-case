# Monorepo Display Case Helm Chart

This Helm chart deploys the complete Monorepo Display Case stack including:

- Frontend application
- Backend service
- PostgreSQL database with Vector extension
- Logging stack (Loki, Promtail, Grafana)

## Prerequisites

- Kubernetes 1.19+
- Helm 3.2.0+
- PV provisioner support in the underlying infrastructure
- Storage directories on the host machine:
  - `/opt/postgres-data`
  - `/opt/loki-data`
  - `/opt/grafana-data`
  - `/opt/promtail-positions`

## Installation Steps

1. **Create required storage directories** on your host machine:
   ```bash
   sudo mkdir -p /opt/{postgres-data,loki-data,grafana-data,promtail-positions}
   sudo chown -R 999:999 /opt/postgres-data
   sudo chown -R 10001:10001 /opt/loki-data
   sudo chown -R 472:472 /opt/grafana-data
   sudo chown -R 10001:10001 /opt/promtail-positions
   ```

2. **Create required secrets** before deployment:

### TLS Secret for HTTPS
This chart expects a TLS secret for HTTPS configuration. You must create this secret manually before deploying:

```bash
# Create TLS secret for HTTPS
kubectl create secret tls madevara24-tls \
  --cert=/path/to/fullchain.pem \
  --key=/path/to/privkey.pem
```

### PostgreSQL Secret
This chart requires a Kubernetes secret named `postgres-secret` with the following keys:
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `POSTGRES_DB`

Create this secret manually:
```bash
# Example (replace with your actual credentials)
kubectl create secret generic postgres-secret \
  --from-literal=POSTGRES_USER=admin \
  --from-literal=POSTGRES_PASSWORD=secure_password \
  --from-literal=POSTGRES_DB=mydatabase
```

3. **Install the chart**:
```bash
helm install my-app ./monorepo-display-case
```

For a custom configuration:
```bash
helm install my-app ./monorepo-display-case -f custom-values.yaml
```

## Accessing the Services

- **Frontend**: http://your-domain.com
- **Backend API**: http://your-domain.com/api
- **Grafana dashboard**: http://your-domain.com/grafana (if you set up appropriate ingress rules)

## Configuration

The following tables list the configurable parameters and their default values.

### Global Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `global.environment` | Environment name | `production` |
| `global.registry` | Docker registry for images | `localhost:5000` |

### Frontend Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `frontend.replicaCount` | Number of frontend replicas | `1` |
| `frontend.image.repository` | Frontend image repository | `frontend` |
| `frontend.image.tag` | Frontend image tag | `latest` |
| `frontend.service.type` | Frontend service type | `ClusterIP` |

### Backend Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `backend.replicaCount` | Number of backend replicas | `1` |
| `backend.image.repository` | Backend image repository | `backend-service` |
| `backend.image.tag` | Backend image tag | `latest` |
| `backend.port` | Backend service port | `8080` |
| `backend.service.type` | Backend service type | `ClusterIP` |

### PostgreSQL Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `postgres.image.repository` | PostgreSQL image repository | `postgres-vector` |
| `postgres.image.tag` | PostgreSQL image tag | `latest` |
| `postgres.storage.size` | PostgreSQL storage size | `10Gi` |
| `postgres.config.maxConnections` | Max DB connections | `100` |
| `postgres.config.sharedBuffers` | Shared buffers size | `128MB` |
| `postgres.config.effectiveCacheSize` | Effective cache size | `384MB` |

### Logging Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `logging.loki.enabled` | Enable Loki | `true` |
| `logging.loki.storage.size` | Loki storage size | `5Gi` |
| `logging.promtail.enabled` | Enable Promtail | `true` |
| `logging.grafana.enabled` | Enable Grafana | `true` |
| `logging.grafana.storage.size` | Grafana storage size | `2Gi` |

### Ingress Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `ingress.enabled` | Enable ingress | `true` |
| `ingress.className` | Ingress class name | `traefik` |
| `ingress.domain` | Domain name | `madevara24.space` |
| `ingress.tls.enabled` | Enable TLS | `true` |
| `ingress.tls.secretName` | TLS secret name | `madevara24-tls` |

## Upgrading

To upgrade the chart:
```bash
helm upgrade my-app ./monorepo-display-case
```

## Uninstalling

To uninstall/delete the deployment:
```bash
helm uninstall my-app
```

## Backup and Restore

### PostgreSQL Backup
```bash
kubectl exec -it $(kubectl get pods -l app=postgres -o name) -- pg_dump -U admin mydatabase > backup.sql
```

### PostgreSQL Restore
```bash
kubectl exec -it $(kubectl get pods -l app=postgres -o name) -- psql -U admin mydatabase < backup.sql
```

## Persistence

This chart uses PersistentVolumeClaims for:

- PostgreSQL data
- Loki storage
- Grafana storage
- Promtail positions

Ensure you have proper PV provisioning in your cluster.

## Secrets

### TLS Secret for HTTPS
This chart expects a TLS secret for HTTPS configuration. You must create this secret manually before deploying:

```bash
# Create TLS secret for HTTPS
kubectl create secret tls madevara24-tls \
  --cert=/path/to/fullchain.pem \
  --key=/path/to/privkey.pem
```

### PostgreSQL Secret
This chart requires a Kubernetes secret named `postgres-secret` with the following keys:
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `POSTGRES_DB`

Create this secret manually:
```bash
# Example (replace with your actual credentials)
kubectl create secret generic postgres-secret \
  --from-literal=POSTGRES_USER=admin \
  --from-literal=POSTGRES_PASSWORD=secure_password \
  --from-literal=POSTGRES_DB=mydatabase
```

## Next Steps

Now that you have your chart set up, continue by adding templates for:

1. Frontend deployment and service
2. PostgreSQL statefulset and service
3. Logging stack components (Loki, Promtail, Grafana)
4. ConfigMaps and secrets templates 
