package stat_script

import (
	_ "embed"
	"fmt"
	"github.com/MenciusCheng/go-text-template/genapi/wtype"
	"github.com/MenciusCheng/go-text-template/parse"
	"regexp"
)

//go:embed sv_data_Of_Micard_230410_all.json
var SvMicardJson string

//go:embed sv_data_Of_Micard_230410_all2.json
var SvMicardJson2 string

//go:embed x.json
var AnniversaryJson string

//go:embed a.sql
var ASQL string

// 参数类型
const (
	ParamFormatInt64   = int32(1)
	ParamFormatStr     = int32(2)
	ParamFormatFloat64 = int32(3)
	ParamFormatTime    = int32(4)
)

var paramFormatName = map[int32]string{
	ParamFormatInt64:   "整型",
	ParamFormatStr:     "字符串",
	ParamFormatFloat64: "浮点型",
	ParamFormatTime:    "时间",
}

// sv_data_Of_Micard_230410_all 导出正式版
func HandleMicardJson2() {
	wt := wtype.NewWType()
	wt.ReadWTypeByStr(SvMicardJson, func(b *wtype.Block) {
		b.Def.Name = "SvMicardJson"
		b.Def.TextType = wtype.TextTypeJson
	})
	jsonCnt := wt.CByName("SvMicardJson").([]map[string]interface{})
	ftMapByField := make(map[string]int32)
	for _, item := range jsonCnt {
		field := item["name"].(string)
		t := int32(item["dataFormatType"].(float64))
		ftMapByField[field] = t
	}

	wt.ReadWTypeByStr(SvMicardJson2, func(b *wtype.Block) {
		b.Def.Name = "SvMicardJson2"
		b.Def.TextType = wtype.TextTypeJson
	})
	jsonCnt2 := wt.CByName("SvMicardJson2").([]map[string]interface{})
	remarkMapByField := make(map[string]string)
	for _, item := range jsonCnt2 {
		field := item["sqlParamEnglishName"].(string)
		t := int32(item["paramFormat"].(float64))

		if t == 2 {
			if ftMapByField[field] == 1 {
				remarkMapByField[field] = fmt.Sprintf("toInt64OrZero(ifNull(%s,'')) as %s,", field, field)
			} else {
				remarkMapByField[field] = fmt.Sprintf("ifNull(%s,'') as %s,", field, field)
			}
		} else if t == 1 || t == 3 {
			remarkMapByField[field] = fmt.Sprintf("ifNull(%s,0) as %s,", field, field)
		} else if t == 4 {
			remarkMapByField[field] = fmt.Sprintf("ifNull(%s,'0001-01-01 00:00:00') as %s,", field, field)
		}
	}

	fieldReg := regexp.MustCompile(`^    (\w+)`)
	wt.ReadWTypeByStr(ASQL, func(b *wtype.Block) {
		b.Def.Name = "asql"
		b.FormatAfter = func(b2 *wtype.Block) {
			arr := make([]string, 0)
			for _, line := range b2.Lines {
				submatch := fieldReg.FindStringSubmatch(line)
				if len(submatch) > 1 {
					if v, ok := remarkMapByField[submatch[1]]; ok {
						arr = append(arr, v)
						fmt.Println(v)
					} else {
						arr = append(arr, line)
						fmt.Println(line)
					}
				} else {
					arr = append(arr, line)
					fmt.Println(line)
				}
			}
			b2.Content = arr
		}
	})
}

func HandleMicardJson4() {
	wt := wtype.NewWType()
	fieldReg := regexp.MustCompile(`^    (\w+)`)
	wt.ReadWTypeByStr(ASQL, func(b *wtype.Block) {
		b.Def.Name = "asql"
		b.FormatAfter = func(b2 *wtype.Block) {
			arr := make([]string, 0)
			for _, line := range b2.Lines {
				submatch := fieldReg.FindStringSubmatch(line)
				if len(submatch) > 1 {
					fmt.Printf("\"%s\",", submatch[1])
				}
			}
			b2.Content = arr
		}
	})
}

func HandleAnniversaryJson() {
	wt := wtype.NewWType()
	wt.ReadWTypeByStr(AnniversaryJson, func(b *wtype.Block) {
		b.Def.TextType = wtype.TextTypeJson
	})
	jsonCnt := wt.CByType(wtype.TextTypeJson).([]map[string]interface{})

	for _, item := range jsonCnt {
		fmt.Printf("ds = append(ds, GetValueInt64(data.%s))\n", parse.UpperFirst(item["name"].(string)))
	}
}
