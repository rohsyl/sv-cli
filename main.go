package main

import (
	"fmt"
	"sv-cli/cmd"
	"sv-cli/utils"
)

func main() {
	utils.LoadEnv()

	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
