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
    "time"
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

	containerJSON, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}

	stats, err := cli.ContainerStats(context.Background(), containerID, false)
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}
	defer stats.Body.Close()

	var containerStats types.StatsJSON
	if err := DecodeJSON(stats.Body, &containerStats); err != nil {
		return MetricResult{Success: false, Error: "Failed to decode container stats"}
	}

	// CPU usage
	cpuDelta := float64(containerStats.CPUStats.CPUUsage.TotalUsage - containerStats.PreCPUStats.CPUUsage.TotalUsage)
	systemDelta := float64(containerStats.CPUStats.SystemUsage - containerStats.PreCPUStats.SystemUsage)
	cpuUsage := (cpuDelta / systemDelta) * float64(len(containerStats.CPUStats.CPUUsage.PercpuUsage)) * 100.0

	// Memory
	memoryUsage := containerStats.MemoryStats.Usage
	memoryLimit := containerStats.MemoryStats.Limit

	// Disk I/O (Read/Write)
	var blkRead, blkWrite uint64
	for _, blk := range containerStats.BlkioStats.IoServiceBytesRecursive {
		switch strings.ToLower(blk.Op) {
		case "read":
			blkRead += blk.Value
		case "write":
			blkWrite += blk.Value
		}
	}

	// Network I/O (Rx/Tx)
	var netRx, netTx uint64
	for _, net := range containerStats.Networks {
		netRx += net.RxBytes
		netTx += net.TxBytes
	}

	return MetricResult{
		Success: true,
		Data: DockerContainer{
			ID:          containerJSON.ID[:12],
			Name:        strings.TrimPrefix(containerJSON.Name, "/"),
			Image:       containerJSON.Config.Image,
			Status:      containerJSON.State.Status,
			Memory:      memoryUsage,
			MemoryLimit: memoryLimit,
			CPU:         cpuUsage,
			DiskRead:    blkRead,
			DiskWrite:   blkWrite,
			NetworkRx:   netRx,
			NetworkTx:   netTx,
			Timestamp:   time.Now(),
		},
	}
}

// DecodeJSON decodes a JSON response from an io.ReadCloser
func DecodeJSON(body io.ReadCloser, target interface{}) error {
	decoder := json.NewDecoder(body)
	return decoder.Decode(target)
}
