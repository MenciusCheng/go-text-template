package main

import (
	_ "embed"
	"fmt"
	"github.com/MenciusCheng/go-text-template/genapi/wtype"
	"github.com/MenciusCheng/go-text-template/parse"
	"io/fs"
	"io/ioutil"
	"strings"
)

//go:embed kuaishou_api.wtype
var kuaishou_api string

func main() {
	wt := wtype.NewWType()
	wtype.ReadWTypeByStr(&wt, kuaishou_api, func(b *wtype.Block) {
		b.Def.TextType = wtype.TextTypeTabrow
		b.FormatAfter = KuaishouAPI
	})
	str := wt.BlockByType[wtype.TextTypeTabrow].Content.(string)

	err := ioutil.WriteFile("out/genksapi.txt", []byte(str), fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

func KuaishouAPI(b *wtype.Block) {
	tabrow, ok := b.Content.(wtype.Tabrow)
	if !ok {
		return
	}

	sb := strings.Builder{}
	var ksapiMap = make(map[string]bool)
	for _, cells := range tabrow {
		if len(cells) < 4 {
			continue
		}

		fieldName := strings.TrimSpace(cells[1])
		if ksapiMap[fieldName] {
			continue
		}
		ksapiMap[fieldName] = true

		var t string
		switch strings.ToLower(strings.TrimSpace(cells[2])) {
		case "string":
			t = "string"
		case "int":
			t = "int64"
		case "long":
			t = "int64"
		case "double":
			t = "float64"
		case "-":
			continue
		default:
			panic(fmt.Sprintf("未检查类型:'%s'", cells[2]))
		}
		camel := parse.SnakeToCamel(fieldName)
		remark := strings.TrimSpace(cells[3])

		str := fmt.Sprintf("%s %s `json:\"%s\"` // %s\n", camel, t, fieldName, remark)
		sb.WriteString(str)
	}
	b.Content = sb.String()
}
