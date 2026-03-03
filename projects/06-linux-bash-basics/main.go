package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func main() {
	// 1. Setup Redis Client
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// 2. Define Routes
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/visit", func(w http.ResponseWriter, r *http.Request) {
		val, err := rdb.Incr(ctx, "visitor_count").Result()
		if err != nil {
			http.Error(w, "Could not reach Redis", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Visitor count: %d", val)
	})

	// 3. Start Server
	port := "9090"
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
