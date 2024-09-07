/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rohitpandeydev/tasks/internal/services"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the list",
	Long: `Add a task to the list. You can use this command to add a task to the list.
  For example:
  tasks add "Buy milk"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		services.WriteTask(task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
