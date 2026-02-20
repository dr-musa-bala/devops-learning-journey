# 🤝 Contributing to the DevOps Status API

First off, thank you for considering contributing! This tool was built to move our team from manual onboarding checks to a **"Digital Heartbeat"** architecture. 

## 🌟 Our Philosophy
We build tools that are:
1. **Lightweight:** Using the Go Standard Library only (no heavy frameworks).
2. **Observable:** Every feature must be reflectable via a JSON endpoint.
3. **Resilient:** Code must handle timeouts and system errors gracefully.

## 🛠 Low-Hanging Fruit (The Roadmap)
I have initialized the foundation. Here is where the team can help us iterate:

* **Automated Remediation:** If a health check fails, can we trigger a script to fix it?
* **Network Audits:** Adding connectivity checks for our private Docker registry or internal VPN.
* **Security:** Adding API Key validation or Basic Auth to protect the `/health` endpoint.
* **Custom Checks:** Add a new function in `main.go` to check for `kubectl`, `terraform`, or `aws-cli`.

## 🚀 How to Contribute
1. **Fork** the repository.
2. **Create a branch** for your feature (`git checkout -b feat/new-check`).
3. **Add your logic** in a new handler function.
4. **Test** locally: `go run main.go` then `curl http://localhost:8080/health`.
5. **Open a Pull Request** with a brief description of the impact.

---
*"DevOps is about making the right thing the easiest thing to do. Let's build it together."*
