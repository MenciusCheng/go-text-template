package genstream

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenStream_ReadCsv(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			args: args{text: "1, John\n2, Jane\n3, Bob"},
			want: [][]string{{"1", " John"}, {"2", " Jane"}, {"3", " Bob"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GenStream{}
			g.ReadCsv(tt.args.text)
			if !reflect.DeepEqual(g.Data, tt.want) {
				t.Errorf("ReadCsv() = %v, want %v", g.Data, tt.want)
			}
		})
	}
}

func TestGenStream_WriteCsv(t *testing.T) {
	type fields struct {
		Data [][]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{
				Data: [][]string{
					{"7", " Sarah"},
					{"8", " Tom"},
					{"9", " Emma"},
				},
			},
			want: "7,\" Sarah\"\n8,\" Tom\"\n9,\" Emma\"\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GenStream{
				Data: tt.fields.Data,
			}
			if got := g.WriteCsv(); got != tt.want {
				t.Errorf("WriteCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 生成文本示例
func TestGenerator_Split(t *testing.T) {
	g := NewGenStream()
	// 文本读取
	g.Read(`1	2023-09-03 00:02:18	2023-09-03 00:02:18	天一说土被厂	http://nufp.ug/ovpqwbd	0
2	2023-09-03 00:10:54	2023-09-03 00:16:33	革深圆划织	http://mbrjhu.eh/lvrpsxfl	25
31	2023-09-03 00:55:20	2023-09-03 00:55:20	工工想力人位	http://wcdywv.cy/mtfa	3`)
	g.PrintJson()

	// 文本分割
	g.Split("\t")
	g.PrintJson()

	// 模版输出
	out := g.WriteTemp(`{{- range $i, $row := .Data }}
{{ $i }}: col count:{{ len ($row) }}, col3:{{ index $row 3 }}
{{- end }}`)
	fmt.Printf("out:%s\n", out)
}

// 行数分组
func TestGenerator_Chunk(t *testing.T) {
	g := NewGenStream()
	// 文本解析
	g.Read(`topic
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
	g.PrintJson()

	// 行数分组
	g.FilterBlank().Chunk(4)
	g.PrintJson()

	// 模版输出
	out := g.WriteTemp(`{{- range $row := .Data }}
[ {{ range $i, $col := . }}{{if gt $i 0}}, {{end}}{{$col}}{{end}} ]
{{- end }}`)
	fmt.Printf("out:%s\n", out)
}

// 文本替换
func TestGenerator_ReplaceAll(t *testing.T) {
	g := NewGenStream()
	g.ReadFile("sqltext/test.sql")
	g.ReplaceAll("abc", "fgh")
	g.WriteFile("sqltext/test.out")
}
