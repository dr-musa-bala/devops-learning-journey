package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dr-musa-bala/WordPress-Setup-on-xFusionCorp-Infra/config"
	"github.com/dr-musa-bala/WordPress-Setup-on-xFusionCorp-Infra/internal/monitor"
	"github.com/dr-musa-bala/WordPress-Setup-on-xFusionCorp-Infra/internal/notifier"
)

func main() {
	configPath := flag.String("config", "config.json", "Path to configuration file")
	flag.Parse()

	// Load configuration
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Println("🚀 Starting Docker Container Monitor")
	fmt.Printf("⏱️  Check interval: %d seconds\n", cfg.CheckInterval)
	fmt.Printf("📦 Monitoring %d containers\n", len(cfg.Containers))

	// Initialize monitor and notifier
	mon := monitor.New(cfg.DockerSocket)
	notif := notifier.New(cfg.AlertWebhook)

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(time.Duration(cfg.CheckInterval) * time.Second)
	defer ticker.Stop()

	// Run initial check
	runCheck(ctx, mon, notif, cfg.Containers)

	// Main loop
	for {
		select {
		case <-ticker.C:
			runCheck(ctx, mon, notif, cfg.Containers)
		case <-sigChan:
			fmt.Println("\n👋 Shutting down gracefully...")
			return
		}
	}
}

func runCheck(ctx context.Context, mon *monitor.Monitor, notif *notifier.Notifier, containers []string) {
	statuses := mon.CheckAll(ctx, containers)
	notif.PrintStatus(statuses)

	// Send alert if webhook is configured
	if err := notif.SendAlert(statuses); err != nil {
		log.Printf("Warning: Failed to send alert: %v", err)
	}
}
