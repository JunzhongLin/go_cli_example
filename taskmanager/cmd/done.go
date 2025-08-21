package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done",
	Long:  `Mark a specific task as done by its ID.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])

		if err != nil || index < 1 {
			fmt.Println("Invalid task ID.")
			return
		}
		tasks := loadTasks()
		if index > len(tasks) {
			fmt.Printf("Task with ID %d does not exist.\n", index)
			return
		}
		tasks[index-1].Done = true
		saveTasks(tasks)
		fmt.Printf("Task with ID %d marked as done.\n", index)
	},
}