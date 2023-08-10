package main

import (
	_ "embed"
	"fmt"
	"github.com/MenciusCheng/go-text-template/genapi/wtype"
	"regexp"
)

//go:embed tabrow.wtype
var tabrowWtype string

func main() {
	wt := wtype.ReadWTypeByStrList(tabrowWtype)
	tableBlock := wt.BlockByName["tables"]
	m := tableBlock.Content.(map[string]interface{})
	arr := m["data"].([]interface{})
	needArr := make([]string, 0)
	compile := regexp.MustCompile("^kefu_session_\\d+")
	for _, item := range arr {
		str := item.(string)
		if compile.MatchString(str) {
			needArr = append(needArr, str)
		}
	}
	tableBlock.Format = needArr

	fmt.Println(wt.GenByTmpl())
}
