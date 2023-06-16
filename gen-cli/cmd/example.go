package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(exampleCmd)
}

var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "show examples of specified command",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		example := args[0]

		switch example {
		case "dbbean":
			fmt.Println("gen-cli dbbean --db ai -t table_name")
		}
	},
}
