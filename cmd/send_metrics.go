package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"sv-cli/metrics"
	"sv-cli/utils"
)

// MetricRequest represents a single metric request
type MetricRequest struct {
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config,omitempty"`
}

// MetricsConfig represents the API response for metric configuration
type MetricsConfig struct {
	Metrics []MetricRequest `json:"metrics"`
}

// MetricsResults represents the JSON payload sent to the API
type MetricsResults struct {
	Results map[string]metrics.MetricResult `json:"results"`
}

// NewFetchMetricsCmd creates the fetch-metrics command
func NewSendMetricsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "send-metrics",
		Short: "Fetches metrics config from API and send results",
		Run: func(cmd *cobra.Command, args []string) {
			// Fetch metric configuration
			config, err := fetchMetricConfig()
			if err != nil {
				fmt.Println("❌ Error fetching configuration:", err)
				return
			}

			// Execute metrics dynamically
			results := executeMetrics(config)

			// Send metrics data to the API
			err = sendMetricsData(results)
			if err != nil {
				fmt.Println("❌ Error submitting metrics:", err)
				return
			}

			fmt.Println("✅ Successfully submitted metrics.")
		},
	}
}

// fetchMetricConfig retrieves metric configurations from the API
func fetchMetricConfig() (*MetricsConfig, error) {
	resp, err := http.Get(utils.GetEnv("METRICS_CONFIG_URL", ""))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	var config MetricsConfig
	err = json.NewDecoder(resp.Body).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// executeMetrics runs the requested metric functions dynamically
func executeMetrics(config *MetricsConfig) MetricsResults {
	results := make(map[string]metrics.MetricResult)

	for _, metric := range config.Metrics {
		// Extract parameters if any
		var params []interface{}
		if metric.Config != nil {
			for _, value := range metric.Config {
				params = append(params, value)
			}
		}

		// Dynamically call the metric function
		results[metric.Type] = metrics.CallMetricFunction(metric.Type, params...)
	}

	return MetricsResults{Results: results}
}

// sendMetricsData submits the collected metrics to the API
func sendMetricsData(results MetricsResults) error {
	jsonData, err := json.Marshal(results)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", utils.GetEnv("METRICS_SUBMIT_URL", ""), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to submit metrics, response code: %d", resp.StatusCode)
	}

	return nil
}
