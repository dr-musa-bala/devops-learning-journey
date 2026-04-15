## 📄 Infrastructure Documentation: Persistence & Auto-Scaling
**Project:** Health-API (Phase 4 & 5)  
**Environment:** Minikube (Development/Testing)  
**Namespace:** `monitoring`

---

### 1. Stateful Storage (Persistent Volumes)
To transition the Health-API from a "stateless" app to a "stateful" one, we implemented a decoupled storage architecture. This ensures that logs, audit trails, and local data survive pod restarts or deployments.

#### **Architecture Components:**
* **PersistentVolume (PV):** A cluster-wide resource representing a 1Gi segment of the host machine's disk (`/mnt/data`).
* **PersistentVolumeClaim (PVC):** A request by the application for a specific amount of storage (1Gi) with `ReadWriteOnce` access.
* **Volume Mount:** The internal mapping that connects the PVC to the container's `/app/data` directory.



#### **Challenges & Resolutions:**
| Challenge | Root Cause | Resolution |
| :--- | :--- | :--- |
| **ImagePullBackOff** | Using `:latest` tag triggered an external pull from Docker Hub, which failed because the image only existed in the local Minikube registry. | Updated `k8s_deployment.yaml` to use a specific, immutable build tag (`health-api:31`) and set `imagePullPolicy: IfNotPresent`. |
| **Stuck Deployment** | Kubernetes was trying to perform a Rolling Update but could not start the new Pod due to the image error. | Manually deleted the failing Pod and updated the Deployment manifest to clear the "stuck" state. |
| **Data Verification** | Uncertainty if the storage was actually persistent across Pod lifecycles. | Performed a "Chaos Test": Created a file in the Pod, deleted the Pod, and verified the file existed in the newly spawned replacement Pod. |

### 💾 Data Persistence
The Health-API uses a Persistent Volume Claim (PVC) to ensure data survives pod restarts. 

![PVC Bound Status](./images/pvc-bound.png)
*Figure 1: Verification of the PVC successfully binding to the PV.*
---

### 2. Elastic Scaling (Horizontal Pod Autoscaler - HPA)
To handle traffic spikes in the **Kaduna Hub** and ensure high availability, we moved from static replicas to dynamic scaling.

#### **Configuration Details:**
* **Minimum Replicas:** 2 (Ensures zero-downtime during updates).
* **Maximum Replicas:** 5 (Prevents resource exhaustion on the host).
* **Scaling Trigger:** 50% average CPU utilization.



#### **Operational Logic:**
1.  **Metrics Collection:** The HPA queries the `metrics-server` every 15 seconds.
2.  **Calculation:** If `CurrentCPU > 50%`, the HPA calculates the required replicas using the formula:
    $$DesiredReplicas = \lceil CurrentReplicas \times \frac{CurrentMetricValue}{TargetValue} \rceil$$
3.  **Cooldown (Stabilization):** To prevent "flapping" (scaling up and down too rapidly), Kubernetes waits for a stabilization window (usually 5 minutes) before scaling back down.

---

### 3. Benefits to the Team
* **Operational Resilience:** The app now self-heals without losing data.
* **Cost & Resource Efficiency:** We only use the compute power we need, scaling down during low-traffic periods in the Kaduna region.
* **Immutable Deployments:** By switching from `:latest` to specific tags (e.g., `:31`), the support team can rollback to a known working state instantly if a bug is discovered.

---

### 🛠️ Final Infrastructure State Check
To view the current "Health" of this documentation's implementation, run:
```bash
# Check all resources in the namespace
kubectl get all,pv,pvc,hpa -n monitoring
```
### 📈 Horizontal Pod Autoscaling (HPA)
We implemented HPA to manage traffic spikes. The system automatically scales between 2 and 5 replicas based on a 50% CPU threshold.

![HPA Target Status](./images/hpa-status.png)
*Figure 2: HPA monitoring the live CPU load of the Kaduna Hub.*

Effective documentation is the final step in any successful DevOps sprint. This document serves as the "Technical SOP" (Standard Operating Procedure) for how secrets are handled within the **Health-API** project in the Kaduna axis.

---

## 🔐 Infrastructure Documentation: Secrets Management
**Project:** Health-API (Phase 6)  
**Security Level:** Protected (Base64 Encoded / Imperative Management)  
**Namespace:** `monitoring`

---

### 1. The Strategy: Imperative vs. Declarative
To maintain a "Zero-Credentials-in-Git" policy, we have adopted an **Imperative Secret Management** strategy. 

* **Declarative (YAML):** Used for non-sensitive configurations (ConfigMaps). Safe for GitHub.
* **Imperative (Command Line):** Used for sensitive credentials (Secrets). **Forbidden** in GitHub. This ensures that even if our repository is compromised, the actual production database passwords remain secure inside the cluster.



---

### 2. Implementation Details

#### **A. The Secret Object**
We created a Kubernetes Secret of type `Opaque`. Unlike a ConfigMap, the values here are stored in Base64 encoding to prevent accidental "shoulder-surfing" exposure.

* **Secret Name:** `health-api-secrets`
* **Data Key:** `DB_PASSWORD`
* **Format:** Base64 Encoded string.

#### **B. The Injection Pattern**
The secret is injected into the Pod at runtime as an **Environment Variable**. This allows the application to consume the credential without needing to know that it is stored in Kubernetes.

**Deployment Snippet:**
```yaml
env:
  - name: DATABASE_PASSWORD
    valueFrom:
      secretKeyRef:
        name: health-api-secrets
        key: DB_PASSWORD
```

---

### 3. Challenges & Security Resolutions

| Challenge | Risk | Resolution |
| :--- | :--- | :--- |
| **Git Leakage** | Committing `k8s_secrets.yaml` would expose passwords to anyone with repo access. | **Deleted** the YAML file. Added `k8s_secrets.yaml` to `.gitignore`. |
| **Encoding vs. Encryption** | Base64 is not true encryption; it is easily decodable. | Implemented the **Imperative Command** method to keep secrets in cluster memory only. |
| **Pod Environment Exposure** | Root users can see env vars via `printenv`. | (Future Step) Discussed transitioning to **CSI Secret Store Drivers** or **Vault** for memory-only injection. |

---

### 4. Operational Instructions (For the Support Team)

Since the secret is not in the Git repository, any engineer setting up this environment must manually create the secret using the following command:

**To Create:**
```bash
kubectl create secret generic health-api-secrets \
  --from-literal=DB_PASSWORD='<your_secure_password>' \
  -n monitoring
```

**To Verify (Security Audit):**
```bash
# Check if the secret exists
kubectl get secrets -n monitoring health-api-secrets

# Verify the value (Internal use only)
kubectl get secret health-api-secrets -n monitoring -o jsonpath='{.data.DB_PASSWORD}' | base64 --decode
```

---

### 🏆 Milestone Verification
The implementation is confirmed successful. 
* **Status:** `health-api` Pods are in `Running` state.
* **Verification:** Running `printenv | grep DATABASE_PASSWORD` inside the pod returns the correct plain-text password, proving the Kubernetes "Decoupling" logic is functioning perfectly.

---
