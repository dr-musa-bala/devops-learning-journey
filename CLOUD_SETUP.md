# 🚀 Cloud-Native Health API (Go + Redis)

A high-performance microservice deployed using a modern, zero-cost DevOps pipeline. This project demonstrates containerization, automated CI/CD, and cloud-native service orchestration.

## 🏗️ Architecture
- **Language:** Go (Golang)
- **Containerization:** Docker & Docker Hub
- **Database:** Upstash (Serverless Redis)
- **Deployment:** Render (Platform-as-a-Service)
- **CI/CD:** GitHub Actions (Automated Builds)

## 🌐 Live Demo
The API is live and reachable at: 
[https://health-api-latest.onrender.com/health](https://health-api-latest.onrender.com/health)

## 🛠️ Infrastructure Setup
This project follows a "Zero-Bill" strategy to maintain industry standards without AWS costs:

1. **Local VM:** Developed and tested inside an Ubuntu VirtualBox environment.
2. **GitHub Actions:** Automatically builds a Docker image on every `push` and sends it to Docker Hub.
3. **Upstash Redis:** Provides a managed Redis instance with SSL encryption.
4. **Render:** Pulls the image from Docker Hub and injects `REDIS_ADDR` and `PORT` via Environment Variables.

## 🚀 How to Run Locally
```bash
docker-compose up -d
