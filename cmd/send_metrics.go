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
	Type   string         `json:"type"`
	Enabled int			  `json:"enabled"`
	Config  *MetricConfig `json:"config,omitempty"`
}

type MetricConfig struct {
	Items []MetricItem     `json:"items,omitempty"`
}

type MetricItem struct {
	Enabled int            `json:"enabled"`
	Name    string         `json:"name,omitempty"`
	Params  []string  	   `json:"params,omitempty"`
}

// MetricsConfig represents the API response for metric configuration
type MetricsConfig struct {
	ID      int              `json:"id"`
	Metrics []MetricRequest  `json:"metrics"`
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
	url := utils.GetEnv("METRICS_CONFIG_URL", "")
	token := utils.GetEnv("API_KEY", "")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
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

		if(metric.Enabled == 0) {
			continue;
		}

		// Extract parameters if any
		if (metric.Config != nil && len(metric.Config.Items) > 0) {
			for _, item := range metric.Config.Items {
				// Dynamically call the metric function
				results[metric.Type] = metrics.CallMetricFunction(metric.Type, item.Params)
			}
		} else {
			results[metric.Type] = metrics.CallMetricFunction(metric.Type, nil)
		}

	}

	return MetricsResults{Results: results}
}

// sendMetricsData submits the collected metrics to the API
func sendMetricsData(results MetricsResults) error {
	token := utils.GetEnv("API_KEY", "")

	jsonData, err := json.Marshal(results)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", utils.GetEnv("METRICS_SUBMIT_URL", ""), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
