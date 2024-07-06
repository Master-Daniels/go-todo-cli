package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version number of the Todo app.",
	Long:    `All software has versions. This is Todo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Todo cli application v0.0.1")
	},
}
