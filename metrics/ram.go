package metrics

import (
	"github.com/shirou/gopsutil/mem"
)

func GetRAMUsage() MetricResult {
	v, err := mem.VirtualMemory()
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}

	return MetricResult{Success: true, Data: RAMUsage{Total: v.Total, Free: v.Free}}
}
