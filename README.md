# monorepo-display-case

Personal website for displaying myself

## Backend Service

Backend to serve question API, query vector DB, and fetch response from openAI

**Todo:**

- [ ] Security
  - [ ] Add API key for FE
  - [ ] Add rate limit per IP
  - [ ] Add CORS
- [x] Finish question API
  - [x] Convert question into embedding
  - [ ] Query vector database
  - [ ] Send query result to OpenAI
  - [ ] Return OpenAI response
- [ ] Populate vector DB
  - [ ] Get some embedding manually (curl)
  - [ ] Make another API to store embeddings into vector DB
- [ ] Move security to API gateway

## API Gateway

**Todo:**

- [ ] Make one
- [ ] Add API key for FE
- [ ] Add rate limit per IP
- [ ] Add CORS

## Frontend

- [ ] Make one
- [ ] Integrate w/ BE API
- [ ] Reroute to API gateway if built

## Logging

Loki + grafana + promtail standard logging

**Todo:**

- [x] Add logger

## Orchestration

`docker-compose up`

**Todo:**

- [x] Setup docker compose
