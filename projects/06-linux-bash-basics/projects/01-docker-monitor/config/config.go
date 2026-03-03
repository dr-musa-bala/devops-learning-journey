package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	CheckInterval int      `json:"check_interval"` // seconds
	DockerSocket  string   `json:"docker_socket"`
	AlertWebhook  string   `json:"alert_webhook"`
	Containers    []string `json:"containers"` // container names to monitor
}

func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	// Set defaults
	if config.DockerSocket == "" {
		config.DockerSocket = "/var/run/docker.sock"
	}
	if config.CheckInterval == 0 {
		config.CheckInterval = 30
	}

	return &config, nil
}
