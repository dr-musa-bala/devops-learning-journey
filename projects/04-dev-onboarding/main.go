package main

import (
	"context"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

// CheckResult holds the outcome of a system check
type CheckResult struct {
	Name    string
	Status  bool
	Message string
}

// runCheck executes a shell command to see if a tool is installed
func runCheck(ctx context.Context, name string, command string, args []string, results chan<- CheckResult, wg *sync.WaitGroup) {
	defer wg.Done()

	// Use CommandContext to allow for timeouts—a key DevOps pattern
	cmd := exec.CommandContext(ctx, command, args...)
	err := cmd.Run()

	if err != nil {
		results <- CheckResult{Name: name, Status: false, Message: "Not found or error occurred"}
		return
	}
	results <- CheckResult{Name: name, Status: true, Message: "Installed successfully"}
}

func main() {
	fmt.Println("🚀 Starting DevOps Onboarding Check...")

	// Define the tools we need to check
	tools := map[string][]string{
		"Git":    {"version"},
		"Docker": {"--version"},
		"Go":     {"version"},
	}

	results := make(chan CheckResult, len(tools))
	var wg sync.WaitGroup

	// Set a timeout for the entire operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Launch checks concurrently
	for name, args := range tools {
		wg.Add(1)
		go runCheck(ctx, name, name, args, results, &wg)
	}

	// Wait for all checks to finish in the background
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results as they come in
	allPassed := true
	for res := range results {
		status := "✅"
		if !res.Status {
			status = "❌"
			allPassed = false
		}
		fmt.Printf("%s %-10s: %s\n", status, res.Name, res.Message)
	}

	if allPassed {
		fmt.Println("\n🎉 Environment is ready! Welcome to the team.")
	} else {
		fmt.Println("\n⚠️  Some tools are missing. Please install them to continue.")
	}
}
