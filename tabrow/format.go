package tabrow

import (
	"fmt"
	"github.com/MenciusCheng/go-text-template/parse"
	"strings"
)

type FormatFunc func(cells []string) string

func FormatRows(rows [][]string, f FormatFunc) string {
	b := strings.Builder{}
	for _, row := range rows {
		b.WriteString(f(row))
	}
	return b.String()
}

var ksapiMap = make(map[string]bool)

func KuaishouAPI(cells []string) string {
	if len(cells) < 3 {
		return ""
	}
	field := strings.TrimSpace(cells[0])

	if ksapiMap[field] {
		return ""
	}

	var t string
	switch strings.ToLower(strings.TrimSpace(cells[1])) {
	case "string":
		t = "string"
	case "int":
		t = "int64"
	case "long":
		t = "int64"
	case "double":
		t = "float64"
	case "-":
		return ""
	default:
		panic(fmt.Sprintf("未检查类型:'%s'", cells[1]))
	}
	camel := parse.SnakeToCamel(field)
	//if parse.SnakeCase(camel) != field {
	//	panic(fmt.Sprintf("驼峰转换失败:'%s' to '%s'", field, camel))
	//}
	remark := strings.TrimSpace(cells[2])

	ksapiMap[field] = true

	return fmt.Sprintf("%s %s `json:\"%s\"` // %s\n", camel, t, field, remark)
}
