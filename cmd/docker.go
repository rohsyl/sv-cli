package cmd

import (
	"context"
	"fmt"
	"sv-cli/utils"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

func NewDockerCmd() *cobra.Command {
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
