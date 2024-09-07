/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rohitpandeydev/tasks/internal/services"

	"github.com/spf13/cobra"
)

var detailOutput bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all the tasks",
	Long: `List all the tasks which are there. You can use flag -a to list all the task.
  If you are not using flag then non compete task will be listed`,
	Run: func(cmd *cobra.Command, args []string) {
		services.ListOutputDetailed(detailOutput)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&detailOutput, "all", "a", false, "list all the tasks")
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
