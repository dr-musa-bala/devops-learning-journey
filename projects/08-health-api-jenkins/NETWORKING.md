# 🌐 Project 08: Ingress & Networking Documentation

## 1. Objective
To transition the **Health API** from a cluster-internal IP to a production-grade DNS-based routing system using the domain `health.kaduna.hub`.

## 2. Infrastructure Components
* **Controller:** NGINX Ingress Controller (Minikube Addon).
* **Resource:** Kubernetes Ingress Manifest (`k8s_ingress.yaml`).
* **Backend:** Flask API Service (`health-app-service`) on Port 80.
* **Metrics:** Prometheus Exporter integrated into the Flask app.

## 3. Implementation Workflow
The networking follows a **North-South** traffic pattern:
1.  **Client Request:** User hits `http://health.kaduna.hub:9876/health`.
2.  **Local DNS:** The Windows `hosts` file intercepts the domain and points it to the local bridge (`127.0.0.1`).
3.  **Port-Forwarding:** `kubectl` tunnels traffic from the local port `9876` to the Ingress Controller service.
4.  **Routing:** NGINX reads the `Host` header and forwards the request to the correct Pod based on path rules.



---

## 4. Challenges & Root Cause Analysis (RCA)

| Challenge | Root Cause | Rectification |
| :--- | :--- | :--- |
| **Connection Timeout** | Minikube's internal IP (`192.168.49.2`) was unreachable from the Windows host browser due to WSL2 networking isolation. | Implemented a **Port-Forward** bridge to `127.0.0.1` and updated the Windows `hosts` file to match. |
| **Address Already in Use** | Zombie processes from previous `kubectl` sessions were still holding onto ports (8080, 9090). | Used `sudo fuser -k [PORT]/tcp` to forcefully clear the port before restarting the tunnel. |
| **404 Not Found** | The Flask app was only configured for the `/health` route, but the browser was hitting the root `/`. | Verified routes via `curl` and confirmed the app was responding correctly on the defined endpoints (`/health` and `/metrics`). |
| **DNS Resolution Failure** | The terminal was trying to ping the old Minikube IP while the bridge was on localhost. | Updated `C:\Windows\System32\drivers\etc\hosts` to ensure the domain pointed to the active tunnel address. |

---

## 5. Verification Commands
To verify the health of the network stack, use the following suite of commands:

```bash
# 1. Check if the "Receptionist" (Ingress) is alive
kubectl get pods -n ingress-nginx

# 2. Check the Ingress rules and address
kubectl get ingress -n monitoring

# 3. Test the application endpoint
curl -i -H "Host: health.kaduna.hub" http://127.0.0.1:9876/health

# 4. Test the observability/metrics endpoint
curl -i -H "Host: health.kaduna.hub" http://127.0.0.1:9876/metrics
```

---

## 6. Lessons Learned
* **Logs are Truth:** Using `kubectl logs` allowed us to see that the 404 error was coming from the *App*, not the *Network*, which saved hours of troubleshooting.
* **Tunneling is Key in WSL:** Standard Ingress IPs often fail in virtualized environments; `port-forward` is a reliable fallback for development.
* **Clean Exits:** Always use `Ctrl+C` to close tunnels to avoid port-binding issues in future sessions.

