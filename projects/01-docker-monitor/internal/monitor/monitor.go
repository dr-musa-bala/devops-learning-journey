package monitor

import (
	"context"
	"encoding/json"
	"os/exec"
	"strings"
)

type ContainerStatus struct {
	Name    string
	ID      string
	State   string
	Healthy bool
	Error   string
}

type Monitor struct {
	// keep field for future options
	_ struct{}
}

func New(_ string) *Monitor {
	return &Monitor{}
}

func (m *Monitor) CheckContainer(ctx context.Context, containerName string) (*ContainerStatus, error) {
	cmd := exec.CommandContext(ctx, "docker", "inspect", containerName)
	out, err := cmd.Output()
	if err != nil {
		// docker inspect returns non-zero if not found; return friendly status
		return &ContainerStatus{
			Name:    containerName,
			State:   "not_found",
			Healthy: false,
			Error:   err.Error(),
		}, nil
	}

	// docker inspect returns a JSON array
	var infos []struct {
		ID    string `json:"Id"`
		State struct {
			Status  string `json:"Status"`
			Running bool   `json:"Running"`
			Health  *struct {
				Status string `json:"Status"`
			} `json:"Health"`
		} `json:"State"`
	}

	if err := json.Unmarshal(out, &infos); err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return &ContainerStatus{
			Name:    containerName,
			State:   "not_found",
			Healthy: false,
			Error:   "no inspect data",
		}, nil
	}

	info := infos[0]
	id := info.ID
	if len(id) >= 12 {
		id = id[:12]
	}

	status := &ContainerStatus{
		Name:    containerName,
		ID:      id,
		State:   info.State.Status,
		Healthy: info.State.Running,
	}

	if info.State.Health != nil {
		status.Healthy = strings.EqualFold(info.State.Health.Status, "healthy")
	}

	return status, nil
}

func (m *Monitor) CheckAll(ctx context.Context, containers []string) []*ContainerStatus {
	results := make([]*ContainerStatus, len(containers))

	type res struct {
		i int
		s *ContainerStatus
	}
	ch := make(chan res, len(containers))

	for i, c := range containers {
		go func(idx int, name string) {
			st, err := m.CheckContainer(ctx, name)
			if err != nil {
				st = &ContainerStatus{
					Name:    name,
					Healthy: false,
					Error:   err.Error(),
				}
			}
			ch <- res{i: idx, s: st}
		}(i, c)
	}

	for i := 0; i < len(containers); i++ {
		r := <-ch
		results[r.i] = r.s
	}

	return results
}
