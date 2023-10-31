package generator

import (
	"encoding/json"
	"fmt"
	"github.com/MenciusCheng/go-text-template/parse"
	"regexp"
	"testing"
	"text/template"
)

// 生成文本示例
func TestGenerator_Exec(t *testing.T) {
	g := NewGenerator(ConfigParser(ParserTabRow))
	// 文本解析
	g.Source(`1	2023-09-03 00:02:18	2023-09-03 00:02:18	天一说土被厂	http://nufp.ug/ovpqwbd	0
2	2023-09-03 00:10:54	2023-09-03 00:16:33	革深圆划织	http://mbrjhu.eh/lvrpsxfl	25
31	2023-09-03 00:55:20	2023-09-03 00:55:20	工工想力人位	http://wcdywv.cy/mtfa	3`)
	t.Log(g.JsonIndent())

	// 模版添加
	err := g.Temp(`{{- range $row := .rows }}
f: {{ len (index $row 2) }}, {{ $row }}
{{- end }}`)
	if err != nil {
		t.Error(err)
		return
	}

	// 生成结果
	t.Log(g.Exec())
}

// 从文件中生成文本示例
func TestGenerator_Exec_FromFile(t *testing.T) {
	g := NewGenerator()
	// 文本解析
	g.SourceFile("exec-tmpl-data/source.txt", ConfigParser(ParserTabRow))
	t.Log(g.JsonIndent())

	// 模版添加
	err := g.TempFile("exec-tmpl-data/source.tmpl", ConfigExecutor(WithTempExecutor(template.New("").Funcs(parse.GetFuncMap()))))
	if err != nil {
		t.Error(err)
		return
	}

	// 生成结果
	t.Log(g.Exec())
	err = g.ExecToFile("source.out")
	if err != nil {
		t.Error(err)
		return
	}
}

// 行数分组
func TestGenerator_Exec_GroupBy(t *testing.T) {
	g := NewGenerator(ConfigParser(WithParserLineGroupByCount(4)))
	// 文本解析
	g.Source(`topic
t_partition
t_offset
status
stat_api_tracking_topic_beta
4
8296139
重复xyreqid
stat_api_tracking_topic_beta
4
8296140
重复xyreqid
`)
	t.Log(g.JsonIndent())

	// 模版添加
	err := g.Temp(`{{- range $row := .rows }}
[ {{ range $i, $col := . }}{{if gt $i 0}}, {{end}}{{$col}}{{end}} ]
{{- end }}`)
	if err != nil {
		t.Error(err)
		return
	}

	// 生成结果
	t.Log(g.Exec())
}

// 从Json中生成文本示例
func TestGenerator_Exec_Json(t *testing.T) {
	g := NewGenerator()
	// 文本解析
	g.SourceFile("exec-tmpl-data/source.json", ConfigParser(ParserJson))

	// 模版添加
	err := g.Temp(`
{{- range $row := .result.data }}
 {{ .creative_type }} {{ .creative_style }} {{ .unit_name }} {{ .create_time }} {{ FloatToIntString .unit_id }}
{{- end }}
`)
	if err != nil {
		t.Error(err)
		return
	}

	// 生成结果
	t.Log(g.Exec())
}

func TestGenerator_LineExecutor(t *testing.T) {
	g := NewGenerator()

	// 文本解析
	g.SourceFile("exec-line-data/source.txt", ConfigParser(WithParserTabRowBySep(" ")))

	// 模版添加
	err := g.TempFile("exec-line-data/source.tmpl")
	if err != nil {
		t.Error(err)
		return
	}

	// 自定义执行器
	executor := WithLineExecutor(func(data map[string]interface{}) func(line string) string {
		remarkMap := make(map[string]string)
		bs, _ := json.Marshal(data["rows"])
		rows := make([][]string, 0)
		_ = json.Unmarshal(bs, &rows)

		for _, row := range rows {
			if len(row) == 3 {
				remarkMap[row[1]] = row[2]
			} else if len(row) == 4 {
				remarkMap[row[1]] = row[3]
			}
		}

		reg := regexp.MustCompile(`json:"([a-zA-Z0-9_]+)"`)

		return func(line string) string {
			submatch := reg.FindStringSubmatch(line)
			if len(submatch) < 2 {
				return line
			}
			if v, ok := remarkMap[submatch[1]]; ok {
				return fmt.Sprintf("%s // %s", line, v)
			}
			return line
		}
	})

	// 生成结果
	err = g.ExecToFile("source.out", ConfigExecutor(executor))
	if err != nil {
		t.Error(err)
		return
	}
}
