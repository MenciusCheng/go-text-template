package generator

import (
	"bytes"
	"github.com/MenciusCheng/go-text-template/parse"
	"log"
	"strings"
	"text/template"
)

type ExecutorFunc func(data map[string]interface{}, tmplText string) string

func defaultTemplate() *template.Template {
	return template.New("").Funcs(parse.GetFuncMap())
}

// 内置模板执行器
func TempExecutor(data map[string]interface{}, tmplText string) string {
	t, err := defaultTemplate().Parse(tmplText)
	if err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer
	err = t.Execute(&b, data)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return b.String()
}

func WithTempExecutor(t *template.Template) ExecutorFunc {
	return func(data map[string]interface{}, tmplText string) string {
		if t == nil {
			t = defaultTemplate()
		}
		var err error

		t, err = t.Parse(tmplText)
		if err != nil {
			log.Fatal(err)
		}

		var b bytes.Buffer
		err = t.Execute(&b, data)
		if err != nil {
			log.Fatal(err)
			return ""
		}
		return b.String()
	}
}

// 逐行格式化模板的执行器，注入格式化函数
func WithLineExecutor(f func(data map[string]interface{}) func(line string) string) ExecutorFunc {
	return func(data map[string]interface{}, tmplText string) string {
		// 结合读取的数据生成格式化函数
		format := f(data)

		tmplLines := strings.Split(tmplText, "\n")
		ds := make([]string, 0, len(tmplLines))
		for _, tmplLine := range tmplLines {
			// 逐行格式化模板数据
			s := format(tmplLine)
			ds = append(ds, s)
		}
		return strings.Join(ds, "\n")
	}
}
