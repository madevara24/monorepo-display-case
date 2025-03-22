# Service-specific variables and targets
REGISTRY=localhost:5000
SERVICE_NAME=frontend
VERSION=$(shell date +%Y%m%d)-$(shell git rev-parse --short HEAD)
CYAN=\033[0;36m
NC=\033[0m # No Color

.PHONY: help build push deploy all clean

help:
	@echo "${CYAN}Available commands:${NC}"
	@echo "make build    - Build the Docker image"
	@echo "make push     - Push to local registry"
	@echo "make deploy   - Deploy to k3s"
	@echo "make all      - Build, push, and deploy"
	@echo "make clean    - Remove local images"

build:
	@echo "${CYAN}Building ${SERVICE_NAME}:${VERSION}${NC}"
	docker build -t ${REGISTRY}/${SERVICE_NAME}:${VERSION} .
	docker tag ${REGISTRY}/${SERVICE_NAME}:${VERSION} ${REGISTRY}/${SERVICE_NAME}:latest

push: build
	@echo "${CYAN}Pushing ${SERVICE_NAME}:${VERSION} to registry${NC}"
	docker push ${REGISTRY}/${SERVICE_NAME}:${VERSION}
	docker push ${REGISTRY}/${SERVICE_NAME}:latest

deploy: push
	@echo "${CYAN}Deploying ${SERVICE_NAME}:${VERSION} to k3s${NC}"
	kubectl rollout restart deployment ${SERVICE_NAME}

all: deploy

clean:
	@echo "${CYAN}Cleaning up local images${NC}"
	docker rmi ${REGISTRY}/${SERVICE_NAME}:${VERSION} || true
	docker rmi ${REGISTRY}/${SERVICE_NAME}:latest || true
