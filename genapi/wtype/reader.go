package wtype

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

func ReadWTypeByStrList(strList ...string) WType {
	wt := NewWType()
	for _, str := range strList {
		ReadWTypeByStr(&wt, str)
	}
	wt.GenMap()
	return wt
}

func ReadWTypeByStr(wt *WType, str string, opts ...FormatFunc) {
	// 读取块内容
	bsc := DefaultBlockScanner()
	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		line := scanner.Text()

		if IsBlockDefLine(line) { // 块内容定义行
			if bsc.Buffered {
				wt.Blocks = append(wt.Blocks, bsc.ToBlock(opts...))
			}
			bsc = NewBlockScanner(line)
		} else {
			bsc.ReadLine(line)
		}
	}
	if bsc.Buffered {
		wt.Blocks = append(wt.Blocks, bsc.ToBlock(opts...))
	}
	wt.GenMap()
}

func ReadWTypeByStr2(str string) {
	split := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
	fmt.Printf("结果：%+v\n", split)
}
