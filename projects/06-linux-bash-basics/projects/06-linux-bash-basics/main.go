package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8" // You'll need to run 'go get' for this
)

var ctx = context.Background()

func main() {
	// 1. Get port from environment or default to 8080
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// 2. Connect to Redis
	// NOTE: "redis-db" matches the service name in docker-compose.yml
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis-db:6379",
	})
        // Inside your main function, right after connecting to Redis:
err := rdb.Ping(ctx).Err()
if err != nil {
    log.Printf("⚠️ Redis not ready: %v", err)
    // We DON'T use log.Fatal here so the API stays alive 
    // and we can actually see the error in the browser/curl.
} 
else {
    log.Println("✅ Successfully connected to Redis!")
}
	// 3. Define the Health Check (Existing)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	// 4. NEW: The Visit Counter Endpoint
	http.HandleFunc("/visit", func(w http.ResponseWriter, r *http.Request) {
		// Increment the "hits" key in Redis
		hits, err := rdb.Incr(ctx, "hits").Result()
		if err != nil {
			http.Error(w, "Database Error", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Hello! You are visitor number: %d\n", hits)
	})

	log.Printf("📡 API starting on port %s...", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
