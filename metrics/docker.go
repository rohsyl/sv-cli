package metrics

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"strings"
)

func GetDockerContainers() MetricResult {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}

	var containerList []DockerContainer
	for _, c := range containers {
		containerList = append(containerList, DockerContainer{
			ID:     c.ID[:12],
			Name:   c.Names[0][1:],
			Image:  c.Image,
			Status: c.State,
		})
	}

	return MetricResult{Success: true, Data: containerList}
}

func GetDockerContainerInfo(containerName string) MetricResult {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}
	cli.NegotiateAPIVersion(context.Background())

	// List all containers (running & stopped) to find the matching one
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}

	var containerID string
	for _, c := range containers {
		for _, name := range c.Names {
			if strings.TrimPrefix(name, "/") == containerName {
				containerID = c.ID
				break
			}
		}
	}

	if containerID == "" {
		return MetricResult{Success: false, Error: fmt.Sprintf("Container %s not found", containerName)}
	}

	// Get container details (inspect)
	containerJSON, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}

	// Get container statistics (for memory & CPU usage)
	stats, err := cli.ContainerStats(context.Background(), containerID, false)
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}
	defer stats.Body.Close()

	// Parse container stats
	var containerStats types.StatsJSON
	if err := DecodeJSON(stats.Body, &containerStats); err != nil {
		return MetricResult{Success: false, Error: "Failed to decode container stats"}
	}

	// Calculate CPU usage
	cpuDelta := float64(containerStats.CPUStats.CPUUsage.TotalUsage - containerStats.PreCPUStats.CPUUsage.TotalUsage)
	systemDelta := float64(containerStats.CPUStats.SystemUsage - containerStats.PreCPUStats.SystemUsage)
	cpuUsage := (cpuDelta / systemDelta) * float64(len(containerStats.CPUStats.CPUUsage.PercpuUsage)) * 100.0

	// Memory usage
	memoryUsage := containerStats.MemoryStats.Usage
	// memoryLimit := containerStats.MemoryStats.Limit

	// Return formatted container details
	return MetricResult{
		Success: true,
		Data: DockerContainer{
			ID:     containerJSON.ID[:12],
			Name:   strings.TrimPrefix(containerJSON.Name, "/"),
			Image:  containerJSON.Config.Image,
			Status: containerJSON.State.Status,
			Memory: memoryUsage,
			CPU:    cpuUsage,
		},
	}
}

// DecodeJSON decodes a JSON response from an io.ReadCloser
func DecodeJSON(body io.ReadCloser, target interface{}) error {
	decoder := json.NewDecoder(body)
	return decoder.Decode(target)
}
