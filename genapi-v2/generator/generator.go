package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MenciusCheng/go-text-template/parse"
	"os"
	"text/template"
)

// Generator 生成器，路径如下：
//
//	文本 => 解析器 => Data(Json Map) => 模版 => 结果
type Generator struct {
	OriData  interface{}
	Data     map[string]interface{} // OriData 序列化为JSON后，再反序列化成 Data
	Parser   func(text string) map[string]interface{}
	Template *template.Template
}

func NewGenerator(opts ...OptionFunc) *Generator {
	g := &Generator{}
	for _, opt := range opts {
		opt(g)
	}
	return g
}

func defaultTemplate() *template.Template {
	return template.New("").Funcs(parse.GetFuncMap())
}

// 读取文本
func (g *Generator) Source(text string, opts ...OptionFunc) {
	for _, opt := range opts {
		opt(g)
	}

	oriData := g.Parser(text)
	g.OriData = oriData
	// 序列化为JSON后再反序列化成 map
	g.Data = MapToJsonToMap(oriData)
}

// 从文件中读取文本
func (g *Generator) SourceFile(filename string, opts ...OptionFunc) {
	g.Source(g.loadFile(filename), opts...)
}

func (g *Generator) loadFile(filename string) string {
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile failed", err)
		return ""
	}
	return string(b)
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
func (g *Generator) Temp(text string, opts ...OptionFunc) error {
	for _, opt := range opts {
		opt(g)
	}

	if g.Template == nil {
		g.Template = defaultTemplate()
	}

	t, err := g.Template.Parse(text)
	if err != nil {
		return err
	}
	g.Template = t
	return nil
}

// 从文件中读取模版
func (g *Generator) TempFile(filename string, opts ...OptionFunc) error {
	return g.Temp(g.loadFile(filename), opts...)
}

// 执行模版生成文本
func (g *Generator) Exec() string {
	var b bytes.Buffer
	err := g.Template.Execute(&b, g.Data)
	if err != nil {
		fmt.Println("Execute failed", err)
		return ""
	}
	return b.String()
}

// 执行模版生成文本至文件
func (g *Generator) ExecToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = g.Template.Execute(file, g.Data)
	if err != nil {
		fmt.Println("Execute Error", err)
		return err
	}

	return nil
}
