package metrics

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

func GetCPUUsage() MetricResult {
	usage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}

	cpuCount, err := cpu.Counts(true)
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}

	return MetricResult{Success: true, Data: CPUUsage{Usage: usage[0], CPUCount: cpuCount}}
}
