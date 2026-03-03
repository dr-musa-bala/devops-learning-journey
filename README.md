# 🚀 DevOps Learning Journey

![CI/CD Pipeline](https://github.com/dr-musa-bala/devops-learning-journey/actions/workflows/main.yml/badge.svg)

> **Status:** All Systems Operational. This repository tracks my transition from Linux administration to automated cloud architecture.

---

## 🌟 Current Milestone: Multi-Container Health API
**Location:** `/projects/06-linux-bash-basics`

This project is a production-ready **Go + Redis** microservice stack.

### 🛠️ Technical Achievements
* **Containerization**: Wrote a multi-stage Dockerfile to keep the Go binary small and secure.
* **Orchestration**: Used `docker-compose` to link the API to a persistent Redis database.
* **CI/CD Pipeline**: Automated testing with **ShellCheck** and multi-platform builds with **Docker Buildx**.

---

## 📂 The Journey So Far (Archive)

### Phase 1: Linux & Shell Scripting 🐧
*This was the foundation of the journey, focusing on the "OS" layer.*
* **Scripting**: Developed `read_json.sh` for automated data parsing.
* **Automation**: Set up Cron jobs for system health monitoring.
* **Tools**: Mastered `grep`, `sed`, `awk`, and Linux file permissions.

### Phase 2: Containers & Go (Current) 🐳
*Moving up the stack into application delivery and automation.*
* **Languages**: Transitioned to **Golang** for high-performance backend services.
* **Registries**: Configured automated pushes to **Docker Hub**.
* **GitHub Actions**: Built the "Robot" that audits every line of code I write.

---

## 🚀 How to Run the Latest Build
```bash
cd projects/06-linux-bash-basics
docker-compose up --build
