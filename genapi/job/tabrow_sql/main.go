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
	fmt.Println(wt.GenByTmpl())
}
