package metrics

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func GetServiceStatus(serviceName string) MetricResult {
	var status string
	var errorCode int
	var success bool = false

	if runtime.GOOS == "linux" {
		out, err := exec.Command("systemctl", "is-active", serviceName).Output()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				errorCode = exitErr.ExitCode()
			} else {
				errorCode = -1 // Unknown error
			}
			status = "unknown"
		} else {
			success = true
			status = strings.TrimSpace(string(out))
		}
	} else if runtime.GOOS == "windows" {
		out, err := exec.Command("sc", "query", serviceName).Output()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				errorCode = exitErr.ExitCode()
			} else {
				errorCode = -1
			}
			status = "unknown"
		} else if strings.Contains(string(out), "RUNNING") {
			success = true
			status = "running"
		} else {
			success = true
			status = "not running"
		}
	} else {
		return MetricResult{Success: false, Error: "Unsupported OS",  ErrorCode: -1}
	}

	fmt.Println(status)
	return MetricResult{
		Success: success, 
		Data: ServiceStatus{Name: serviceName, Status: status},
		ErrorCode: errorCode,
	}
}
