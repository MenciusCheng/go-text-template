package demo

import (
	"os"
	"strings"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

// 解析对象示例
func PrintInventory(inventory Inventory) error {
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}\n")
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, inventory)
	if err != nil {
		return err
	}
	return nil
}

// 解析 map
func ParseMap(text string, data map[string]interface{}) (res string, err error) {
	tmpl, err := template.New("").Parse(text)
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

// 移除文本空格
func TextAndSpaces(text string) (string, error) {
	return ParseMap(text, nil)
}
