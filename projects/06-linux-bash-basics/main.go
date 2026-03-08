package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func main() {
	// 1. Get the connection string from Render
	redisURL := os.Getenv("REDIS_ADDR")
	if redisURL == "" {
		fmt.Println("CRITICAL: REDIS_ADDR not set, using localhost")
		redisURL = "redis://localhost:6379"
	}

	// 2. Parse the URL (This fixes the "too many colons" error)
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("CRITICAL: Invalid Redis URL: %v", err)
	}

	rdb = redis.NewClient(opt)

	// 3. Immediate Ping Test (The "Senior" Move)
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("CRITICAL: Could not connect to Redis: %v", err)
	}
	fmt.Println("Successfully connected to Redis!")

	// 4. Handlers (Keep these exactly as they were)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		val, err := rdb.Incr(ctx, "visitor_count").Result()
		if err != nil {
			fmt.Printf("Redis error: %v\n", err)
			http.Error(w, "Could not reach Redis", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Visitor count: %d", val)
	})

	// 5. Start Server
	port := "9090"
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
