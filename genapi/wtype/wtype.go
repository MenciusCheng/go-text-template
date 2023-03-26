package wtype

import (
	"fmt"
	"github.com/MenciusCheng/go-text-template/parse"
	"strings"
	"text/template"
)

type WType struct {
	Blocks []Block `json:"blocks"`

	BlockByID    map[string]*Block
	BlockByName  map[string]*Block
	BlockByType  map[string]*Block
	BlocksByName map[string][]*Block
	BlocksByType map[string][]*Block
}

func (t *WType) GenMap() {
	t.initMap()

	for i := range t.Blocks {
		v := &t.Blocks[i]
		t.BlockByID[t.Blocks[i].Def.ID] = v
		t.BlockByName[t.Blocks[i].Def.Name] = v
		t.BlockByType[t.Blocks[i].Def.TextType] = v
		t.BlocksByName[t.Blocks[i].Def.Name] = append(t.BlocksByName[t.Blocks[i].Def.Name], v)
		t.BlocksByType[t.Blocks[i].Def.TextType] = append(t.BlocksByType[t.Blocks[i].Def.TextType], v)
	}
}

func (t *WType) initMap() {
	t.BlockByID = make(map[string]*Block)
	t.BlockByName = make(map[string]*Block)
	t.BlockByType = make(map[string]*Block)
	t.BlocksByName = make(map[string][]*Block)
	t.BlocksByType = make(map[string][]*Block)
}

func NewWType() WType {
	wt := WType{}
	wt.initMap()
	return wt
}

func (t *WType) GenByTmpl() string {
	block, ok := t.BlockByType[TextTypeTmpl]
	if !ok {
		return ""
	}

	tplStr, ok := block.Content.(string)
	if !ok {
		return ""
	}

	tpl, err := template.New("").Funcs(parse.GetFuncMap()).Parse(tplStr)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return ""
	}

	sb := strings.Builder{}
	err = tpl.Execute(&sb, t)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return ""
	}
	return sb.String()
}
