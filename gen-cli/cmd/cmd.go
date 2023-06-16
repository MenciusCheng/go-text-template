package cmd

import (
	"fmt"
	"github.com/MenciusCheng/go-text-template/utils/fileutil"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"text/template"
)

func init() {
	rootCmd.AddCommand(cmdCmd)
}

var cmdCmd = &cobra.Command{
	Use:   "cmd",
	Short: "init subcommand code",
	Args:  cobra.MinimumNArgs(1),
	Run:   cmdCmdRun,
}

func cmdCmdRun(cmd *cobra.Command, args []string) {
	name := args[0]
	remark := ""
	if len(args) > 1 {
		remark = args[1]
	}

	cmdParser, err := template.New("").Parse(cmdTmpl)
	if err != nil {
		fmt.Println("Parse error", err)
		os.Exit(1)
	}

	// create file
	fileName := name
	if !strings.HasSuffix(fileName, ".go") {
		fileName = fmt.Sprintf("%s.go", fileName)
	}
	if fileutil.IsExist(fileName) {
		fmt.Println("file exist")
		os.Exit(1)
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Create error", err)
		os.Exit(1)
	}

	// generate content
	err = cmdParser.Execute(file, map[string]interface{}{
		"name":   name,
		"remark": remark,
	})
	if err != nil {
		fmt.Println("Execute error", err)
		os.Exit(1)
	}
}

var cmdTmpl = `package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand({{ .name }}Cmd)
}

var {{ .name }}Cmd = &cobra.Command{
	Use:   "{{ .name }}",
	Short: "{{ .remark }}",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}`
