package main

import (
	_ "embed"
	"fmt"
	"github.com/MenciusCheng/go-text-template/genapi/wtype"
	"io/fs"
	"io/ioutil"
	"regexp"
	"strings"
)

//go:embed kuaishou_api.wtype
var kuaishou_api string

//go:embed origin.txt
var origin string

func main() {
	wt := wtype.NewWType()

	wtype.ReadWTypeByStr(&wt, kuaishou_api, func(b *wtype.Block) {
		b.Def.TextType = wtype.TextTypeTabrow
	})
	tabrow := wt.BlockByType[wtype.TextTypeTabrow].Content.(wtype.Tabrow)
	ksRemarkMap := make(map[string]string)
	for _, row := range tabrow {
		ksRemarkMap[row[1]] = row[3]
	}

	wtype.ReadWTypeByStr(&wt, origin, func(b *wtype.Block) {
		b.Def.Name = "origin"
	})
	compile := regexp.MustCompile("json:\"(\\w+)\"")
	newRows := make([]string, 0)
	for _, line := range wt.BlockByName["origin"].Lines {
		ms := compile.FindStringSubmatch(line)
		if len(ms) < 2 {
			continue
		}
		newRows = append(newRows, fmt.Sprintf("%s // %s", line, ksRemarkMap[ms[1]]))
	}

	str := strings.Join(newRows, "\n")
	err := ioutil.WriteFile("out/addremark.txt", []byte(str), fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
