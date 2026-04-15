# Use the official Golang image to build our 'Source of Truth'
FROM golang:1.23-alpine

# Install the 'Glue' tools (curl and jq) inside the container
RUN apk add --no-cache bash curl jq

# Set the working directory inside the container
WORKDIR /app

# Copy your local files into the container
COPY . .

# Build your Go API
RUN go build -o api main.go

# Make sure your Bash script is executable
RUN chmod +x read_json.sh

# Expose the port your API runs on
EXPOSE 8080

# The command to start your API
CMD ["./api"]
