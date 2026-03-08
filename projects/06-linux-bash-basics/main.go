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
	// 1. Setup Redis Client
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		fmt.Println("CRITICAL: REDIS_ADDR environment variable is NOT set. Falling back to localhost.")
		redisAddr = "localhost:6379"
	} else {
		fmt.Printf("Attempting to connect to Redis at: %s\n", redisAddr)
	}

	rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// 2. Immediate Ping Test
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("CRITICAL: Could not connect to Redis: %v", err)
	}
	fmt.Println("Successfully connected to Redis!")

	// 3. Handlers
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		val, err := rdb.Incr(ctx, "visitor_count").Result()
		if err != nil {
			fmt.Printf("Redis error: %v\n", err)
			http.Error(w, "Could not reach Redis", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Visitor count: %d", val)
	})

	// 4. Start Server
	port := "9090"
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
} // <--- MAKE SURE THIS CLOSING BRACE IS HERE
