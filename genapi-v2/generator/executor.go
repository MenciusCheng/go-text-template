package generator

import (
	"bytes"
	"github.com/MenciusCheng/go-text-template/parse"
	"log"
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
