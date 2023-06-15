package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	rootCmd.AddCommand(timeCmd)
}

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "显示当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("当前时间", time.Now().Format("2006-01-02 15:04:05"))
	},
}
