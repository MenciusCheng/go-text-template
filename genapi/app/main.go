package main

import (
	"fmt"
	"github.com/MenciusCheng/go-text-template/genapi/input"
	"github.com/MenciusCheng/go-text-template/genapi/wtype"
	"github.com/MenciusCheng/go-text-template/parse"
	"strings"
	"text/template"
)

func main() {
	GenWType()
}

func ReadWType() {
	wt := wtype.ReadWTypeByStrList(input.RequestKuaishowGettokenWType)
	fmt.Printf("%+v\n", wt)
}

func GenWType() {
	wt := wtype.ReadWTypeByStrList(input.RequestKuaishowGettokenWType, input.RequestGenWType)

	for _, block := range wt.BlocksByType[wtype.TextTypeTmpl] {
		tplStr, ok := block.Content.(string)
		if !ok {
			continue
		}
		tpl, err := template.New("").Funcs(parse.GetFuncMap()).Parse(tplStr)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}

		sb := strings.Builder{}
		err = tpl.Execute(&sb, wt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("out:\n%s\n", sb.String())
	}
}
