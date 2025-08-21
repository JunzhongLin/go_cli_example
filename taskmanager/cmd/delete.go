package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var deleteCmd = &cobra.Command{
	Use: "delete [task number]",
	Short: "Delete a task",
	Long: `Delete a specific task from the task list by its number.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil || index < 1 {
			fmt.Println("Invalid task number.")
			return
		}
		tasks := loadTasks()
		if index > len(tasks) {
			fmt.Printf("Task with number %d does not exist.\n", index)
			return
		}
		tasks = append(tasks[:index-1], tasks[index:]...) // Remove the task at index
		saveTasks(tasks)
		fmt.Printf("Task with number %d has been deleted.\n", index)
	},
}