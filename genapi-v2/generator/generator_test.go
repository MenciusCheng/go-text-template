package generator

import "testing"

// 生成文本示例
func TestGenerator_Exec(t *testing.T) {
	g := NewGenerator(ParserTabRow)
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
	g := NewGenerator(ParserTabRow)
	// 文本解析
	g.SourceFile("source.txt")
	t.Log(g.JsonIndent())

	// 模版添加
	err := g.TempFile("source.tmpl")
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
