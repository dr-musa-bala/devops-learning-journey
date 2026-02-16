---

# 📡 DevOps Health Check API (Go-Powered Microservice)

**Status:** 🚀 Active Service | **Focus:** Observability & Monitoring Integration

## 🎯 Overview

In modern DevOps, tools must be able to communicate with one another. This project transitions from a standalone CLI tool to a **persistent background service**. It exposes system health data via a RESTful JSON API, allowing monitoring tools (like Prometheus or custom Bash scripts) to query the environment's readiness in real-time.

## 🛠 Technical Implementation (The Go Standard Library)

This project serves as the "capstone" for the Go learning phase, implementing several high-level concepts:

* **`net/http` Server:** Orchestrated a persistent HTTP server using the Go standard library, avoiding heavy frameworks to keep the binary footprint minimal.
* **JSON Marshalling:** Utilized `encoding/json` to translate complex Go `structs` into machine-readable JSON formats.
* **Environment-Driven Configuration:** Implemented `os.Getenv` for port assignment, a "Cloud Native" requirement for deploying to Docker and Kubernetes.
* **Persistent State:** Tracked system uptime by calculating the delta between a global `startTime` variable and the current request time.

## 📊 Performance & Service Metrics

This service is optimized for high-frequency polling with minimal resource overhead.

| Metric | Target / Result | DevOps Significance |
| --- | --- | --- |
| **Response Format** | `application/json` | Universal compatibility with Bash, Python, and JS. |
| **Startup Time** | < 10ms | Critical for fast-scaling container environments. |
| **Memory Footprint** | ~2MB (Idle) | Allows for "sidecar" deployment without resource theft. |
| **Configuration** | Env-Variable Injectable | Decouples code from infrastructure (12-Factor App). |

---

## 📈 DevOps Learning Path: From Go to Bash

This project marks the successful completion of the Go foundations and serves as the **bridge** to Linux Bash Scripting.

1. **Phase 1 (Go):** Building the "Source of Truth" (This API).
2. **Phase 2 (Bash):** Building the "Consumer." My next step is to write Bash scripts that `curl` this API, parse the JSON with `jq`, and automate system alerts based on the results.

## 🤝 Team Iteration & Roadmap

Following the principle of **Team Ownership**, this API is designed for easy expansion:

* **Integration:** Add a function to the `healthCheckHandler` that calls the "Onboarding Auditor" code to report tool status via the web.
* **Security:** Implement a simple API Key check using HTTP headers to restrict access.
* **Logging:** Add the `log` package to record every request's IP address and timestamp to the terminal.

---

## 🚀 Quick Start

```bash
# Set a custom port (optional)
export APP_PORT=9000

# Run the service
go run main.go

# Query from another terminal
curl http://localhost:9000/health

```

---

