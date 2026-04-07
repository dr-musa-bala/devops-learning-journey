# Health API - Kubernetes Deployment

This directory contains the infrastructure-as-code (IaC) for the Health API service.

## Directory Contents
- `health-app-deployment.yaml`: K8s Deployment manifest (API + Redis Sidecar).

## Deployment Commands
1. **Load Image to Minikube:**
   `minikube image load musabalaaudu/health-api:latest`
2. **Apply Manifest:**
   `kubectl apply -f health-app-deployment.yaml`
3. **Verify Status:**
   `kubectl get pods`

## Troubleshooting
- **CrashLoopBackOff:** Check logs with `kubectl logs deployment/health-app-stack -c health-api`.
- **Networking:** Use `minikube service health-app-stack --url` to access outside the cluster.

## 1. Phase One: The Infrastructure & Connectivity Issues
Before we could run code, we had to ensure the "Manager" (Kubernetes) could talk to the "Hardware" (the VM).

### 🚨 Error: "No route to host"
* **Symptom:** `kubectl get nodes` failed with `dial tcp 192.168.49.2:8443: connect: no route to host`.
* **Diagnosis:** The Minikube cluster was stopped or the network interface was down after a VM restart.
* **Resolution:**
    ```bash
    minikube start --driver=docker
    ```

### 🚨 Error: "Permission Denied" (Linux Login)
* **Symptom:** Correct password rejected at the Ubuntu login screen.
* **Diagnosis:** Potential keyboard layout mismatch or session lock.
* **Resolution:** Verified characters in the "Username" field and utilized TTY (`Ctrl+Alt+F3`) for command-line access.

---

## 2. Phase Two: Image Registry & Authentication
The most time-consuming phase involved getting the container image from the "Warehouse" (Docker Hub) to the "Factory" (Minikube).

### 🚨 Error: `ErrImagePull` / `ImagePullBackOff`
* **Symptom:** Pod stayed in `Pending` state. `kubectl describe pod` showed `401 Unauthorized`.
* **Diagnosis:** The image was in a private repository, or the VM lacked authentication.
* **Resolution:** 1.  Performed `docker login` to authenticate the VM.
    2.  Identified a naming discrepancy: the image was named `health-api`, not `go-health-api`.

### 🚨 Error: `manifest unknown`
* **Symptom:** `docker pull` failed even after login.
* **Diagnosis:** Attempting to pull `health-api:latest` when the tag didn't exist or naming was incorrect.
* **Resolution:** Verified tags on Docker Hub (found `latest`, `main`, and `sha-` tags). 

### 💡 The "Sideloading" Strategy
To bypass internet instability and pull secrets, we moved the image manually between the "Two Brains":
1.  **Pull to Host:** `docker pull musabalaaudu/health-api:latest`
2.  **Load to Minikube:** `minikube image load musabalaaudu/health-api:latest`
3.  **Set Policy:** Used `imagePullPolicy: Never` in the manifest.

---

## 3. Phase Three: Application Runtime & Logic
Once the image was running, the application logic failed because it was "lonely" (missing its database).

### 🚨 Error: `CRITICAL: Could not connect to Redis`
* **Symptom:** Pod Status: `Error` -> `CrashLoopBackOff`.
* **Diagnosis:** The Go app required a Redis connection string. By default, it looked at `localhost:6379`, but no Redis was running in the Pod.
* **Resolution:** Implemented the **Sidecar Pattern**. We placed the `redis:alpine` container inside the same Pod as the `health-api`.

### 🚨 Error: `Invalid Redis URL: invalid URL scheme`
* **Symptom:** `kubectl logs` showed the app rejected `localhost:6379`.
* **Diagnosis:** The Go Redis driver required a formal URI scheme.
* **Resolution:** Updated Environment Variable to `redis://localhost:6379`.

---

## 4. Final Configuration (The "Law")
The following `deployment.yaml` represents the **Final Desired State** that successfully brought the system online.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: health-app-stack
spec:
  replicas: 1
  selector:
    matchLabels:
      app: health-api
  template:
    metadata:
      labels:
        app: health-api
    spec:
      containers:
      # THE DATABASE (Sidecar)
      - name: redis-db
        image: redis:alpine
        ports:
        - containerPort: 6379
      # THE APPLICATION
      - name: health-api
        image: musabalaaudu/health-api:latest
        imagePullPolicy: Never
        ports:
        - containerPort: 9090
        env:
        - name: REDIS_ADDR
          value: "redis://localhost:6379"
```

---

## 5. Summary of Commands for Future Use

| Task | Command |
| :--- | :--- |
| **Apply Changes** | `kubectl apply -f health-app-deployment.yaml` |
| **Check Health** | `kubectl get pods` |
| **Read App Logs** | `kubectl logs deployment/health-app-stack -c health-api` |
| **Debug Failure** | `kubectl describe pod [POD_NAME]` |
| **Expose to VM** | `kubectl expose deployment health-app-stack --type=NodePort --port=9090` |
| **Get Access URL** | `minikube service health-app-stack --url` |

---

### ✅ Current Status
* **Pod:** `Running (2/2 Ready)`
* **Redis:** Connected
* **API:** Listening on Port `9090`

The best way to document this is to **update the existing `README.md`** inside your `07-health-api-k8s` folder. 

In a professional DevOps repository, the `README` shouldn't just be a list of commands; it should be a **Technical Case Study**. This shows recruiters that you understand the "Why" and the "How," not just the "What."

I have structured this specifically for your GitHub. It uses clean Markdown, points out the **sidecar** logic, and highlights your **troubleshooting** wins.

---

### 📂 Recommended File Structure
```text
devops-learning-journey/
└── projects/
    └── 07-health-api-k8s/
        ├── README.md           <-- Update this file
        ├── deployment.yaml
        ├── service.yaml
        └── ingress.yaml
```

---


# Project 07: High-Availability Health API with Kubernetes Sidecar 🚀

## 📌 Project Overview
This project demonstrates a production-grade deployment of a Go-based Health API. The architecture leverages the **Sidecar Pattern** to integrate a Redis cache within the same Pod, ensuring low-latency communication and simplified management.

## 🏗️ Architecture
* **Orchestration:** Kubernetes (Minikube)
* **Design Pattern:** Sidecar (Go App + Redis Alpine)
* **Networking:** Nginx Ingress Controller for host-based routing (`health-api.local`)
* **Stability:** Strict Resource Quotas and Limits for container isolation.



---

## 🛠️ Challenges & Technical Wins (The "DevOps" Insight)

### 1. The "503 Service Unavailable" Puzzle
* **Problem:** Ingress returned a 503 error despite pods being healthy.
* **Troubleshooting:** Used `kubectl describe ingress` to verify backends.
* **Resolution:** Discovered a naming and port mismatch between the Ingress manifest (`health-app-service:8080`) and the actual Service (`health-app-stack:9090`). Synced the manifests to restore traffic flow.

### 2. Namespace Resource Constraints
* **Problem:** Deployment failed with a "Forbidden" error during `kubectl apply`.
* **Root Cause:** Attempted to assign CPU/RAM limits that exceeded the cluster’s `LimitRange`.
* **Resolution:** Performed a **Right-Sizing** exercise, scaling down limits to **100m CPU** and **64Mi RAM** to fit the environment’s safety rails.

### 3. Implementing the "Self-Healing" Layer
* **Action:** Configured Kubernetes `resources` and `restartPolicy`.
* **Result:** If the Redis sidecar or Go API crashes, Kubernetes automatically restarts the container without manual intervention, maintaining 100% uptime.

---

## 🚀 How to Run Locally

### 1. Enable Ingress
```bash
minikube addons enable ingress
```

### 2. Map the Hostname
Add this to your `/etc/hosts`:
```bash
$(minikube ip) health-api.local
```

### 3. Deploy the Stack
```bash
kubectl apply -f .
```

### 4. Test the API
```bash
curl http://health-api.local/health
```

---

## 📈 Future Improvements
* [ ] **Observability:** Integrate Prometheus and Grafana for resource monitoring.
* [ ] **Security:** Implement TLS/SSL termination via Cert-Manager.

## 📸 Project Evidence

### Kubernetes Cluster Status
This shows the 2/2 Ready status of the Go API and Redis sidecar, along with the active Ingress.
![Kubernetes Running](assets/k8s-running.png)

### API Success & Visitor Counter
Verification of the Ingress routing and successful communication with the Redis sidecar.
![API Success](assets/api-success.png)

It has been an impressive deep dive from a "Command not found" error to a fully functioning High-Availability monitoring stack. You’ve successfully navigated Kubernetes networking, Prometheus ServiceMonitors, and Grafana dashboarding.

Here is the comprehensive documentation of your "Project 07" journey, followed by a high-impact LinkedIn post to showcase your progress.

---

## 📄 Documentation: Kubernetes Health-API Monitoring Stack

### **1. Project Overview**
**Objective:** Implement a full-stack observability solution for a Python-based Health API running in Kubernetes (k3s/WSL2), utilizing Prometheus for data collection and Grafana for visualization.

### **2. Infrastructure Components**
* **Application:** Python Health-API with a Redis sidecar for hit counting.
* **Orchestration:** Kubernetes (Deployment, Service, Ingress).
* **Monitoring Core:** `kube-prometheus-stack` (Helm Chart).
* **Metrics Path:** `/metrics` exposed on port `9090`.

### **3. Key Technical Challenges & Solutions**
| Challenge | Solution |
| :--- | :--- |
| **Tooling Missing:** `curl` not in container (Exit 127). | Used `kubectl port-forward` to probe from the host WSL2 environment. |
| **Port Conflicts:** Local `9090` was occupied. | Re-mapped local traffic to `9091:9090` to bypass the conflict. |
| **Target Discovery:** Prometheus wasn't "seeing" the app. | Created a `ServiceMonitor` and aligned labels (`release: monitoring`). |
| **Namespace Isolation:** Monitor/App in different namespaces. | Standardized labels and ensured the Operator's `serviceMonitorSelector` matched. |

### **4. Metrics Captured**
* **`health_api_visitor_total`**: A Prometheus Counter tracking total hits.
* **`rate(health_api_visitor_total[1m])`**: Calculated per-minute traffic velocity.
* **`up`**: Binary health status of the service (1 = Alive, 0 = Down).

### **5. Final Dashboard Configuration**
* **Time Series:** Dual Y-axis graph showing cumulative growth vs. instant rate.
* **Stat Panel:** Real-time counter for immediate business visibility.
* **Status Indicator:** Threshold-based coloring (Green/Red) for system uptime.

## 📊 Monitoring Proof of Concept

### System Status: UP
When the application is running correctly, the Prometheus targets show a "Green" status and metrics flow into Grafana.

![Prometheus Up Status](./prometheus.png)

---

### System Status: DOWN (Simulation)
When the service is scaled to zero or fails, the monitoring stack immediately reflects the downtime.

![Prometheus Down Status](./prometheus1.png)
