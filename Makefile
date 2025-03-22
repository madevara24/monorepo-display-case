# Variables
REGISTRY=localhost:5000
CYAN=\033[0;36m
NC=\033[0m # No Color

.PHONY: help backend frontend all clean k8s-init k8s-storage k8s-secrets k8s-deploy k8s-verify k8s-all

help:
	@echo "${CYAN}Available commands:${NC}"
	@echo "make backend     - Build backend service"
	@echo "make frontend    - Build frontend service"
	@echo "make all        - Build all services"
	@echo "make clean      - Clean all services"
	@echo ""
	@echo "${CYAN}Kubernetes commands:${NC}"
	@echo "make k8s-init    - Initialize directories and permissions"
	@echo "make k8s-storage - Apply storage configurations"
	@echo "make k8s-secrets - Apply secrets (requires manual setup first)"
	@echo "make k8s-deploy  - Deploy all services"
	@echo "make k8s-verify  - Verify deployments"
	@echo "make k8s-all     - Run all k8s steps (except secrets)"

backend:
	@echo "${CYAN}Building backend service...${NC}"
	$(MAKE) -C backend-service all

frontend:
	@echo "${CYAN}Building frontend service...${NC}"
	$(MAKE) -C frontend-display-case all

all: backend frontend

clean:
	$(MAKE) -C backend-service clean
	$(MAKE) -C frontend-display-case clean

# Kubernetes Commands
k8s-init:
	@echo "${CYAN}Creating required directories...${NC}"
	sudo mkdir -p /opt/postgres-data /opt/loki-data /opt/grafana-data /opt/promtail-positions
	@echo "${CYAN}Setting permissions...${NC}"
	sudo chown -R 999:999 /opt/postgres-data
	sudo chown -R 10001:10001 /opt/loki-data
	sudo chown -R 472:472 /opt/grafana-data

k8s-storage:
	@echo "${CYAN}Applying storage configurations...${NC}"
	kubectl apply -f k8s/logging/storage.yml
	kubectl apply -f k8s/postgres/configmap.yml

k8s-secrets:
	@echo "${CYAN}Checking if secret.yml exists...${NC}"
	@if [ ! -f k8s/postgres/secret.yml ]; then \
		echo "${CYAN}Error: k8s/postgres/secret.yml not found. Please create it from secret.example.yml${NC}"; \
		exit 1; \
	fi
	@echo "${CYAN}Applying secrets...${NC}"
	kubectl apply -f k8s/postgres/secret.yml

k8s-deploy:
	@echo "${CYAN}Building and pushing images...${NC}"
	docker build -t $(REGISTRY)/postgres-vector:latest -f postgres.Dockerfile .
	docker push $(REGISTRY)/postgres-vector:latest
	@echo "${CYAN}Deploying PostgreSQL...${NC}"
	kubectl apply -f k8s/postgres/statefulset.yml
	kubectl apply -f k8s/postgres/service.yml
	@echo "${CYAN}Deploying logging stack...${NC}"
	kubectl apply -f k8s/logging/loki-config.yml
	kubectl apply -f k8s/logging/promtail-config.yml
	kubectl apply -f k8s/logging/deployments.yml
	kubectl apply -f k8s/logging/services.yml

k8s-verify:
	@echo "${CYAN}Verifying deployments...${NC}"
	kubectl get pods
	@echo "\n${CYAN}Storage status:${NC}"
	kubectl get pv,pvc
	@echo "\n${CYAN}Services status:${NC}"
	kubectl get services

k8s-all: k8s-init k8s-storage k8s-deploy k8s-verify
	@echo "${CYAN}All k8s deployments completed!${NC}"
	@echo "${CYAN}Note: Secrets were not applied. Run 'make k8s-secrets' after setting up secret.yml${NC}"