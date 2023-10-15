package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MenciusCheng/go-text-template/parse"
	"text/template"
)

// Generator 生成器，路径如下：
//
//	文本 => 解析器 => Data(Json Map) => 模版 => 结果
type Generator struct {
	Data      map[string]interface{}
	Parser    func(text string) map[string]interface{}
	Templater *template.Template
}

func NewGenerator(parser func(text string) map[string]interface{}) *Generator {
	return &Generator{
		Data:      make(map[string]interface{}),
		Parser:    parser,
		Templater: template.New("").Funcs(parse.GetFuncMap()),
	}
}

// 读取文本
func (g *Generator) Source(text string) {
	oriData := g.Parser(text)
	// 序列化为JSON后再反序列化成 map
	data := MapToJsonToMap(oriData)
	g.Data = data
}

// 打印读取的文本Json
func (g *Generator) Json() string {
	dataJson, err := json.Marshal(g.Data)
	if err != nil {
		return ""
	}
	return string(dataJson)
}

// 打印读取的文本Json，带缩进
func (g *Generator) JsonIndent() string {
	dataJson, err := json.MarshalIndent(g.Data, "", "\t")
	if err != nil {
		return ""
	}
	return string(dataJson)
}

// 读取模版
func (g *Generator) Temp(text string) error {
	t, err := g.Templater.Parse(text)
	if err != nil {
		return err
	}
	g.Templater = t
	return nil
}

// 执行模版生成文本
func (g *Generator) Exec() string {
	var b bytes.Buffer
	err := g.Templater.Execute(&b, g.Data)
	if err != nil {
		fmt.Println("Execute failed", err)
		return ""
	}
	return b.String()
}
