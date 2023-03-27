package stat_dimension

import (
	_ "embed"
	"fmt"
	"github.com/MenciusCheng/go-text-template/genapi/wtype"
	"github.com/MenciusCheng/go-text-template/parse"
	"io/fs"
	"io/ioutil"
	"regexp"
	"strings"
)

//go:embed dimension.json
var DimensionJson string

//go:embed unit_report_resp_details.txt
var UnitReportRespDetailsTxt string

//go:embed unit_report_resp_details.tmpl
var UnitReportRespDetailsTmpl string

//go:embed unit_report_req_details.txt
var UnitReportReqDetailsTxt string

//go:embed map_express.txt
var MapExpressTxt string

// 参数类型
const (
	ParamFormatInt64   = int32(1)
	ParamFormatStr     = int32(2)
	ParamFormatFloat64 = int32(3)
)

var paramFormatName = map[int32]string{
	ParamFormatInt64:   "整型",
	ParamFormatStr:     "字符串",
	ParamFormatFloat64: "浮点型",
}

// 生成响应参数
func GenRespDetailField() {
	wt := wtype.NewWType()
	wt.ReadWTypeByStr(UnitReportRespDetailsTxt, func(b *wtype.Block) {
		b.Def.Name = "respTxt"
		b.Def.TextType = wtype.TextTypeTabrow
		b.FormatAfter = KuaishouApiField
	})
	wt.ReadWTypeByStr(UnitReportRespDetailsTmpl, func(b *wtype.Block) {
		b.Def.TextType = wtype.TextTypeTmpl
	})

	err := ioutil.WriteFile("out.txt", []byte(wt.GenByTmpl()), fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

// 生成请求参数
func GenReqDetailField() {
	wt := wtype.NewWType()
	wt.ReadWTypeByStr(UnitReportReqDetailsTxt, func(b *wtype.Block) {
		b.Def.Name = "respTxt"
		b.Def.TextType = wtype.TextTypeTabrow
		b.FormatAfter = KuaishouApiField
	})
	wt.ReadWTypeByStr(UnitReportRespDetailsTmpl, func(b *wtype.Block) {
		b.Def.TextType = wtype.TextTypeTmpl
	})

	err := ioutil.WriteFile("out.txt", []byte(wt.GenByTmpl()), fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

// 根据字段给表达式增加备注
func AddRemarkToExpByField() {
	wt := wtype.NewWType()
	wt.ReadWTypeByStr(DimensionJson, func(b *wtype.Block) {
		b.Def.TextType = wtype.TextTypeJson
	})
	jsonCnt := wt.CByType(wtype.TextTypeJson).([]map[string]interface{})
	remarkMapByField := make(map[string]string)
	for _, item := range jsonCnt {
		remark := fmt.Sprintf("%s，%s", item["name"].(string), paramFormatName[int32(item["param_format"].(float64))])
		remarkMapByField[item["english_name"].(string)] = remark
	}

	fieldReg := regexp.MustCompile(`"(\w+)":`)
	wt.ReadWTypeByStr(MapExpressTxt, func(b *wtype.Block) {
		b.Def.Name = "mexp"
		b.FormatAfter = func(b2 *wtype.Block) {
			arr := make([]string, 0)
			for _, line := range b2.Lines {
				submatch := fieldReg.FindStringSubmatch(line)
				v := fmt.Sprintf("%s // %s", line, remarkMapByField[submatch[1]])
				arr = append(arr, v)
				fmt.Println(v)
			}
			b2.Content = arr
		}
	})
}

func checkList() bool {
	wt := wtype.NewWType()
	wt.ReadWTypeByStr(DimensionJson, func(b *wtype.Block) {
		b.Def.TextType = wtype.TextTypeJson
	})
	content := wt.Blocks[0].Content.([]map[string]interface{})
	fmt.Printf("len: %d\n", len(content))
	for _, item := range content {
		if item["english_name"] != item["alias"] {
			return false
		}
	}
	return true
}

func KuaishouApiField(b *wtype.Block) {
	tabrow := b.Content.(wtype.Tabrow)
	var ksapiMap = make(map[string]bool)
	tabrow2 := wtype.Tabrow{}
	for _, row := range tabrow {
		if len(row) < 4 {
			continue
		}

		fieldName := strings.TrimSpace(row[1])
		if ksapiMap[fieldName] {
			continue
		}
		ksapiMap[fieldName] = true

		camel := parse.SnakeToCamel(fieldName)

		var t string
		switch strings.ToLower(strings.TrimSpace(row[2])) {
		case "string":
			t = "string"
		case "int":
			t = "int64"
		case "long":
			t = "int64"
		case "double":
			t = "float64"
		case "-":
			continue
		default:
			panic(fmt.Sprintf("未检查类型:'%s'", row[2]))
		}

		row2 := make([]string, 0)
		row2 = append(row2, row...)
		row2 = append(row2, camel, t)
		tabrow2 = append(tabrow2, row2)
	}
	b.Content = tabrow2
}
