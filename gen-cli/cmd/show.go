package cmd

import (
	"fmt"
	"github.com/MenciusCheng/go-text-template/utils/fileutil"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {
	showCmd.Flags().StringVarP(&ShowConfig.Line, "line", "l", "", "选中的行")
	showCmd.MarkFlagRequired("line")

	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "显示选中的文本",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ShowConfig.File = args[0]

		fileLines, err := fileutil.ReadFileByLine(ShowConfig.File)
		if err != nil {
			fmt.Println("read file error", err)
			os.Exit(1)
		}

		line, err := strconv.Atoi(ShowConfig.Line)
		if err != nil {
			fmt.Println("Line error", err)
			os.Exit(1)
		}

		if line <= 0 || line > len(fileLines) {
			fmt.Println("wrong line", line)
			os.Exit(1)
		}

		fmt.Println(fileLines[line-1])
	},
}

var ShowConfig struct {
	Line string `json:"line"`
	File string `json:"file"`
}
