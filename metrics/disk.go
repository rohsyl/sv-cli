package metrics

import (
	"github.com/shirou/gopsutil/disk"
)

func GetDiskUsage() MetricResult {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}

	var disks []DiskUsage
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}
		disks = append(disks, DiskUsage{
			Path:       p.Mountpoint,
			TotalSpace: usage.Total,
			FreeSpace:  usage.Free,
		})
	}

	return MetricResult{Success: true, Data: disks}
}
