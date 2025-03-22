# Variables
CYAN=\033[0;36m
NC=\033[0m # No Color

.PHONY: help backend frontend all clean

help:
	@echo "${CYAN}Available commands:${NC}"
	@echo "make backend  - Build backend service"
	@echo "make frontend - Build frontend service"
	@echo "make all      - Build all services"
	@echo "make clean    - Clean all services"

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

.PHONY: infra
infra:
    @echo "${CYAN}Deploying infrastructure services...${NC}"
    kubectl apply -f k8s/postgres/
    kubectl apply -f k8s/logging/