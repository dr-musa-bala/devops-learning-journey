# 📑 Project 08: Operational Report
## Automated CI/CD Integration for the Health API
**Engineer:** Dr. Musa  
**System Date:** April 2026  
**Status:** **FULLY OPERATIONAL & DEPLOYED**

---

## 1. Project Objective
The primary objective of this project was to establish a fully automated **Continuous Integration and Continuous Deployment (CI/CD) Pipeline** on a local WSL2 environment. This pipeline is designed to:
* **Validate** source code through automated checks.
* **Containerize** a Python Flask "Health API" using Docker.
* **Orchestrate** the deployment of that containerized application into a Kubernetes (Minikube) cluster.
* **Ensure Security** by using scoped keys and the Principle of Least Privilege for service accounts.
* **Synchronize** with GitHub as the "Single Source of Truth."

---

## 2. Phase I: Environment Provisioning (The Foundation)
The first goal was to install the automation engine (Jenkins) directly onto the Linux subsystem (WSL2) to ensure a high-performance, native connection to the Docker daemon.

### The Sequential Procedure:
1.  **Dependency Installation:** Installed Java 17 (JRE) and Fontconfig.
2.  **Security Handshake:** Added the Jenkins repository keys to the system.

### 🚩 The First Twist: "The Invalid Guardian"
* **Challenge:** The standard installation commands failed; `apt` reported the Jenkins package was "obsolete."
* **The Cause:** In early 2026, Jenkins rotated their GPG signing keys. The system refused the repository because the signature didn't match the new standards.
* **The Rectification:** We performed a **Clean Reset**, implemented a **Scoped Keyring** at `/etc/apt/keyrings/`, and used the `signed-by` flag to explicitly trust the new 2026 key for that repository only.

---

## 3. Phase II: Service Integration (The Handshake)
Objective: Give Jenkins "eyes and ears" to communicate with Docker and Kubernetes.

### 🚩 The Second Twist: "The Permission Barrier"
* **Challenge:** Pipeline command `docker --version` failed with a permission error.
* **The Cause:** The `jenkins` service user lacked rights to talk to the Docker engine (root-level socket access).
* **The Rectification:** We performed a **Group Elevation** (`sudo usermod -aG docker jenkins`), allowing Jenkins to inherit Docker permissions safely.

---

## 4. Phase III: Kubernetes Authentication (The Final Boss)
Objective: Allow Jenkins to trigger deployments inside the Minikube cluster.

### 🚩 The Third Twist: "The Certificate Paradox"
* **Challenge:** Pipeline failed with `unable to read client-cert`.
* **The Cause:** The SSL certificates were stored in `/home/dr-musa/`. Linux security prevents the `jenkins` service account from reaching into a user's home directory.
* **The Rectification:** We executed a **Certificate Migration Strategy**. We relocated the `.minikube` directory to `/var/lib/jenkins/`, updated ownership, and used `sed` to programmatically update the `kubeconfig` paths to point to the new location.

---

## 5. Phase IV: Production Alignment & Flattening (The Final Push)
Objective: Move from a "Workspace Test" to a "Source of Truth" deployment using a GitHub repository.

### The Sequential Procedure:
1.  **Code Creation:** Developed the Flask `app.py`, `Dockerfile`, and requirements.
2.  **Infrastructure as Code:** Drafted Kubernetes manifests for Deployment and Service.

### 🚩 The Fourth Twist: "The Directory Labyrinth"
* **Challenge:** Creating sub-directories like `k8s/deployment.yaml` caused pathing errors in both the shell and the Jenkins Pipeline scripts.
* **The Cause:** Forward slashes in filenames and missing parent directories caused "No such file or directory" errors during automation.
* **The Rectification:** We adopted a **Flat File Strategy**. We renamed the manifests to `k8s_deployment.yaml` and `k8s_service.yaml` at the root level. This simplified the Jenkinsfile logic and made the repository "Self-Documenting."

---

## 6. Phase V: GitHub Synchronization & Automated Build
1.  **Identity Setup:** Configured Git global parameters and initialized the repo in `~/devops-learning-journey/projects/08-health-api-jenkins`.
2.  **Secure Push:** Utilized a **Personal Access Token (PAT)** to push to GitHub, bypassing the deprecated password authentication.
3.  **Webhook Simulation:** Configured Jenkins to use **Pipeline from SCM**, pointing it to the GitHub URL as the definitive source.

---

## 7. Final Result: Success & Verification
With the final Jenkinsfile executed, the following results were achieved:
* **Docker Build:** Image `health-api:latest` successfully created.
* **K8s Deployment:** Two replicas of the API are running in Minikube.
* **Network Access:** Service is exposed on NodePort `30080`.

**Console Output:**
```text
[Pipeline] stage (Kubernetes Deployment)
+ kubectl apply -f k8s_deployment.yaml
deployment.apps/health-api created
+ kubectl apply -f k8s_service.yaml
service/health-api-service created
Finished: SUCCESS
```

## 8. Conclusion
Project 08 successfully transitioned from a broken, permission-locked environment to a **Hardened, Automated Pipeline**. By resolving the twists of GPG keys, Linux permissions, certificate paths, and file structure, we have realized a production-ready CI/CD workflow.
