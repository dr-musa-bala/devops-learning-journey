# 🐳 Phase 2: Containerization & Portability

This phase focuses on taking the **Source of Truth** (Go API) and the **System Glue** (Bash) and packaging them into a standardized, immutable unit using **Docker**.

## 🏗 The Architecture
The project is now encapsulated in a Docker container using a multi-layered approach:

* **Base OS:** Alpine Linux (Lightweight, security-focused).
* **Runtime:** Go 1.23 Toolchain.
* **System Tools:** `curl`, `jq`, and `bash` baked into the image.
* **Networking:** Port 8080 exposed for external health-check interrogation.

---

## 🛠 Challenges & Resolutions
This section documents the technical hurdles encountered during the containerization process and the engineering logic used to resolve them.

### 1. The "Go Module" Ghost
* **Problem:** The Docker build failed at the compilation stage with the error: `go.mod file not found`.
* **Discovery:** Docker’s isolated environment requires a defined Go module to manage the build context.
* **Resolution:** Executed `go mod init health-api` in the build context.

### 2. The "Context Gap" (Missing Source)
* **Problem:** `RUN go build -o api main.go` failed because `main.go` was absent.
* **Discovery:** Identified a mismatch between the local working directory and the Docker Build Context.
* **Resolution:** Reorganized the project structure to ensure all "ingredients" were present for the `COPY . .` instruction.

### 3. The Toolchain Version Conflict
* **Problem:** `Error: go.mod requires go >= 1.25.7 (running go 1.23.12)`.
* **Discovery:** A version mismatch occurred between the Windows host and the Docker Alpine image.
* **Resolution:** Modified `go.mod` to target `go 1.21`, ensuring backward compatibility.
