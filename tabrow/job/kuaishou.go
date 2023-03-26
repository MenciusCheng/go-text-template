package job

import (
	"fmt"
	"github.com/MenciusCheng/go-text-template/tabrow"
	"io/fs"
	"io/ioutil"
	"strings"
)

func FixKuaishou() {
	ksapiFile := "tabrow/input/kuaishou_api.txt"
	ksRows := tabrow.LinesToRowsByTab(tabrow.ReadFileLines(ksapiFile))
	ksRemarkMap := make(map[string]string)
	for _, row := range ksRows {
		ksRemarkMap[strings.TrimSpace(row[0])] = strings.TrimSpace(row[2])
	}

	fileName := "tabrow/input/origin.txt"
	rows := tabrow.LinesToRowsByReg(tabrow.ReadFileLines(fileName), "json:\"(\\w+)\"")
	newRows := make([]string, 0)
	for _, row := range rows {
		newRows = append(newRows, fmt.Sprintf("%s // %s", row[0], ksRemarkMap[row[2]]))
	}
	res := strings.Join(newRows, "\n")

	err := ioutil.WriteFile("out/out.txt", []byte(res), fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

// kuaishou api 转 model
func Kuaishou() {
	fileName := "tabrow/input/kuaishou_api.txt"
	rows := tabrow.LinesToRowsByTab(tabrow.ReadFileLines(fileName))

	res := tabrow.FormatRows(rows, tabrow.KuaishouAPI)

	err := ioutil.WriteFile("out/out.txt", []byte(res), fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
