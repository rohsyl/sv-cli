package metrics

import (
	"reflect"
	"fmt"
)

// Registry for mapping metric names to function pointers
var metricsMap = map[string]interface{}{
	"ram":              GetRAMUsage,
	"disk":             GetDiskUsage,
	"cpu":              GetCPUUsage,
	"service":          GetServiceStatus,       // Requires extra parameters
	"docker-container": GetDockerContainerInfo, // Requires extra parameters
	"database":			CheckDatabaseStatus,
}

// CallMetricFunction dynamically executes the corresponding function
func CallMetricFunction(name string, params []string) MetricResult {
	if fn, exists := metricsMap[name]; exists {
		fnValue := reflect.ValueOf(fn)

		// Ensure the function exists and can be called
		if fnValue.Kind() == reflect.Func {
			// Convert params into reflect.Value slice
			args := make([]reflect.Value, len(params))
			for i, param := range params {
				args[i] = reflect.ValueOf(param)
			}

			// Call the function dynamically
			results := fnValue.Call(args)
			fmt.Println(results)

			// Ensure we return the expected MetricResult type
			if len(results) > 0 {
				if metricResult, ok := results[0].Interface().(MetricResult); ok {
					fmt.Println(metricResult)
					return metricResult
				}
			}
		}
	}
	return MetricResult{Success: false, Error: "Metric function not found or invalid parameters"}
}
