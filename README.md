<<<<<<< HEAD
---

# Docker Container Health Monitor (Go + WSL)

**Status:** ✅ Complete | **Time:** ~6 hours | **Impact:** Automated Container Visibility

---

## 📊 Metrics at a Glance

| Metric               | Value                   |
| -------------------- | ----------------------- |
| Time to Build        | ~6 hours                |
| Lines of Code        | ~420                    |
| Containers Monitored | Unlimited               |
| Alert Method         | HTTP Webhook            |
| Check Interval       | Configurable            |
| Environment          | WSL2 + Docker Desktop   |
| Dependencies         | None (Std Library Only) |

---

## 🎯 What It Does

A lightweight CLI monitoring tool that automatically checks Docker container health and sends alerts when issues occur.

**Monitors Containers For:**

* ❌ Not running
* ⚠️ Unhealthy
* 🔍 Not found
* ✅ Running normally

**Before:** Manually running `docker ps` and `docker inspect` repeatedly
**After:** Run one command → Continuous automated monitoring

---

## 🚀 Quick Start

### Run It (WSL Recommended)

```bash
cd /mnt/c/Users/<your-user>/go-practice
go run .
```

---

### Example Output

```
🚀 Starting Docker Container Monitor
⏱️  Check interval: 30 seconds
📦 Monitoring 3 containers

=== Container Health Check ===
Time: 2026-02-08 21:47:11

✅ nginx
   ID: abc123456789
   State: running

❌ redis
   ID: def987654321
   State: exited
```

If any container is unhealthy → JSON webhook alert is sent automatically.

---

## 🧠 How It Works

### Architecture Flow

1. Load `config.json`
2. Start ticker based on `check_interval`
3. For each container:

   * Execute `docker inspect <name>`
   * Parse JSON output
4. Print results to terminal
5. Send webhook alert if unhealthy

---

## 📁 Project Structure

```
go-practice/
├── config.json
├── go.mod
├── main.go
│
├── config/
│   └── config.go
│
├── internal/
│   ├── monitor/
│   │   └── monitor.go
│   └── notifier/
│       └── notifier.go
```

---

## 📄 What Each File Does

**config.json** – Runtime configuration (interval, webhook, containers)
**go.mod** – Module definition
**main.go** – Entry point, ticker loop, graceful shutdown
**config.go** – Parses config and applies defaults
**monitor.go** – Executes `docker inspect` and evaluates health
**notifier.go** – Prints status + sends webhook alerts

---

## 🔧 Core Code Highlights

### Module Definition — `go.mod`

```go
module github.com/dr-musa-bala/WordPress-Setup-on-xFusionCorp-Infra

go 1.20
```

---

### Runtime Configuration — `config.json`

```json
{
  "docker_socket": "/var/run/docker.sock",
  "alert_webhook": "http://localhost:9999/alert",
  "containers": ["nginx", "postgres", "redis"]
}
```

---

### Entry Point Logic — `main.go` (Excerpt)

```go
mon := monitor.New(cfg.DockerSocket)
notif := notifier.New(cfg.AlertWebhook)

ticker := time.NewTicker(time.Duration(cfg.CheckInterval) * time.Second)

for {
    select {
    case <-ticker.C:
        runCheck(ctx, mon, notif, cfg.Containers)
    case <-sigChan:
        fmt.Println("\n👋 Shutting down gracefully...")
        return
    }
}
```

**Key Points**

* Uses Go ticker for periodic checks
* Graceful shutdown via OS signals
* Separates monitoring and notification concerns

---

### Config Loader — `config.go` (Excerpt)

```go
if config.DockerSocket == "" {
    config.DockerSocket = "/var/run/docker.sock"
}
if config.CheckInterval == 0 {
    config.CheckInterval = 30
}
```

**Insight:** Safe defaults prevent runtime crashes.

---

### Container Inspection — `monitor.go` (Excerpt)

```go
cmd := exec.CommandContext(ctx, "docker", "inspect", containerName)
out, err := cmd.Output()
```

**Why CLI instead of Socket?**

* Fewer permission issues
* Cross-platform reliability
* Works seamlessly inside WSL

**Concurrency Used**

```go
go func(idx int, name string) {
    st, _ := m.CheckContainer(ctx, name)
    ch <- res{i: idx, s: st}
}(i, c)
```

Parallel container checks significantly reduce wait time.

---

### Webhook Alerts — `notifier.go` (Excerpt)

```go
resp, err := n.client.Post(
    n.webhookURL,
    "application/json",
    bytes.NewBuffer(payload),
)
```

* Sends structured JSON
* Timeout-controlled HTTP client
* Only triggers when unhealthy containers exist

---

## 📈 Performance

| Containers | Manual Time | Automated Time | Improvement |
| ---------- | ----------- | -------------- | ----------- |
| 3          | 3–5 min     | Instant        | ~95%        |
| 10         | 10–15 min   | Instant        | ~98%        |

**Impact**

* Eliminates repetitive terminal checks
* Enables passive background monitoring
* Ideal for local DevOps experimentation

---

## 🎓 What I Learned

| Concept       | Application               | Confidence |
| ------------- | ------------------------- | ---------- |
| Concurrency   | Parallel container checks | 🟢 Solid   |
| JSON Parsing  | Docker inspect + config   | 🟢 Solid   |
| OS Signals    | Graceful shutdown         | 🟢 Solid   |
| CLI Execution | `exec.Command`            | 🟢 Solid   |
| HTTP Clients  | Webhook alerts            | 🟢 Solid   |

---

### Technical Insights

* **WSL + Docker Integration** is more stable than native Windows CLI.
* **CLI over Docker Socket** avoided permission failures.
* **Safe Defaults** reduce runtime configuration errors.
* **Modular Design** simplified debugging and testing.

---

## 🐛 Challenges & Solutions

| Challenge                       | Solution                   |
| ------------------------------- | -------------------------- |
| Docker socket permission errors | Switched to CLI inspection |
| Import path mismatch            | Fixed `go.mod` module path |
| Go not found in WSL             | Installed Go inside Ubuntu |
| Webhook testing difficulty      | Local Python receiver      |

---

## 🔄 Future Enhancements

* Prometheus metrics
* Grafana dashboards
* Slack / Email alerts
* Web UI dashboard
* Live config reload
* CPU / Memory tracking

---

## 🎯 Use Cases

**Personal**

* Local Docker container monitoring
* Learning Go + DevOps integration

**Professional**

* Development container oversight
* Pre-deployment validation

**DevOps**

* CI/CD pipeline checks
* Automated alert experiments

---

## 🏷️ Tech Stack

* **Language:** Go 1.20+
* **Runtime:** Docker Desktop
* **Environment:** WSL2 Ubuntu
* **Alerting:** HTTP Webhooks
* **Monitoring:** Docker CLI (`docker inspect`)
* **Dependencies:** None

---

## 📊 Development Statistics

**Time Allocation**

* Planning — 45 min
* WSL + Docker Setup — 2 hrs
* Coding — 2 hrs
* Debugging — 1 hr
* Documentation — 30 min

**Code Composition**

* Logic — 75%
* Comments — 15%
* Config/Structure — 10%

---

## 👤 Author

**Dr. Musa Bala Audu**

---

## 📄 License

MIT License recommended for open-source distribution.

---
=======
# 🚀 DevOps Learning Journey

> A documented journey from beginner to DevOps engineer - building real tools, one concept at a time.

[![GitHub followers](https://img.shields.io/github/followers/dr-musa-bala?style=social)](https://github.com/dr-musa-bala)
[![GitHub stars](https://img.shields.io/github/stars/dr-musa-bala/devops-learning-journey?style=social)](https://github.com/dr-musa-bala/devops-learning-journey)

## 👋 About This Journey

I'm learning DevOps by building real automation tools and documenting everything I learn. This repository is my public learning log - tracking progress, sharing insights, and building in public.

**Start Date:** February 06, 2026  
**Current Day:** 1 of 100  
**Current Focus:** Go Programming Fundamentals

## 📊 Progress at a Glance

| Metric | Count |
|--------|-------|
| 🗓️ Days Learning | 1 |
| 🛠️ Projects Built | 2 |
| 📚 Concepts Learned | 5 |
| 💻 Lines of Code | 163 |
| 📝 Git Commits | 1 |

**Last Updated:** Feb 06, 2026

## 🛠️ Projects Built

### 0. [Hello World](./projects/00-hello-world) - Day 1
**What it does:** My first Go program  
**Key learning:** Go package system, build process, environment setup  
**Status:** ✅ Complete

### 1. [Automated File Organizer](./projects/02-file-organizer) - Day 1
**What it does:** Organizes files by type automatically  
**Key learning:** File I/O, pattern matching, categorization  
**Status:** ✅ Complete  
**Impact:** Reduced file organization time from 20 minutes to 2 seconds

## 📚 What I'm Learning

### Current Week (Week 1)
- [x] Go basics (variables, functions, types)
- [x] File system operations
- [x] Error handling patterns
- [ ] Docker fundamentals
- [ ] Git workflows

### This Month
- Build 4 automation tools in Go
- Learn Docker container management
- Understand CI/CD basics
- Set up GitHub Actions

## 🎯 Learning Goals

**Short-term (Month 1):**
- Master Go programming fundamentals
- Understand Docker and containers
- Learn basic Kubernetes concepts

**Long-term (6 Months):**
- Build production-ready DevOps tools
- Contribute to open-source projects
- Land a DevOps engineering role
- Help others on their DevOps journey

## 💡 Latest Insights

### Day 1 - Key Takeaway
> "Go compiles to standalone binaries with zero dependencies. This is a game-changer 
> for DevOps tools - no Python virtual environments, no Node.js modules, just one 
> executable that runs anywhere!"

[Read all daily insights →](./docs/learnings/)

## 🔧 Tech Stack

**Currently Learning:**
- Go 1.25+
- Git & GitHub
- VS Code
- Docker (upcoming)
- Kubernetes (upcoming)

**Development Environment:**
- OS: Windows 11
- Terminal: PowerShell, GitBash
- Editor: VS Code with Go extension

## 📖 How to Use This Repository

Each project includes:
- ✅ Complete, working source code
- ✅ Detailed README explaining what I learned
- ✅ Challenges faced and solutions found
- ✅ Time invested and metrics
- ✅ Next steps and future improvements

Feel free to:
- ⭐ Star this repo if you find it helpful
- 🔀 Fork it to start your own learning journey
- 💬 Open issues with questions or suggestions
- 🤝 Connect with me on [LinkedIn](www.linkedin.com/in/musa-bala-audu-o-d-57b906113/)

## 🌱 Why I'm Learning in Public

1. **Accountability** - Public commitment keeps me consistent
2. **Documentation** - Future me will thank present me
3. **Community** - Helping others who are on the same path
4. **Portfolio** - Evidence of continuous learning
5. **Growth** - Feedback makes me better

## 🔗 Connect With Me

- 💼 [LinkedIn](https://www.linkedin.com/in/musa-bala-audu-o-d-57b906113/) - Professional updates
- 🐦 [Twitter/X](@sight_musa) - Daily progress
- 📧 [Email](freshabdullaah@gmail.com) - Let's talk DevOps!

## 📈 Weekly Progress

### Week 1 (Jan 06-13, 2026)
- ✅ Set up Go development environment
- ✅ Built Hello World program
- ✅ Created file organizer tool
- 🔄 Learning Docker basics
- 📅 Planning first container project

## 🙏 Inspired By

- The #100DaysOfCode community
- #DevOps community on Twitter/LinkedIn
- Every developer who learns in public

---

**"The expert in anything was once a beginner."** - Helen Hayes


---

Last commit: Just getting started!  
Next milestone: 10 days of consistent learning
💪 Let's build something amazing, one day at a time.
>>>>>>> cfd4bd7bdf1634de16bb7d4fb51f7fb2e615e7a6

