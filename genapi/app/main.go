package main

import (
	"fmt"
	"github.com/MenciusCheng/go-text-template/genapi/input"
	"github.com/MenciusCheng/go-text-template/genapi/wtype"
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
	fmt.Printf("out:\n%s\n", wt.GenByTmpl())
}
