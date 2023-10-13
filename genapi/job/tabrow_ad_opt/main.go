package main

import (
	_ "embed"
	"fmt"
	"github.com/MenciusCheng/go-text-template/genapi/wtype"
)

//go:embed tabrow.wtype
var tabrowWtype string

func main() {
	wt := wtype.ReadWTypeByStrList(tabrowWtype)

	tabrow := wt.BlockByName["ad"].Content.(wtype.Tabrow)

	myTabrow := wtype.Tabrow{}
	myLine := make([]string, 0)
	for _, line := range tabrow {
		if len(line) == 0 || len(line[0]) == 0 {
			continue
		}
		myLine = append(myLine, line[0])
		if len(myLine) >= 5 {
			myTabrow = append(myTabrow, myLine)
			myLine = make([]string, 0)
		}
	}
	if len(myLine) >= 5 {
		myTabrow = append(myTabrow, myLine)
	}
	wt.BlockByName["ad"].Content = myTabrow

	fmt.Println(wt.GenByTmpl())
}
