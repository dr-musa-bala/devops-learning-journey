# Contributing to the Status API

This service is a baseline for our team's observability. We follow the "Standard Library First" principle to keep our DevOps tools lightweight and secure.

## 🛠 How to Help
We are currently looking to iterate on the following:
1. **New Endpoints:** Add checks for specific databases or cloud providers.
2. **Self-Healing:** Help us write the logic to trigger a service restart if a health check fails.
3. **Security:** Implement Basic Auth or API Key validation in the headers.

## 🚀 Development Workflow
1. **Fork** the repo.
2. **Add your logic** in a new handler function in `main.go`.
3. **Register the route** in the `main()` function.
4. **Test** using `curl http://localhost:8080/your-endpoint`.

*"Automation is a team sport. Let's build together."*
