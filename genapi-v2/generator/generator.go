package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Generator 生成器，路径如下：
//
//	文本 => 解析器(Parser) => Data(Json Map) + 模板 => 执行器(Executor) => 结果
type Generator struct {
	OriData  interface{}                              // 解释器读取文本后生成的数据
	Data     map[string]interface{}                   // OriData 序列化为JSON后，再反序列化成 Data
	Parser   func(text string) map[string]interface{} // 解释器
	TmplText string                                   // 模板
	Executor ExecutorFunc                             // 执行器
}

func NewGenerator(opts ...OptionFunc) *Generator {
	g := &Generator{
		Executor: TempExecutor,
	}
	for _, opt := range opts {
		opt(g)
	}
	return g
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

func (g *Generator) writeFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("WriteFile failed", err)
		return err
	}
	return nil
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

	g.TmplText = text
	return nil
}

// 从文件中读取模版
func (g *Generator) TempFile(filename string, opts ...OptionFunc) error {
	return g.Temp(g.loadFile(filename), opts...)
}

// 执行模版生成文本
func (g *Generator) Exec(opts ...OptionFunc) string {
	for _, opt := range opts {
		opt(g)
	}

	return g.Executor(g.Data, g.TmplText)
}

// 执行模版生成文本至文件
func (g *Generator) ExecToFile(filename string, opts ...OptionFunc) error {
	return g.writeFile(filename, g.Exec(opts...))
}

// 执行模版生成至详细日志
func (g *Generator) ExecToDebugLog(opts ...OptionFunc) error {
	// 创建目录
	dir := filepath.Join(".", "out")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	ts := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s/%s.out", dir, ts)

	var b bytes.Buffer

	b.WriteString("====================data====================\n")
	b.WriteString(g.JsonIndent())

	b.WriteString("\n\n====================tmpl====================\n")
	b.WriteString(g.TmplText)

	b.WriteString("\n\n====================res====================\n")
	res := g.Exec(opts...)
	b.WriteString(res)

	fmt.Println(res)

	return g.writeFile(filename, b.String())
}
