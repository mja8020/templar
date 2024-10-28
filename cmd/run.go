/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs pre-configured commands against the target folders.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) == 0 {
			return fmt.Errorf("At least one command must be specified")
		}

		fmt.Println("run called")
		return
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
