package metrics

import (
	"os"
	"runtime"
)

func GetSystemInfo() MetricResult {
	hostname, err := os.Hostname()
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}

	return MetricResult{
		Success: true,
		Data: map[string]string{
			"hostname": hostname,
			"os":       runtime.GOOS,
		},
	}
}
