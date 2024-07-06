package cmd

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/MasterDaniels/todo/todos"
	"github.com/spf13/cobra"
)

var (
	updateText     string
	updatePriority int
)

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringVarP(&updateText, "update", "u", "", "Text that the todo should be updated to.")
	editCmd.Flags().IntVarP(&updatePriority, "update-priority", "p", 2, "Update priority of the task: (1,2,3)")
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a todo",
	Long:  `Edit the text of an existing todo in the todo app.`,
	Run:   editRun,
}

func editRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Printf("Error: %v.", errors.New("an argument must be passed to the edit command"))
		return
	}

	items, err := todos.ReadItems(GetdataFile())
	if err != nil {
		fmt.Printf("%v.\nError: %v", errors.New("an error happened when reading todo file"), err)
		return
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("\"%v\" is not a valid label.\nEnter a number between 1 and %v.\nOr try wrapping the update string in quotes", args[0], len(items))
		return
	}

	if i > 0 && i < len(items) {
		if updateText != "" {
			items[i-1].Text = updateText
		}

		if items[i-1].Priority != updatePriority {
			items[i-1].SetPriority(updatePriority)
		}

		fmt.Printf("%q updated.\n", items[i-1].Text)

		sort.Sort(todos.ByPri(items))
		todos.SaveItems(GetdataFile(), items)
	} else {
		fmt.Println("Error: ", i, "doesn't match any items.")
	}
}
