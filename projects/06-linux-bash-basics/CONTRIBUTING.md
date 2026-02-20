# 🤝 Contributing to the Observability Pipeline

First off, thank you for helping us build a more resilient system! This project bridge's the gap between **Go Microservices** and **Linux Automation**.

## 🛠 How You Can Contribute

We are looking to expand our monitoring capabilities in the following areas:

### 1. New Bash Probes
If you want to add a new system check:
* Create a new script named `check_<feature>.sh`.
* Ensure it uses `jq` for parsing if it talks to an API.
* Always include a "Silent" mode using `> /dev/null` to keep logs clean.

### 2. Go API Enhancements
Help us make the "Source of Truth" more detailed:
* Add new fields to the JSON response in `main.go` (e.g., `cpu_usage`, `uptime`).
* Ensure the `status` field remains a standard string for our Bash logic to read.

### 3. Automation (Crontab)
Help us optimize how often these scripts run. If you find a better scheduling frequency for production, open a Pull Request!

## 🚀 Submission Process

1. **Fork** the repository.
2. **Create a Branch**: `git checkout -b feat/your-improvement`.
3. **Test Your Script**: Ensure `chmod +x` is set and it runs without errors on Ubuntu.
4. **Update the README**: If you add a new tool, mention it in the metrics table.
5. **Open a PR**: Describe the problem your contribution solves.

## 📝 Coding Standards
* Use **ShellCheck** to validate your Bash scripts.
* Use the **Shebang** `#!/bin/bash` at the top of every file.
* Variables should be `UPPERCASE` and quoted (e.g., `"$VAR"`).

---
*“Observability is about making the invisible, visible.”*
