package main

import (
	"flag"
	"fmt"
	"github.com/MenciusCheng/go-text-template/cmd/gen"
)

func main() {
	configS := flag.String("c", "cmd/config/dao1", "Configuration file")
	flag.Parse()
	if configS == nil {
		fmt.Printf("arg is empty\n")
		return
	}
	dirName := *configS

	if err := gen.GenByDirConfig(dirName); err != nil {
		fmt.Printf("GenByDirConfig err: %s\n", err)
		return
	}
}
