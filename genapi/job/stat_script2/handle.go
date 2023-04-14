package stat_script

import (
	_ "embed"
	"fmt"
	"github.com/MenciusCheng/go-text-template/genapi/wtype"
	"regexp"
)

//go:embed sv_user_anni.json
var AnniversaryJson string

//go:embed sv_user_anni2.json
var AnniversaryJson2 string

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

// anni 导出正式版
func HandleMicardJson2() {
	wt := wtype.NewWType()
	wt.ReadWTypeByStr(AnniversaryJson, func(b *wtype.Block) {
		b.Def.Name = "AnniversaryJson"
		b.Def.TextType = wtype.TextTypeJson
	})
	jsonCnt := wt.CByName("AnniversaryJson").([]map[string]interface{})
	ftMapBySqlField := make(map[string]int32)
	for _, item := range jsonCnt {
		field := item["name"].(string)
		t := int32(item["dataFormatType"].(float64))
		ftMapBySqlField[field] = t
	}

	wt.ReadWTypeByStr(AnniversaryJson2, func(b *wtype.Block) {
		b.Def.Name = "AnniversaryJson2"
		b.Def.TextType = wtype.TextTypeJson
	})
	jsonCnt2 := wt.CByName("AnniversaryJson2").([]map[string]interface{})
	remarkMapByField := make(map[string]string)
	for _, item := range jsonCnt2 {
		ckSQLfield := item["sqlParamEnglishName"].(string)
		ckField := item["paramEnglishName"].(string)
		t := int32(item["paramFormat"].(float64))

		if ckField == "updated_at" || ckField == "create_at" {
			remarkMapByField[ckSQLfield] = fmt.Sprintf("formatDateTime(ifNull(%s,'0001-01-01 00:00:00'),'%s') as %s,", ckField, "%Y-%m-%d %H:%M:%S", ckSQLfield)
		} else if t == 2 {
			if ftMapBySqlField[ckSQLfield] == 1 {
				remarkMapByField[ckSQLfield] = fmt.Sprintf("toInt64OrZero(ifNull(%s,'')) as %s,", ckField, ckSQLfield)
			} else {
				remarkMapByField[ckSQLfield] = fmt.Sprintf("ifNull(%s,'') as %s,", ckField, ckSQLfield)
			}
		} else if t == 1 || t == 3 {
			remarkMapByField[ckSQLfield] = fmt.Sprintf("ifNull(%s,0) as %s,", ckField, ckSQLfield)
		} else if t == 4 {
			remarkMapByField[ckSQLfield] = fmt.Sprintf("ifNull(%s,'0001-01-01 00:00:00') as %s,", ckField, ckSQLfield)
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
