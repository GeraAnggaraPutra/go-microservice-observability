MONITORING_DIR := .
SERVICES_DIR   := services

SERVICE_A := $(SERVICES_DIR)/service-a
SERVICE_B := $(SERVICES_DIR)/service-b

DC := docker compose

# DEFAULT
help:
	@echo ""
	@echo "Available commands:"
	@echo "  make up             -> Build + start monitoring & all services"
	@echo "  make start          -> Start existing containers (no build)"
	@echo "  make stop           -> Stop containers (no remove)"
	@echo "  make down           -> STOP + REMOVE containers, volumes & networks"
	@echo "  make restart        -> Restart everything (clean)"
	@echo ""
	@echo "  make monitoring     -> Start monitoring only"
	@echo "  make service-a      -> Start service A only"
	@echo "  make service-b      -> Start service B only"
	@echo ""
	@echo "  make logs           -> Show monitoring logs"
	@echo "  make logs-grafana   -> Show Grafana logs"
	@echo "  make logs-prometheus-> Show Prometheus logs"
	@echo "  make logs-loki      -> Show Loki logs"
	@echo ""

# BUILD + UP
up: monitoring services
	@echo "✅ All services are running"

monitoring:
	@echo "🚀 Starting monitoring stack (build)..."
	cd $(MONITORING_DIR) && $(DC) up -d --build

services: service-a service-b

service-a:
	@echo "🚀 Starting service A (build)..."
	cd $(SERVICE_A) && $(DC) up -d --build

service-b:
	@echo "🚀 Starting service B (build)..."
	cd $(SERVICE_B) && $(DC) up -d --build

# START
start: start-monitoring start-services
	@echo "▶️  All services started"

start-monitoring:
	@echo "▶️  Starting monitoring stack..."
	cd $(MONITORING_DIR) && $(DC) start

start-services: start-service-a start-service-b

start-service-a:
	@echo "▶️  Starting service A..."
	cd $(SERVICE_A) && $(DC) start

start-service-b:
	@echo "▶️  Starting service B..."
	cd $(SERVICE_B) && $(DC) start

# STOP
stop:
	@echo "⏸️  Stopping all containers..."
	cd $(SERVICE_A) && $(DC) stop || true
	cd $(SERVICE_B) && $(DC) stop || true
	cd $(MONITORING_DIR) && $(DC) stop || true

# DOWN (REMOVE EVERYTHING)
down:
	@echo "💣 Stopping & REMOVING containers, volumes, and networks..."
	cd $(SERVICE_A) && $(DC) down -v --remove-orphans || true
	cd $(SERVICE_B) && $(DC) down -v --remove-orphans || true
	cd $(MONITORING_DIR) && $(DC) down -v --remove-orphans || true

restart: down up

# LOGS
logs:
	cd $(MONITORING_DIR) && $(DC) logs -f

logs-grafana:
	cd $(MONITORING_DIR) && $(DC) logs -f grafana

logs-prometheus:
	cd $(MONITORING_DIR) && $(DC) logs -f prometheus

logs-loki:
	cd $(MONITORING_DIR) && $(DC) logs -f loki
