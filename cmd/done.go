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
	rootCmd.AddCommand(doneCmd)
}

var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do", "D"},
	Short:   "Mark a Todo item as Done",
	Long: `Mark a Todo as done by entering the number of the Todo on the list.
Toggle done on a Todo by running the same command again.`,
	Run: doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Printf("Error: %v.", errors.New("an argument must be passed to the done command"))
		return
	}

	items, err := todos.ReadItems(GetdataFile())

	if err != nil {
		fmt.Printf("%v.\nError: %v", errors.New("an error happened when reading todo file"), err)
		return
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("\"%v\" is not a valid label.\nEnter a number between 1 and %v.", args[0], len(items))
		return
	}

	if i > 0 && i < len(items) {
		items[i-1].Done = !items[i-1].Done

		fmt.Printf("%q updated to Done: %v.\n", items[i-1].Text, items[i-1].Done)

		sort.Sort(todos.ByPri(items))
		todos.SaveItems(GetdataFile(), items)

	} else {
		fmt.Println("Error: ", i, "doesn't match any items.")
	}
}
