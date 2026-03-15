## **Project Documentation: Proactive Observability & ChatOps Alerting**

**System:** Go-Health API Infrastructure

**Core Stack:** Prometheus, Alertmanager, Discord Webhooks

### **1. Overview**

To move from a reactive "ticket-based" administration to a proactive "reliability-first" model, I implemented a centralized alerting pipeline. The goal was to ensure zero blind spots in the containerized environment, specifically targeting service availability and resource saturation. 

### **2. Architecture & Implementation**

The alerting loop follows a three-stage process:

* 
**Metric Collection:** Prometheus scrapes the `/metrics` endpoint of the Go-Health API every 15 seconds. 


* 
**Evaluation:** Custom `alert.rules` files define the logic for "Critical" vs. "Warning" states. 


* 
**Notification:** Alertmanager handles grouping, silencing, and routing the alerts to a dedicated Discord "DevOps-Incidents" channel via a Webhook. 



#### **Key Alerting Rules Defined:**

* **InstanceDown:** Triggered if the scraper cannot reach the Go API for >30 seconds.
* **HighLatency:** Triggered if the 95th percentile of request duration exceeds 500ms.
* 
**RedisConnectionLost:** Specifically monitors the connectivity between the API and the stateful layer. 



---

### **3. Challenges Encountered & Solutions**

#### **Challenge A: The "Network Partition" Silence**

* 
**Problem:** During initial testing, a network bridge misconfiguration caused the API to lose its route to the scraper.  Because the scraper couldn't see the API, it didn't register a "High Latency" alert—it simply saw nothing.


* **Solution:** Implemented an **"Absence Alert"** (`absent(up{job="go-api"})`). This ensures that if the system goes silent, an alert is fired for "Missing Data" rather than assuming everything is fine. 



#### **Challenge B: Alert Fatigue (Noise)**

* **Problem:** Standard flapping (a service being down for 1 second during a restart) was triggering "Critical" alerts in Discord, creating unnecessary noise.
* **Solution:** Introduced a `for: 1m` buffer in the Alertmanager configuration. This requires the failure state to persist for a full minute before notifying a human, allowing for transient network blips to resolve themselves.

---

### **4. Business Impact**

* **Mean Time to Detect (MTTD):** Reduced from "User Reporting" (minutes/hours) to <60 seconds.
* 
**System Visibility:** Real-time visibility into container health without manual CLI checks. 


* **Industrialized Reliability:** Established a repeatable pattern that can be scaled to the other 50+ SaaS products in the portfolio.

## Monitoring Workflow

| Status | Screenshot |
|------|------|
| Grafana Healthy | ![](docs/screenshots/healthy_grafana.png) |
| Grafana Unhealthy | ![](docs/screenshots/unhealthy_grafana.png) |
| Prometheus Down | ![](docs/screenshots/prometheus_down.png) |
| Prometheus Up | ![](docs/screenshots/prometheus_up.png) |
| Discord Alert Resolved | ![](docs/screenshots/resolved_discord.png) |
