package cmd

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"io"

	"github.com/MasterDaniels/todo/todos"
	"github.com/spf13/cobra"
)

var doneOpt, allOpt bool

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&doneOpt, "done", "d", false, "Show 'Done' Todos.")
	listCmd.Flags().BoolVarP(&allOpt, "all", "a", false, "Show all Todos.")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the tasks of the todo app.",
	Long:  `List all tasks have been saved on the todo app.`,
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todos.ReadItems(GetdataFile())
	if err != nil {
		fmt.Printf("%v.\nError: %v", errors.New("an error happened when reading todo file"), err)
	} else {

		sort.Sort(todos.ByPri(items))
		w := new(tabwriter.Writer).Init(io.Writer(os.Stdout), 3, 0, 1, byte(' '), uint(0))

		for _, i := range items {
			if allOpt || i.Done == doneOpt {
				fmt.Fprintln(w, i.Label()+"\t"+i.PrettyDone()+"\t"+i.PrettyP()+"\t"+i.Text+"\t")
			}
		}

		w.Flush()
	}
}
