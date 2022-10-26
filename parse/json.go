package parse

import (
	"encoding/json"
	"fmt"
	"strings"
	"text/template"
)

func JsonMap(text string, data map[string]interface{}) (res string, err error) {
	tmpl, err := template.New("").Funcs(GetFuncMap()).Parse(text)
	if err != nil {
		return
	}
	// strings.Builder 转化字符串时比 bytes.Buffer 更快
	sb := strings.Builder{}
	err = tmpl.Execute(&sb, data)
	if err != nil {
		return
	}
	return sb.String(), nil
}

func StringToMap(jsonString string) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func PrintJsonRows(jsonString string) error {
	data, err := StringToMap(jsonString)
	if err != nil {
		return err
	}

	tmpl := `{{ range $i, $item := .rows }}
{{- $i }}. [{{ index . 0 }}] [{{ index . 1 }}]
{{ end -}}
`
	res, err := JsonMap(tmpl, data)
	if err != nil {
		return err
	}
	fmt.Printf("res: \n%+v\n", res)

	return nil
}
