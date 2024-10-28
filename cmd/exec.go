/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mja8020/templar/internal/exec"
	"github.com/spf13/cobra"
)

var (
	executable string
	envVars    map[string]string
	args       []string
)

// execCmd is the execute command for Cobra
var execCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute a command with specified environment variables and arguments",
	RunE: func(cmd *cobra.Command, _ []string) error {
		// Call the Exec function from internal/exec package
		result, err := exec.Run(executable, envVars, args)
		if err != nil {
			return fmt.Errorf("execution failed: %w", err)
		}

		// Print the results
		fmt.Printf("StdOut:\n%s\n", result.StdOut)
		fmt.Printf("StdErr:\n%s\n", result.StdErr)
		fmt.Printf("ExitCode: %d\n", result.ExitCode)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	execCmd.Flags().StringVarP(&executable, "executable", "e", "/bin/bash", "Path to the executable")
	execCmd.Flags().StringToStringVarP(&envVars, "env", "v", nil, "Environment variables (key=value)")
	execCmd.Flags().StringArrayVarP(&args, "args", "a", nil, "Arguments to pass to the executable")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
