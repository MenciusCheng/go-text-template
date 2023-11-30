package exec_stat_data

import (
	"fmt"
	"github.com/MenciusCheng/go-text-template/genapi-v2/generator"
	"testing"
)

func TestGenerator_StatEventParamTab(t *testing.T) {
	g := generator.NewGenerator(generator.ConfigParser(generator.ParserTabRow))
	// 文本解析
	g.Source(StatEventParamTab.Data)

	// 模版添加
	err := g.Temp(StatEventParamTab.Tmpl)
	if err != nil {
		t.Error(err)
		return
	}

	// 生成结果
	fmt.Println(g.Exec())
}

func TestGenerator_StatEventParamJson(t *testing.T) {
	g := generator.NewGenerator(generator.ConfigParser(generator.ParserJson))
	// 文本解析
	g.Source(StatEventParamJson.Data)

	// 模版添加
	err := g.Temp(StatEventParamJson.Tmpl)
	if err != nil {
		t.Error(err)
		return
	}

	// 生成结果
	fmt.Println(g.Exec())
}
