# 🐧 Linux Bash & Go: The Observability Pipeline

This project marks my transition from building standalone Go Microservices to managing them via **Linux Systems Engineering**. I’ve built a bridge between high-performance compiled code (Go) and flexible system automation (Bash).

## 🚀 The DevOps Workflow
1. **The Core:** A Go API (05-status-api) provides a "Source of Truth" for system health.
2. **The Monitor:** Bash scripts automate the "polling" of this API.
3. **The Parser:** Integrated `jq` to transform raw JSON strings into actionable system logs.

## 📊 Key Performance Metrics
| Metric | Improvement | Technical Detail |
| :--- | :--- | :--- |
| **Response Time** | < 10ms | Native Go HTTP handling |
| **Parsing Speed** | Instant | Optimized `jq` filtering |
| **Memory Footprint** | ~2.4MB | Zero-dependency Go binary + Bash |
| **Automation** | 100% | Cron-ready monitoring scripts |

## 🛠 Skills Mastered
* **I/O Redirection:** Using `> /dev/null` and `>>` for clean, persistent logging.
* **Piping (`|`):** Chaining `curl`, `grep`, and `jq` to create data pipelines.
* **Exit Codes:** Utilizing `$?` and conditional logic for fail-safe automation.
* **Command Substitution:** Using `$(date)` and `$(curl)` to inject dynamic data into scripts.

## 🏁 How to Run
1. Start the Go API: `go run main.go`
2. Execute the Monitor: `./check_status.sh`
3. Parse the JSON: `./read_json.sh`

## Screenshots
![Go JSON Observability before and after power on](screenshots/json_before_and_after_go_api_on.png) 
