package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dr-musa-bala/WordPress-Setup-on-xFusionCorp-Infra/internal/monitor"
)

type Notifier struct {
	webhookURL string
	client     *http.Client
}

func New(webhookURL string) *Notifier {
	return &Notifier{
		webhookURL: webhookURL,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

type Alert struct {
	Timestamp  time.Time                  `json:"timestamp"`
	Message    string                     `json:"message"`
	Containers []*monitor.ContainerStatus `json:"containers"`
}

func (n *Notifier) SendAlert(statuses []*monitor.ContainerStatus) error {
	unhealthy := make([]*monitor.ContainerStatus, 0)
	for _, status := range statuses {
		if !status.Healthy {
			unhealthy = append(unhealthy, status)
		}
	}

	if len(unhealthy) == 0 {
		return nil // Nothing to report
	}

	alert := Alert{
		Timestamp:  time.Now(),
		Message:    fmt.Sprintf("⚠️  %d unhealthy container(s) detected", len(unhealthy)),
		Containers: unhealthy,
	}

	payload, err := json.Marshal(alert)
	if err != nil {
		return err
	}

	resp, err := n.client.Post(n.webhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("webhook returned status %d", resp.StatusCode)
	}

	return nil
}

func (n *Notifier) PrintStatus(statuses []*monitor.ContainerStatus) {
	fmt.Println("\n=== Container Health Check ===")
	fmt.Printf("Time: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))

	for _, status := range statuses {
		healthIcon := "✅"
		if !status.Healthy {
			healthIcon = "❌"
		}

		fmt.Printf("%s %s\n", healthIcon, status.Name)
		fmt.Printf("   ID: %s\n", status.ID)
		fmt.Printf("   State: %s\n", status.State)
		if status.Error != "" {
			fmt.Printf("   Error: %s\n", status.Error)
		}
		fmt.Println()
	}
}
