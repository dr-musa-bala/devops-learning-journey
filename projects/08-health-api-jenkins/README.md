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

# Troubleshooting & Lessons Learned
🛠️ Troubleshooting & Engineering Logs
During the deployment of the health-api via Jenkins, several architectural and configuration challenges were encountered and resolved.

1. The "Command Not Found" (Path Disparity)
The Error: Jenkins console reported minikube: command not found.

The Cause: The minikube binary was installed in a local user directory (/home/dr-musa/.local/bin), which was not in the jenkins system user's $PATH.

The Resolution: * Migrated the binary to a universal location: sudo cp /home/dr-musa/.local/bin/minikube /usr/local/bin/minikube.

Explicitly defined the PATH within the Jenkinsfile environment block to ensure environment parity between the developer shell and the automation agent.

2. Groovy Compilation Error (Syntax Conflict)
The Error: unexpected char: '#' at the Kubernetes Deployment stage.

The Cause: Attempted to use Bash-style comments (#) outside of a shell (sh) block. Jenkins pipelines use Groovy, which requires // for comments.

The Resolution: Refactored the Jenkinsfile to use standard Groovy comment syntax for script logic and reserved # strictly for multi-line shell scripts.

3. The "Ghost" Deployment (Namespace Isolation)
The Error: Pipeline finished successfully, but kubectl get pods -n monitoring showed no new pods.

The Cause: The deployment defaults to the default namespace if not specified. The pods were being created, but in the wrong "room."

The Resolution: * Updated all kubectl commands in the pipeline with the explicit -n monitoring flag.

Standardized the deployment target to ensure the API sits alongside the Prometheus/Grafana stack for better observability.

4. NodePort Resource Contention
The Error: spec.ports[0].nodePort: Invalid value: 30080: provided port is already allocated.

The Cause: A previous service instance in the default namespace was still squatting on the physical port 30080.

The Resolution: * Identified the conflict using kubectl get svc --all-namespaces.

Performed a manual eviction of the stale service: kubectl delete svc health-api-service -n default.

Re-ran the pipeline to allow the new service to claim the port in the monitoring namespace.


## 🛡️ Resilience & Resource Governance Documentation

### 1. Health Monitoring (Liveness & Readiness Probes)
To ensure the application is self-healing and high-performing, we implemented Kubernetes Probes. These allow the cluster to monitor the actual "health" of the Flask process rather than just the state of the container.

#### **A. Readiness Probe**
* **Purpose:** Determines if the pod is ready to accept client traffic.
* **Mechanism:** Checks the `/health` endpoint on Port 80 every 5 seconds.
* **Impact:** If the probe fails (e.g., during a slow startup or database reconnection), Kubernetes removes the pod from the Service load balancer so users don't see 500 errors.

#### **B. Liveness Probe**
* **Purpose:** Determines if the pod is "alive" or stuck in a deadlocked state.
* **Mechanism:** Checks the `/health` endpoint on Port 80 every 10 seconds (after an initial 20-second delay).
* **Impact:** If the app freezes, Kubernetes automatically kills and restarts the pod to restore service.



---

### 2. Resource Quotas (Requests & Limits)
To prevent "Noisy Neighbor" syndrome and ensure cluster stability, we implemented strict resource boundaries.

| Metric | Request (Guaranteed) | Limit (Maximum) |
| :--- | :--- | :--- |
| **CPU** | 100m (0.1 Core) | 200m (0.2 Core) |
| **Memory** | 64Mi | 128Mi |

* **Logic:** By setting these values, we ensure the `monitoring` namespace remains stable. If the Health API encounters a memory leak, it will be capped at 128Mi and restarted (OOMKilled) before it can impact other services like Grafana or Prometheus.

---

### 3. Engineering Challenges & Resolutions

#### **Challenge A: The Port 80 vs 5000 Ambiguity**
* **Issue:** Standard Flask applications run on port 5000, but our probes were targeting port 80.
* **Diagnostic:** Checked internal logs using `kubectl logs [POD_NAME]`.
* **Resolution:** Confirmed the application was explicitly bound to `0.0.0.0:80` within the code/Dockerfile. We synchronized the probe configuration to target Port 80 to ensure the "Manager" was knocking on the correct "Chef's" door.

#### **Challenge B: Minimalist Image Restrictions (The "No-Tool" Container)**
* **Issue:** Attempted to simulate a crash using `pkill` and `kill` via `kubectl exec`, but both commands were missing from the `$PATH`.
* **Diagnostic:** The image is based on a "slim" distribution for security, which lacks standard process-management utilities.
* **Resolution:** Conducted "Chaos Engineering" from the control-plane level by deleting the pod manually. This successfully verified the **ReplicaSet** logic, as a new pod was instantly provisioned to maintain the desired state.

#### **Challenge C: Probe Delay Calibration**
* **Issue:** Pods initially reporting "Unhealthy" during the first few seconds of deployment.
* **Resolution:** Adjusted `initialDelaySeconds` for the Liveness probe to 20s to allow the Python interpreter and Flask server enough time to fully initialize before Kubernetes begins monitoring.

---
This troubleshooting phase is a vital part of your project history. It marks the transition from "it works on my machine" to "it works under the strict rules of Kubernetes." 

Here is the comprehensive documentation for the **Port Conflict** and **ConfigMap Integration** phase.

---

## 🛠️ Troubleshooting & Configuration Decoupling Log

### 1. The "Port Mismatch" Crisis
**Issue:** After implementing health probes, the application entered a `CrashLoopBackOff` state. Jenkins reported that the deployment "exceeded its progress deadline."

**Diagnostic Steps:**
* **Command:** `kubectl describe pod [POD_NAME] -n monitoring`
* **Finding:** The Events log showed `Liveness probe failed: connect: connection refused`.
* **The "Aha!" Moment:** The application logs confirmed the Flask app was listening on **Port 80**, but the Kubernetes manifest was instructing the cluster to check **Port 5000**. 

**Resolution:**
* Updated `k8s_deployment.yaml` to point both `livenessProbe` and `readinessProbe` to **Port 80**.
* Verified the fix by checking the pod status; the `Exit Code 137` (Forced Kill) disappeared and was replaced by a `Running` status.



---

### 2. Environment Decoupling (ConfigMaps)
To follow "Twelve-Factor App" best practices, we moved hardcoded environment variables out of the deployment logic and into a managed Kubernetes object.

**Implementation:**
* **Artifact:** Created `k8s_configmap.yaml` to store key-value pairs like `APP_LOCATION`.
* **Integration:** Used `valueFrom` and `configMapKeyRef` in the Deployment manifest to inject these variables at runtime.

**Benefit to the Team:**
* **Zero-Rebuild Config:** We can now change the application's "Location" or "Environment Name" by updating a single ConfigMap without needing to re-build the Docker image or modify the core Deployment code.

---

### 3. Key Challenges & Technical Lessons

| Challenge | Root Cause | Resolution |
| :--- | :--- | :--- |
| **CrashLoopBackOff** | Port Mismatch (App on 80, Probe on 5000) | Synchronized Probe ports to match Flask configuration. |
| **Exit Code 137** | Liveness Probe failure threshold met | Corrected the port so the probe received a `200 OK` heartbeat. |
| **`sh` Command Errors** | Confusion between Jenkins Pipeline syntax and Linux Shell | Removed `sh` prefix and quotes when running commands directly in the Ubuntu terminal. |

---

### 💡 Pro-Tip: The "Previous" Log Trick
One of the most valuable tools discovered during this phase was the `--previous` flag:
`kubectl logs [POD_NAME] --previous`

This allowed us to see exactly why the container crashed *before* it restarted, which is the only way to catch "silent" Python errors in a fast-cycling crash loop.

---

### 🏁 Final Verification Checklist
1. [x] Flask app listening on **Port 80**.
2. [x] Kubernetes Probes checking **Port 80**.
3. [x] `APP_LOCATION` pulled dynamically from **ConfigMap**.
4. [x] Resource limits set to **128Mi** to ensure cluster stability.
