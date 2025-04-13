package metrics

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

func GetSystemInfo() MetricResult {
	hostname, err := os.Hostname()
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}

	info := map[string]string{
		"hostname": hostname,
		"os":       runtime.GOOS,
	}

	var osInfo map[string]string

	if runtime.GOOS == "windows" {
		osInfo = getWindowsOSInfo()
	} else {
		osInfo, err = parseOSRelease("/etc/os-release")
		if err != nil {
			return MetricResult{Success: false, Error: err.Error()}
		}
	}

	// Merge relevant OS fields into info
	for _, key := range []string{"PRETTY_NAME", "ID", "VERSION_ID", "ID_LIKE"} {
		if val, ok := osInfo[key]; ok {
			info[strings.ToLower(key)] = val
		}
	}

	return MetricResult{
		Success: true,
		Data:    info,
	}
}

func parseOSRelease(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			key := parts[0]
			val := strings.Trim(parts[1], `"`)
			info[key] = val
		}
	}
	return info, scanner.Err()
}

