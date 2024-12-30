package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"sv-cli/utils"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

func NewDockerCmd() *cobra.Command {
	dockerCmd := &cobra.Command{
		Use:   "docker",
		Short: "Manage Docker containers",
	}

	// Add subcommands
	dockerCmd.AddCommand(NewDockerShowCmd(), NewDockerListCmd())

	return dockerCmd
}

func NewDockerShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "docker [container_name]",
		Short: "Check if a Docker container is running",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			containerName := args[0]

			cli, err := client.NewClientWithOpts(client.FromEnv)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			cli.NegotiateAPIVersion(context.Background())

			containers, _ := cli.ContainerList(context.Background(), container.ListOptions{All: true})
			status := "not found"
			for _, c := range containers {
				if c.Names[0] == "/"+containerName {
					status = "running"
				}
			}

			result := map[string]string{"container": containerName, "status": status}
			if format == "json" {
				utils.PrintJSON(result)
			} else {
				utils.PrintTable([][]string{{containerName, status}}, []string{"Container", "Status"})
			}
		},
	}
}

func NewDockerListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all running Docker containers",
		Run: func(cmd *cobra.Command, args []string) {
			// Retrieve the format flag
			format, _ := cmd.Flags().GetString("format")

			// Create Docker client
			cli, err := client.NewClientWithOpts(client.FromEnv)
			if err != nil {
				fmt.Printf("Error creating Docker client: %v\n", err)
				return
			}
			cli.NegotiateAPIVersion(context.Background())

			// List running containers
			containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
			if err != nil {
				fmt.Printf("Error listing containers: %v\n", err)
				return
			}

			if format == "json" {
				output, _ := json.MarshalIndent(containers, "", "  ")
				fmt.Println(string(output))
			} else {
				// Prepare table data
				tableData := [][]string{}
				for _, c := range containers {
					name := ""
					if len(c.Names) > 0 {
						name = c.Names[0][1:] // Remove the leading '/'
					}

					labels := ""
					for key, value := range c.Labels {
						labels += fmt.Sprintf("%s=%s ", key, value)
					}

					tableData = append(tableData, []string{
						c.ID[:12], // Truncated container ID
						name,
						c.Image,
						c.State,
						c.Status,
					})
				}

				// Print table
				utils.PrintTable(tableData, []string{"Container ID", "Name", "Image", "State", "Status"})
			}
		},
	}
}
