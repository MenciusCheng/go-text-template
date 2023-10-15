package generator

import (
	"bytes"
	"testing"
	"text/template"
)

func TestParserTabRow(t *testing.T) {
	text := `1	2023-09-03 00:02:18	2023-09-03 00:02:18	天一说土被厂	http://nufp.ug/ovpqwbd	0
2	2023-09-03 00:10:54	2023-09-03 00:16:33	革深圆划织	http://mbrjhu.eh/lvrpsxfl	25
31	2023-09-03 00:55:20	2023-09-03 00:55:20	工工想力人位	http://wcdywv.cy/mtfa	3`

	// 解析数据源
	oriData := ParserTabRow(text)
	// 序列化为JSON后再反序列化成 map
	data := MapToJsonToMap(oriData)
	// JSON 打印
	t.Log(PrintMapToJson(data))

	// 模版打印
	tmplText := `
{{- range $row := .rows }}
f: {{ len (index $row 2) }}, {{ $row }}
{{- end }}
`
	tmpl, _ := template.New("").Parse(tmplText)

	var b bytes.Buffer
	tmpl.Execute(&b, data)
	t.Log(b.String())
}

func TestParserLineGroupBy(t *testing.T) {
	text := `455092177

2

https://api.ads.heytapmobi.com/api/uploadActiveData

111.19.93.185

oaid1111

455092182

1

https://api.ads.heytapmobi.com/api/uploadActiveData

120.219.10.208

oaid1111`

	data := ParserLineGroupBy(text)
	t.Log(PrintMapToJson(data))
}

func TestParserSQL(t *testing.T) {
	text := "CREATE TABLE `bookmark` (\n  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',\n  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',\n  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',\n  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',\n  `url` varchar(5000) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '网址',\n  `folder_id` bigint NOT NULL DEFAULT '0' COMMENT '所属文件夹ID',\n  PRIMARY KEY (`id`)\n) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='书签';"

	data := ParserSQL(text)
	t.Log(PrintMapToJson(data))
}
