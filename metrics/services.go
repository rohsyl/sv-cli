package metrics

import (
	"os/exec"
	"runtime"
	"strings"
)

func GetServiceStatus(serviceName string) MetricResult {
	var status string
	if runtime.GOOS == "linux" {
		out, err := exec.Command("systemctl", "is-active", serviceName).Output()
		if err != nil {
			status = "unknown"
		} else {
			status = strings.TrimSpace(string(out))
		}
	} else if runtime.GOOS == "windows" {
		out, err := exec.Command("sc", "query", serviceName).Output()
		if err != nil {
			status = "unknown"
		} else if strings.Contains(string(out), "RUNNING") {
			status = "running"
		} else {
			status = "not running"
		}
	} else {
		return MetricResult{Success: false, Error: "Unsupported OS"}
	}

	return MetricResult{Success: true, Data: ServiceStatus{Name: serviceName, Status: status}}
}
