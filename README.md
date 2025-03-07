# monorepo-display-case

Personal website for displaying myself

## Backend Service

Backend to serve question API, query vector DB, and fetch response from openAI

**Todo:**

- [ ] Security
  - [ ] Add API key for FE
  - [ ] Add rate limit per IP
  - [x] Add CORS
- [x] Finish question API
  - [x] Convert question into embedding
  - [x] Query vector database
  - [x] Send query result to OpenAI
  - [x] Return OpenAI response
- [x] Populate vector DB
  - [x] Get some embedding manually (curl)
  - [x] Make another API to store embeddings into vector DB
- [ ] Move security to API gateway
- [ ] Allow previous message to be taken into context
  - [ ] Take a look into langchain

## API Gateway

**Todo:**

- [ ] Make one
- [ ] Add API key for FE
- [ ] Add rate limit per IP
- [ ] Add CORS

## Frontend

- [x] Make one
- [x] Integrate w/ BE API
- [ ] Reroute to API gateway if built


## RAG Script/Content

- [ ] Add more details in a sentence type structure
  - [ ] Add details that is not mentioned on the web or CV
- [ ] Tinker around with summary structured data

## Logging

Loki + grafana + promtail standard logging

**Todo:**

- [x] Add logger

## Orchestration

`docker-compose up`

**Todo:**

- [x] Setup docker compose
