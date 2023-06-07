package generate_method

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func WriteBean(beanMap map[string][]BeanColumn, config *Config) {

	for key, value := range beanMap {
		//log.Println(key, value)
		var b bytes.Buffer

		b.WriteString(`package  dbbean ` + " \n \n")
		b.WriteString(`import  ( "time" )  ` + " \n \n")
		b.WriteString("//这个文件是自动生成的，不要修改。当再次自动生成的时候，修改会被冲掉 \n \n")

		b.WriteString("type " + RemoveUnderscoreCapitalize(key) + " struct {")
		for _, v2 := range value {
			var dataType string
			var name string
			var jsonname string
			jsonname = RemoveUnderscoreLowercase(v2.Name)
			name = string(unicode.ToUpper(rune(jsonname[0]))) + jsonname[1:]
			//if name == "Id" {
			//	name = "ID"
			//}
			if strings.Contains(v2.Type, "char") || strings.Contains(v2.Type, "text") {
				dataType = "string"
			} else if strings.Contains(v2.Type, "int") {
				dataType = "int"
			} else {
				dataType = v2.Type
			}

			switch {

			case strings.Contains(v2.Type, "char") || strings.Contains(v2.Type, "text") || strings.Contains(v2.Type, "json"):
				dataType = "string"

			case strings.Contains(v2.Type, "int"):
				dataType = "int"

			case strings.Contains(v2.Type, "double"):
				dataType = "float64"

			//case strings.Contains(v2.Type, "timestamp"):
			//	dataType = "JsonTime"

			case strings.Contains(v2.Type, "timestamp") ||
				strings.Contains(v2.Type, "date"):
				dataType = "time.Time"

			default:
				dataType = v2.Type

			}

			//b.WriteString(name + ` ` + dataType + "`json:\"" + jsonname + "\"`" + ` // ` + v2.COLUMN_COMMENT + "\n")
			b.WriteString(name + ` ` + dataType + "`json:\"" + v2.Name + "\"`" + ` // ` + v2.COLUMN_COMMENT + "\n")
		}
		b.WriteString("}")
		b.WriteString(fmt.Sprintf(`
func (%s) TableName() string {
	return "%s"
}
`, RemoveUnderscoreCapitalize(key), key))

		// /Users/zhang/xy/zhi_dun/src/bean

		//writeFile(`../dbbean/`, key+`_generate.go`, b.Bytes(), true)
		writeFile(config.BeanDir, key+`_generate.go`, b.Bytes(), true)

	}

}
