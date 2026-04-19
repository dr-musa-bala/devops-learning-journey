# 📑 DevOps Project 08: Post-Mortem & Technical Documentation
**Project:** Health API Automated CI/CD Pipeline  
**Engineer:** Musa Bala Audu  
**Environment:** WSL2 (Ubuntu), Jenkins (Service), Minikube (K8s)

---

## 🏗️ 1. Pipeline Overview
The goal was to automate the build, push, and deployment of a Python-based Health API. Every time code is pushed to GitHub, Jenkins builds a Docker image, pushes it to Docker Hub, and updates the Kubernetes deployment in the `monitoring` namespace.

---

## 🚧 2. Challenges & Resolutions

### Challenge A: The "Localhost" Loopback (Network Isolation)
* **The Symptom:** Jenkins failed the `Deploy` stage with `dial tcp 127.0.0.1:XXXXX: connect: connection refused`.
* **The Cause:** `127.0.0.1` inside a Jenkins process refers to the Jenkins environment itself, not the host machine where the Minikube API is actually listening.
* **The Rectification:** * Applied the `--validate=false` flag to the `kubectl apply` command.
    * This forces `kubectl` to send the manifest directly to the server without trying to reach the local discovery port.

### Challenge B: Kubernetes Namespace Mismatch
* **The Symptom:** `error: the namespace from the provided object "default" does not match the namespace "monitoring"`.
* **The Cause:** The `k8s_deployment.yaml` file had `namespace: default` hardcoded in the metadata, but the command was targeting the `monitoring` namespace.
* **The Rectification:** * Manually edited `k8s_deployment.yaml` to set `namespace: monitoring`.
    * Synced the Jenkinsfile variable `K8S_NAMESPACE` to match.

### Challenge C: Linux File Permissions (The "Security Wall")
* **The Symptom:** `open /home/dr-musa/.kube/config: permission denied` followed by `client.key: permission denied`.
* **The Cause:** Jenkins runs as a restricted service user (`jenkins`). It was physically blocked from reading the Kubernetes "map" and "keys" located in the `/home/dr-musa/` directory.
* **The Rectification:** * Modified directory permissions: `sudo chmod +rx /home/dr-musa` and `sudo chmod +rx /home/dr-musa/.kube`.
    * Set file-level read permissions: `sudo chmod 644 /home/dr-musa/.kube/config`.
    * Unlocked Minikube security keys: `sudo chmod -R 644 /home/dr-musa/.minikube/profiles/minikube/*.key`.

### Challenge D: WSL DNS Failure ("No Such Host")
* **The Symptom:** `ERROR: failed to solve: python:3.9-slim: dial tcp: lookup registry-1.docker.io: no such host`.
* **The Cause:** The WSL network bridge lost its "phonebook" (DNS), making it impossible to find Docker Hub.
* **The Rectification:** * Avoided a WSL restart (which would have changed K8s ports and broken Challenge A again).
    * Manually reconstructed the DNS file: `echo "nameserver 8.8.8.8" | sudo tee /etc/resolv.conf`.
    * Locked the file using `sudo chattr +i /etc/resolv.conf` to prevent WSL from overwriting it.

---

## 🛠️ 3. The Final "Green" Jenkinsfile
```groovy
pipeline {
    agent any
    environment {
        DOCKER_HUB_USER = 'musabalaaudu'
        APP_NAME = 'health-api'
        IMAGE_TAG = "${env.BUILD_NUMBER}"
        K8S_NAMESPACE = 'monitoring'
        DOCKER_HUB_CREDS = credentials('docker-hub-credentials') 
        KUBECONFIG = '/home/dr-musa/.kube/config'
        PROJECT_DIR = 'projects/08-health-api-jenkins'
    }
    stages {
        stage('🏗️ Build Docker Image') {
            steps {
                dir("${PROJECT_DIR}") {
                    sh "docker build -t ${DOCKER_HUB_USER}/${APP_NAME}:${IMAGE_TAG} ."
                }
            }
        }
        stage('🚀 Push to Docker Hub') {
            steps {
                sh "echo \$DOCKER_HUB_CREDS_PSW | docker login -u \$DOCKER_HUB_CREDS_USR --password-stdin"
                sh "docker push ${DOCKER_HUB_USER}/${APP_NAME}:${IMAGE_TAG}"
            }
        }
        stage('☸️ Deploy to Kubernetes') {
            steps {
                dir("${PROJECT_DIR}") {
                    sh "sed -i 's|image:.*|image: docker.io/${DOCKER_HUB_USER}/${APP_NAME}:${IMAGE_TAG}|g' k8s_deployment.yaml"
                    sh "kubectl apply -f k8s_deployment.yaml -n ${K8S_NAMESPACE} --validate=false"
                }
            }
        }
    }
    post {
        always { sh "docker logout" }
    }
}
```

---

## 🎓 4. Key Takeaways
1.  **Automation is Fragile:** A simple permission change or DNS glitch can halt a whole pipeline.
2.  **Service Users Matter:** Always remember that Jenkins is a separate user from you. It doesn't have your "Sudo" powers by default.
3.  **Persistence Pays Off:** Most DevOps work is 10% coding and 90% troubleshooting networking and permissions.
