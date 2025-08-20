package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/load"
	"go.mongodb.org/mongo-driver/mongo/description"
)

type Task struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var addCmd = &cobra.Command{
	Use:  "add [task description]",
	Short: "Add a new task",
	Long: `Add a new task to the task list with the specified description.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description = args[0]
		tasks := loadTasks()
		tasks = append(tasks, Task{Description: description, Done: false})
		saveTasks(tasks)
		fmt.Printf("Tasks added: %s \n", description)
	}
}

func loadTasks() []Task {

}