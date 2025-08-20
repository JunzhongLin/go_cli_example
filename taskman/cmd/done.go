package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done",
	Long:  `Mark a specific task as done by its ID.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a task ID.")
			return
		}
		taskID := args[0]
		markTaskAsDone(taskID)
	},
}