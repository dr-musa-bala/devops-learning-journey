package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Status defines the JSON structure for our DevOps API
type Status struct {
	Service   string `json:"service"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Uptime    string `json:"uptime"`
}

// Global start time to calculate uptime
var startTime = time.Now()

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create the data object
	data := Status{
		Service:   "Onboarding-Auditor-API",
		Status:    "Healthy",
		Timestamp: time.Now().Format(time.RFC3339),
		Uptime:    time.Since(startTime).String(),
	}

	// Set header so the browser/Bash knows it is JSON
	w.Header().Set("Content-Type", "application/json")
	
	// Encode the struct into JSON and send it
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Get port from Environment Variable (A key DevOps skill)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" 
	}

	http.HandleFunc("/health", healthCheckHandler)

	fmt.Printf("📡 DevOps API starting on port %s...\n", port)
	fmt.Println("👉 Access locally at: http://localhost:8080/health")
	
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Fatal Error: %s\n", err)
	}
}
