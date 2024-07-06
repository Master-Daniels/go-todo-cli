package cmd

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/MasterDaniels/todo/todos"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo",
	Long:  `Delete an existing todo in the todo app by passing the number of the todo.`,
	Run:   deleteRun,
}

func deleteRun(cmd *cobra.Command, args []string) {
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
		items = append(items[:i-1], items[i:]...)

		sort.Sort(todos.ByPri(items))
		err = todos.SaveItems(GetdataFile(), items)
		if err != nil {
			fmt.Printf("%v.\nError: %v", errors.New("an error happened when saving todo file"), err)
			return
		} else {
			fmt.Printf("Todo number %v deleted.\n", i)
		}
	} else {
		fmt.Printf("Error: %v.", errors.New("the number you entered is not a valid todo"))
		return
	}
}
