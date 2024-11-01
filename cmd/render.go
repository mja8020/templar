/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type renderCmdFlags struct {
	check bool
}

var RenderCmdFlags renderCmdFlags

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Renders/Checks template content",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		fmt.Println("render called")
		return
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)

	renderCmd.PersistentFlags().BoolVarP(&RenderCmdFlags.check, "check", "c", false, "Checks the rendered templates without saving and returns non-zero exit code for pending changes")
}
