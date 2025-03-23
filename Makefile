# Variables
REGISTRY=localhost:5000
CYAN=\033[0;36m
GREEN=\033[0;32m
RED=\033[0;31m
NC=\033[0m # No Color

# Directories
POSTGRES_DIR=/opt/postgres-data
LOKI_DIR=/opt/loki-data
GRAFANA_DIR=/opt/grafana-data
PROMTAIL_DIR=/opt/promtail-positions

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
	@echo "${CYAN}Checking and creating required directories...${NC}"
	@for dir in $(POSTGRES_DIR) $(LOKI_DIR) $(GRAFANA_DIR) $(PROMTAIL_DIR); do \
		if [ ! -d $$dir ]; then \
			echo "Creating $$dir..."; \
			sudo mkdir -p $$dir || { echo "${RED}Failed to create $$dir${NC}"; exit 1; }; \
			echo "${GREEN}Created $$dir${NC}"; \
		else \
			echo "$$dir already exists"; \
		fi \
	done
	@echo "${CYAN}Setting correct permissions...${NC}"
	@# PostgreSQL user (999:999)
	@sudo chown -R 999:999 $(POSTGRES_DIR) && \
		echo "${GREEN}Set permissions for PostgreSQL data directory${NC}" || \
		{ echo "${RED}Failed to set PostgreSQL permissions${NC}"; exit 1; }
	@# Loki user (10001:10001)
	@sudo chown -R 10001:10001 $(LOKI_DIR) && \
		echo "${GREEN}Set permissions for Loki data directory${NC}" || \
		{ echo "${RED}Failed to set Loki permissions${NC}"; exit 1; }
	@# Grafana user (472:472)
	@sudo chown -R 472:472 $(GRAFANA_DIR) && \
		echo "${GREEN}Set permissions for Grafana data directory${NC}" || \
		{ echo "${RED}Failed to set Grafana permissions${NC}"; exit 1; }
	@# Promtail directory (same as Loki)
	@sudo chown -R 10001:10001 $(PROMTAIL_DIR) && \
		echo "${GREEN}Set permissions for Promtail positions directory${NC}" || \
		{ echo "${RED}Failed to set Promtail permissions${NC}"; exit 1; }
	@echo "${GREEN}All directories are set up correctly${NC}"

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
	@echo "${CYAN}Building and pushing infrastructure images...${NC}"
	@echo "Building postgres-vector..."
	docker build -t $(REGISTRY)/postgres-vector:latest -f postgres.Dockerfile . || \
		{ echo "${RED}Failed to build postgres-vector${NC}"; exit 1; }
	docker push $(REGISTRY)/postgres-vector:latest || \
		{ echo "${RED}Failed to push postgres-vector${NC}"; exit 1; }

	@echo "${CYAN}Building application images...${NC}"
	@echo "Building backend-service..."
	docker build -t $(REGISTRY)/backend-service:latest ./backend-service || \
		{ echo "${RED}Failed to build backend-service${NC}"; exit 1; }
	docker push $(REGISTRY)/backend-service:latest || \
		{ echo "${RED}Failed to push backend-service${NC}"; exit 1; }

	@echo "Building frontend..."
	docker build -t $(REGISTRY)/frontend:latest ./frontend-display-case || \
		{ echo "${RED}Failed to build frontend${NC}"; exit 1; }
	docker push $(REGISTRY)/frontend:latest || \
		{ echo "${RED}Failed to push frontend${NC}"; exit 1; }

	@echo "${CYAN}Deploying infrastructure services...${NC}"
	@echo "Deploying PostgreSQL..."
	kubectl apply -f k8s/postgres/statefulset.yml
	kubectl apply -f k8s/postgres/service.yml

	@echo "Deploying logging stack..."
	kubectl apply -f k8s/logging/loki-config.yml
	kubectl apply -f k8s/logging/promtail-config.yml
	kubectl apply -f k8s/logging/deployments.yml
	kubectl apply -f k8s/logging/services.yml

	@echo "${CYAN}Deploying application services...${NC}"
	kubectl apply -f k8s/backend/deployment.yml
	kubectl apply -f k8s/backend/service.yml
	kubectl apply -f k8s/frontend/deployment.yml
	kubectl apply -f k8s/frontend/service.yml

	@echo "${GREEN}All services deployed successfully${NC}"

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