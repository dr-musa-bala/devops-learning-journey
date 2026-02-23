🐳 Phase 2: Containerization & Portability
This phase focuses on taking the Source of Truth (Go API) and the System Glue (Bash) and packaging them into a standardized, immutable unit using Docker.

🏗 The Architecture
The project is now encapsulated in a Docker container using a multi-layered approach:

Base OS: Alpine Linux (Lightweight, security-focused).

Runtime: Go 1.23 Toolchain.

System Tools: curl, jq, and bash baked into the image.

Networking: Port 8080 exposed for external health-check interrogation.

🛠 Challenges & Resolutions
This section documents the technical hurdles encountered during the containerization process and the engineering logic used to resolve them.

1. The "Go Module" Ghost
Problem: The Docker build failed at the compilation stage with the error: go.mod file not found.

Discovery: Docker’s isolated environment requires a defined Go module to manage the build context, even for single-file applications.

Resolution: Executed go mod init health-api in the build context to provide the necessary metadata for the Go compiler.

2. The "Context Gap" (Missing Source)
Problem: RUN go build -o api main.go failed because main.go was absent from the build directory.

Discovery: Identified a mismatch between the local working directory and the Docker Build Context. The file was residing in a parent directory.

Resolution: Utilized Linux find and mv commands to reorganize the project structure, ensuring all "ingredients" were present for the COPY . . instruction.

3. The Toolchain Version Conflict
Problem: Error: go.mod requires go >= 1.25.7 (running go 1.23.12).

Discovery: A version mismatch occurred because the Windows host had a newer Go version than the Docker Alpine image. Go is strictly backward-compatible but not forward-compatible by default.

Resolution: Modified the go.mod file to target go 1.21 and removed the specific toolchain requirement. This ensured the code remained backward compatible and portable across different environments.

🚀 Technical Implementation (The Dockerfile)
Dockerfile
# Start with a stable Go environment
FROM golang:1.23-alpine

# Install essential 'Glue' tools
RUN apk add --no-cache bash curl jq

WORKDIR /app

# Copy project files (Source of Truth + Glue)
COPY . .

# Compile the Go binary
RUN go build -o api main.go

# Ensure the monitoring script has execution rights
RUN chmod +x read_json.sh

EXPOSE 8080

# Execute the API on startup
CMD ["./api"]
📊 Performance & Mastery Metrics
Environment Parity: The "It works on my machine" problem is 100% eliminated. The API runs identically on WSL2, Windows, or a Cloud Server.

Image Optimization: Used Alpine Linux, reducing the operating system footprint significantly compared to standard distributions.

Portability: The entire stack (API + Monitoring Tools) can be deployed with a single command: docker run.

🏁 Final Verification Commands
Bash
# Build the image
docker build -t health-monitor .

# Run the containerized service
docker run -d -p 8080:8080 --name my-health-service health-monitor

# Interrogate the container internals
docker exec -it my-health-service ./read_json.sh
