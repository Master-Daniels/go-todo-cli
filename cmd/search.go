package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MasterDaniels/todo/todos"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&searchText, "search", "s", "", "Text to search for in the todo.")
	searchCmd.MarkFlagRequired("search")
}

var searchText string
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a Todo.",
	Long:  "Search for a todo by text",
	Run:   searchRun,
}

func searchRun(cmd *cobra.Command, args []string) {
	items, err := todos.ReadItems(GetdataFile())
	if err != nil {
		fmt.Printf("%v.\nError: %v", errors.New("an error happened when reading todo file"), err)
		return
	}

	item := todos.Item{}

	if searchText != "" {
		for _, itemSearch := range items {
			if strings.Contains(itemSearch.Text, searchText) {
				item = itemSearch
				break
			}
		}
	}

	fmt.Println("Todo found.\n", item.Label()+"  "+item.PrettyDone()+"  "+item.PrettyP()+" "+item.Text)
}
