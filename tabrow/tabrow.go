package tabrow

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func LinesToRowsByReg(lines []string, regStr string) [][]string {
	compileReg := regexp.MustCompile(regStr)
	res := make([][]string, 0)
	for _, line := range lines {
		subs := compileReg.FindStringSubmatch(line)
		if len(subs) > 0 {
			res = append(res, append([]string{line}, subs...))
		}
	}
	return res
}

// 根据tab分割符转化成rows
func LinesToRowsByTab(lines []string) [][]string {
	res := make([][]string, 0)
	for _, line := range lines {
		subs := strings.Split(line, "\t")
		if len(subs) > 0 {
			res = append(res, subs)
		}
	}
	return res
}

// 读取文本，根据换行符分割成数组
func ReadFileLines(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Open err:%v\n", err)
		return nil
	}
	defer func() {
		_ = f.Close()
	}()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return fileLines
}
