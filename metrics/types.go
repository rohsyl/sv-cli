package metrics

import (
    "time"
)

// MetricResult standardizes the return value for all metric functions
type MetricResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
	ErrorCode int       `json:"error_code,omitempty"`
}

// DiskUsage represents the disk information
type DiskUsage struct {
	Path       string `json:"path"`
	TotalSpace uint64 `json:"total_space"`
	FreeSpace  uint64 `json:"free_space"`
}

// RAMUsage represents total and free RAM
type RAMUsage struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

// CPUUsage represents CPU usage percentage
type CPUUsage struct {
	Usage    float64 `json:"usage"`
	CPUCount int     `json:"cpu_count"`
}

// DockerContainer represents basic container details
type DockerContainer struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Status      string    `json:"status"`
	Memory      uint64    `json:"memory_usage,omitempty"`
	MemoryLimit uint64    `json:"memory_limit,omitempty"`
	CPU         float64   `json:"cpu_usage,omitempty"`

	DiskRead    uint64    `json:"disk_read_bytes,omitempty"`
	DiskWrite   uint64    `json:"disk_write_bytes,omitempty"`
	NetworkRx   uint64    `json:"network_rx_bytes,omitempty"`
	NetworkTx   uint64    `json:"network_tx_bytes,omitempty"`

	Timestamp   time.Time `json:"timestamp"`
}



// ServiceStatus represents the status of a system service
type ServiceStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// DatabaseStatus represents the connectivity status of a database
type DatabaseStatus struct {
	Type   string `json:"type"`
	Status string `json:"status"`
}
