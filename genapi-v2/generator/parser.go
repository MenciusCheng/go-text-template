package generator

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func ParserTabRow(text string) map[string]interface{} {
	res := make(map[string]interface{})

	rows := make([][]string, 0)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		// 清洗
		lineData := strings.TrimSpace(line)
		if len(lineData) == 0 {
			continue
		}

		// 解析
		cols := strings.Split(lineData, "\t")
		values := make([]string, 0)
		values = append(values, lineData)
		values = append(values, cols...)
		rows = append(rows, values)
	}

	res["rows"] = rows
	return res
}

// 按行和列分组，自定义列的分隔符
func WithParserTabRowBySep(sep string) func(text string) map[string]interface{} {
	return func(text string) map[string]interface{} {
		res := make(map[string]interface{})

		rows := make([][]string, 0)
		lines := strings.Split(text, "\n")
		for _, line := range lines {
			// 清洗
			lineData := strings.TrimSpace(line)
			if len(lineData) == 0 {
				continue
			}

			// 解析
			cols := strings.Split(lineData, sep)
			values := make([]string, 0)
			values = append(values, lineData)
			values = append(values, cols...)
			rows = append(rows, values)
		}

		res["rows"] = rows
		return res
	}
}

func ParserLineGroupBy(text string) map[string]interface{} {
	res := make(map[string]interface{})

	rows := make([][]string, 0)
	values := make([]string, 0)

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		// 清洗
		lineData := strings.TrimSpace(line)
		if len(lineData) == 0 {
			continue
		}

		// 解析
		values = append(values, lineData)
		if len(values) >= 5 {
			rows = append(rows, values)
			values = make([]string, 0)
		}
	}
	if len(values) >= 5 {
		rows = append(rows, values)
	}

	res["rows"] = rows
	return res
}

// 根据行数分组，自定义行数
func WithParserLineGroupByCount(count int) func(text string) map[string]interface{} {
	return func(text string) map[string]interface{} {
		res := make(map[string]interface{})

		rows := make([][]string, 0)
		values := make([]string, 0)

		lines := strings.Split(text, "\n")
		for _, line := range lines {
			// 清洗
			lineData := strings.TrimSpace(line)
			if len(lineData) == 0 {
				continue
			}

			// 解析
			values = append(values, lineData)
			if len(values) >= count {
				rows = append(rows, values)
				values = make([]string, 0)
			}
		}
		if len(values) >= count {
			rows = append(rows, values)
		}

		res["rows"] = rows
		return res
	}
}

func ParserSQL(text string) map[string]interface{} {
	res := make(map[string]interface{})

	rows := make([]interface{}, 0)

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		// 清洗
		lineData := strings.TrimSpace(line)
		if len(lineData) == 0 {
			continue
		}

		// 解析
		headReg := regexp.MustCompile("^\\s*CREATE\\s*TABLE\\s*([a-zA-Z0-9_`]+)\\s*\\($")
		fieldReg := regexp.MustCompile("^\\s*([a-zA-Z0-9_`]+)")
		footReg := regexp.MustCompile("^\\s*\\)\\s*ENGINE")

		switch {
		case headReg.MatchString(lineData):
			fmt.Println("head:", lineData)
			submatch := headReg.FindStringSubmatch(lineData)
			tableName := strings.Trim(submatch[1], "`")
			res["table"] = tableName
		case fieldReg.MatchString(lineData):
			fmt.Println("filed:", lineData)
			submatch := fieldReg.FindStringSubmatch(lineData)
			fieldOrg := submatch[1]
			if fieldOrg == "PRIMARY" {
				continue
			} else if fieldOrg == "UNIQUE" {
				continue
			}
			fieldName := strings.Trim(fieldOrg, "`")

			fieldTypeReg := regexp.MustCompile(fmt.Sprintf("%s\\s+([a-zA-Z]+)", fieldOrg))
			fieldType := fieldTypeReg.FindStringSubmatch(lineData)[1]

			commentReg := regexp.MustCompile("COMMENT '(.+)'")
			comment := commentReg.FindStringSubmatch(lineData)[1]

			rows = append(rows, map[string]string{
				"name":    fieldName,
				"type":    fieldType,
				"comment": comment,
			})
		case footReg.MatchString(lineData):
			fmt.Println("foot:", lineData)
			res["rows"] = rows
		default:
			fmt.Println("not case:", lineData)
			continue
		}
	}

	return res
}

func ParserJson(text string) map[string]interface{} {
	res := make(map[string]interface{})
	err := json.Unmarshal([]byte(text), &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
