package cmd

import (
	"fmt"

	"github.com/MasterDaniels/todo/todos"
	"github.com/spf13/cobra"
)

var priority int

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority of the task: 1,2,3")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the todo app.",
	Long:  `Add will add a new task(with the ability to psecify the priority of the task) to the todo app.`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todos.ReadItems(GetdataFile())
	if err != nil {
		fmt.Printf("%v.\nRectifying...\n", err)
	}

	for _, task := range args {
		todo := todos.Item{Text: task}
		todo.SetPriority(priority)
		items = append(items, todo)
	}

	if err := todos.SaveItems(GetdataFile(), items); err != nil {
		fmt.Printf("%v", err.Error())
	}
	fmt.Println("Todo added!")
}
