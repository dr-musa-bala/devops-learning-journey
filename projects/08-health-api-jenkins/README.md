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

To ensure your **Project 08 README** is a complete "Master Class" in troubleshooting, you should add these final three phases. These cover the specific "hidden" errors we solved regarding branch naming, nested paths, and the Docker environment mismatch.

### 5. Phase V: Source Control & Branch Synchronization
**Objective:** Transition from local testing to a GitHub-driven "Source of Truth."
* **🚩 The "Master vs. Main" Twist:**
    * **Challenge:** Pipeline failed with `GitException: Status code 128` during fetch.
    * **The Cause:** Modern GitHub repositories use `main` as the default branch, but Jenkins legacy settings default to `master`.
    * **The Rectification:** Explicitly updated the **Branch Specifier** to `*/main` in the Jenkins Job Configuration.
* **🚩 The "Nested Path" Twist:**
    * **Challenge:** Jenkins could not find the `Jenkinsfile`.
    * **The Cause:** The file was located in a sub-directory (`projects/08-health-api-jenkins/`) rather than the repository root.
    * **The Rectification:** Configured the **Script Path** in Jenkins to point to the specific sub-directory path.

### 6. Phase VI: The Environment Sync (Docker-in-K8s)
**Objective:** Resolve the "Image Visibility" gap between the builder and the orchestrator.
* **🚩 The Container Boundary Twist:**
    * **Challenge:** Pods stuck in `ErrImageNeverPull` status.
    * **The Cause:** Jenkins built the Docker image in the **WSL2 host daemon**, but Minikube’s Kubelet was searching its own **isolated internal registry**.
    * **The Rectification:** Modified the `Jenkinsfile` to include `dir()` blocks for context and injected `minikube docker-env` variables. By manually pointing `DOCKER_CERT_PATH` to the migrated certificates in `/var/lib/jenkins/`, we forced the build to happen **inside** the Minikube node.

### 7. Phase VII: The Network Bridge (Final Validation)
**Objective:** Confirm the API is reachable and reporting correct metadata from the "Kaduna Axis."
* **🚩 The Network Island Twist:**
    * **Challenge:** `curl` timed out when hitting the Minikube IP directly from WSL.
    * **The Cause:** Network isolation between the WSL2 virtual bridge and the Minikube Docker-driver network.
    * **The Rectification:** Implemented a **Minikube Tunnel** (`minikube service health-api-service`) to create a loopback bridge to `127.0.0.1`.
* **Verification:** Confirmed the API correctly identifies its hardcoded deployment context:
    ```json
    {
      "location": "Kaduna-Axis-DC",
      "status": "UP",
      "system": "Kubernetes-WSL"
    }
    ```
Project Name: Health API - Dynamic Infrastructure MigrationObjective: Transition from static, hardcoded configurations to a dynamic, environment-aware CI/CD pipeline.1. Technical ArchitectureApplication: Flask-based Health API.Infrastructure: Kubernetes (Minikube) on WSL2.CI/CD: Jenkins (Automated Pipeline).Communication: Port 80 (Standard HTTP).2. Key TransformationsFeatureOld State (Phase 08)New State (Phase 09)Location LogicHardcoded in app.pyInjected via K8s env variablesNetwork PortPort 5000 (Development)Port 80 (Production Standard)ImportsMissing datetimeRobust os and datetime integrationDeploymentManual overwriteJenkins-managed Rolling Update3. Environment Variable MappingThe application now uses os.getenv("APP_LOCATION"). This allows the same container image to be deployed across different regions (e.g., Kaduna-Axis, Lagos, or Global Cloud) by simply changing the value in the k8s_deployment.yaml.

Project Name: Health API - Infrastructure Hygiene & Self-Healing

Status: Operational / Optimized

1. Resource Optimization (Storage)
Challenge: Continuous CI/CD builds were creating "Dangling Images" and filling the Build Cache, consuming 9.7GB of local storage.

Solution: * Implemented an automated post-build cleanup in Jenkins using docker image prune -f.

Performed a deep system purge using docker system prune -a --volumes to clear the legacy build cache.

Result: Reclaimed ~7.2GB of disk space; reduced total image footprint from 9.7GB to 2.4GB.

2. Self-Healing Infrastructure
Feature: Added Kubernetes Liveness and Readiness Probes.

Logic:

Readiness: Ensures the API is fully loaded before allowing the Service to send traffic.

Liveness: Automatically restarts the container if the /health endpoint stops responding.

Benefit: Zero-downtime deployments and automated recovery from application hangs.
