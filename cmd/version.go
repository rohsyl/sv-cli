package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "sv-cli/internal"
)

func VersionCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "version",
        Short: "Print the version number",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("sv-cli " + internal.Version)
        },
    }
}