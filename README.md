# Go Microservices Observability

This project demonstrates a high-performance observability pipeline for a **Go-based microservices architecture**. It utilizes **Grafana Alloy** as a central collector to unify logs and metrics across multiple services, routing them to **Loki** and **Prometheus** for centralized visualization in **Grafana**.

## 🌟 Key Features

* **Distributed Telemetry Collection**: Uses **Grafana Alloy** as a single agent to aggregate logs and metrics from multiple independent services (`service-a` and `service-b`).
* **Automated Log Pipeline**: Real-time JSON log parsing and smart filtering to remove noise (e.g., health checks) before storage.
* **Isolated Service Environments**: Each microservice is containerized with its own environment and lifecycle, managed via a centralized orchestration layer.
* **Unified Monitoring Stack**: A shared infrastructure for Loki (Logs), Prometheus (Metrics), and Grafana (Dashboards).

## 🛠 Tech Stack

* **Language**: Go (Golang)
* **Collector/Agent**: Grafana Alloy
* **Log Storage**: Grafana Loki
* **Metrics Database**: Prometheus
* **Visualization**: Grafana
* **Orchestration**: Docker Compose & Makefile

## 📁 Project Structure

```text
.
├── services/
│   ├── service-a/
│   │   ├── app/                # Core logic (logger, metrics, middleware, server)
│   │   ├── config/             
│   │   │   └── alloy.hcl       # Alloy config for Service A
│   │   ├── .env                # Service A environment
│   │   ├── docker-compose.yml
│   │   ├── Dockerfile          
│   │   ├── go.mod / go.sum
│   │   └── main.go             
│   └── service-b/   
├── .env                        # Root monitoring environment
├── .gitignore
├── alloy.hcl                   # Main root Alloy configuration
├── docker-compose.yml          # Core monitoring stack (Grafana, Prometheus, Loki, Alloy)
├── grafana.yml                 # Grafana provisioning
├── loki.yml                    # Loki configuration
├── Makefile                    # Project management automation
└── prometheus.yml              # Prometheus configuration
```

## How to run

1. Copy file `.env.example` to `.env`.
  ```sh
  cp .env.example .env
  ```

2. Run the Project.
  ```sh
  make up
  ```

3. Makefile Commands
  ```sh
   make up           // Build and start monitoring + all services
   make start        // Start existing containers without rebuilding
   make stop         // Stop all running containers
   make down         // Stop and remove all containers, volumes, and networks
   make monitoring   // Start only the core monitoring stack
   make logs         // View real-time logs for the monitoring stack
  ```

4. Access Dashboards:
  * Grafana: http://localhost:3000
  * Prometheus: http://localhost:9090
  * Alloy Dashboard: http://localhost:12345

## ⚙️ How It Works
1. **Go Application**: Each microservice generates JSON-formatted logs and exposes time-series metrics via a /metrics endpoint.

2. **Grafana Alloy**: Acts as a central collector that discovered containers, scrapes metrics, and captures log streams.

3. **Metrics Processing**: Alloy pushes collected metrics to **Prometheus** via its remote-write receiver.

4. **Log Aggregation**: Alloy tails log streams from the services and ships them to Loki for centralized storage.

5. **Visualization**: **Grafana** pulls data from both Loki and Prometheus to provide a unified observability dashboard.

## 📊 Troubleshooting
To check component-specific logs, run:
  ```sh
   make logs-grafana
   make logs-prometheus
   make logs-loki
  ```