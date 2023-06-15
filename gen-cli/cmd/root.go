package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Source string

func init() {
	rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
}

var rootCmd = &cobra.Command{
	Use:   "gencli",
	Short: "gencli is a very nice generator",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
