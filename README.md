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

## 🤖 CI/CD Automation
This repo is now automated! 

- **Workflow:** `.github/workflows/main.yml`
- **Docker Image:** `musabalaaudu/health-api:main`
- **Port Mapping:** `-p 8080:8080`

### Quick Commands
```bash
# Pull and Run the Cloud version
docker pull musabalaaudu/health-api:main
docker run -d -p 8080:8080 musabalaaudu/health-api:main

---

# 🚀 DevOps Journey: From Broken Pipelines to Containerized Excellence

## 📌 Project Overview

This project demonstrates a full CI/CD lifecycle for a Go-based Health API and its associated Bash automation tools. It covers the transition from local development to automated linting, containerization, and cross-environment networking.

---

## 🏗️ The Architecture

1. **Backend**: A Golang API providing system health metrics.
2. **Automation**: Bash scripts (`check_status.sh`, `read_json.sh`) for monitoring.
3. **CI/CD**: GitHub Actions utilizing `ShellCheck` for linting and Docker for image distribution.
4. **Containerization**: A multi-purpose Docker image used to run the automation suite.

---

## 🛠️ The "DevOps Trial by Fire": Debugging Log

### 1. The Invisible CI/CD Gatekeeper

**The Issue:** We integrated `ShellCheck` into `main.yml`, but the pipeline was passing even when the code was intentionally broken.
**The Discovery:** An indentation error in the YAML file caused the `ShellCheck` step to be ignored by GitHub Actions.
**The Fix:** We correctly aligned the steps, triggering the first successful "Build Failure"—a milestone in DevOps safety.

### 2. The Windows vs. Linux Showdown (`\r` Carriage Returns)

**The Error:** `SC1017: The parser reached the end of the file while looking for a corresponding 'done'.`
**The Culprit:** Windows uses `CRLF` (`\r\n`) for line endings, while Linux uses `LF` (`\n`). ShellCheck flagged the `\r` as a syntax error.
**The Fix:** We used the "In-place" stream editor to strip the carriage returns:

```bash
sed -i 's/\r$//' projects/06-linux-bash-basics/*.sh

```

### 3. The "Empty File" Crisis & Git Recovery

**The Issue:** During a fix, we accidentally ran a redirection command (`>`) that wiped the contents of `read_json.sh`.
**The Recovery:** We utilized Git as a "Time Machine" to restore the lost code:

```bash
git checkout projects/06-linux-bash-basics/read_json.sh

```

### 4. Word Splitting & Variable Quoting (SC2086)

**The Error:** ShellCheck flagged unquoted variables: `echo $RESPONSE`.
**The Lesson:** In Bash, unquoted variables are subject to "Word Splitting." If an API returns a string with spaces, the script breaks.
**The Fix:** Wrapped all variables in double quotes: `echo "$RESPONSE"`.

---

## 🌐 The Networking Breakthrough

One of the most complex hurdles was enabling a Docker container to talk to a Go API running in a WSL2 environment.

### The Problem: The "Loopback" Trap

A container trying to reach `localhost:8080` fails because `localhost` refers to itself.

### The Solution: Cross-Environment Bridging

1. **Go API Listener**: We ensured `main.go` was listening on `0.0.0.0` (all interfaces) rather than just `127.0.0.1`.
```go
http.ListenAndServe(":"+port, nil)

```


2. **IP Injection**: We identified the WSL2 IP (`hostname -I`) and injected it into the container via Environment Variables.
3. **Docker Host Gateway**: Used the `--add-host` flag to bridge the gap.

---

## 🚀 How to Run the Ecosystem

### 1. Start the Go API (Host Machine)

```bash
cd projects/06-linux-bash-basics
go run main.go

```

### 2. Run the Containerized Auditor

To test the API from within Docker, run:

```bash
WSL_IP=$(hostname -I | awk '{print $1}')
docker run --rm \
  -e API_URL="http://$WSL_IP:8080/health" \
  musabalaaudu/health-api:latest \
  /bin/bash /app/check_status.sh

```

---

## 🏆 Key DevOps Takeaways

* **Green Pipelines are Earned**: A passing build means nothing if the tests aren't actually running.
* **Linting is Non-Negotiable**: ShellCheck catches bugs that would only appear in production.
* **Environment Parity**: Always account for the differences between Windows (Host) and Linux (Docker/WSL).

---

**Current Status**: All pipelines are **Green**. Docker Image is **Verified**.

