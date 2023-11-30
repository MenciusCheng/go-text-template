package exec_stat_data

type ExecData struct {
	Data string
	Tmpl string
}

// 事件tab
var StatEventParamTab = ExecData{
	Data: `
sex	性别	字符串	性别（1:男；2：女）
phone	手机号MD5	字符串	手机号MD5
`,
	Tmpl: `
{{- range $row := .rows }}
			{
				EventParamName:         "{{ index $row 2 }}",
				EventParamEnglishName:  "{{ index $row 1 }}",
				EventParamDesc:         "{{ index $row 4 }}",
				EventParamFormat:       {{ if eq (index $row 3) "字符串" }}record.ParamFormatStr{{ else if eq (index $row 3) "整型" }}record.ParamFormatInt64{{ else if eq (index $row 3) "浮点型" }}record.ParamFormatFloat64{{ end }}, // {{ index $row 3 }}
				EventParamUniqueSwitch: record.NoUnique,
			},
{{- end }}
`,
}

// 事件json
var StatEventParamJson = ExecData{
	Data: `
{
    "data": [
        {
            "name": "广告",
            "english_name": "ad",
            "desc": "创意名称",
            "param_format": 2,
            "unique_switch": 1
        }
    ]
}
`,
	Tmpl: `
{{- range $row := .data }}
			{
				EventParamName:         "{{ .name }}",
				EventParamEnglishName:  "{{ .english_name }}",
				EventParamDesc:         "{{ .desc }}",
				EventParamFormat:       {{ if eq .param_format 2.0 }}record.ParamFormatStr, // 字符串{{ else if eq .param_format 1.0 }}record.ParamFormatInt64, // 整型{{ else if eq .param_format 3.0 }}record.ParamFormatFloat64, // 浮点型{{ end }}
				EventParamUniqueSwitch: record.NoUnique,
			},
{{- end }}
`,
}
