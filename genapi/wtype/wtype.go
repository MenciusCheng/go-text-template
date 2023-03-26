package wtype

type WType struct {
	Blocks []Block `json:"blocks"`

	BlockByID    map[string]*Block
	BlockByName  map[string]*Block
	BlockByType  map[string]*Block
	BlocksByName map[string][]*Block
	BlocksByType map[string][]*Block
}

func (t *WType) GenMap() {
	for i := range t.Blocks {
		v := &t.Blocks[i]
		t.BlockByID[t.Blocks[i].Def.ID] = v
		t.BlockByName[t.Blocks[i].Def.Name] = v
		t.BlockByType[t.Blocks[i].Def.TextType] = v
		t.BlocksByName[t.Blocks[i].Def.Name] = append(t.BlocksByName[t.Blocks[i].Def.Name], v)
		t.BlocksByType[t.Blocks[i].Def.TextType] = append(t.BlocksByType[t.Blocks[i].Def.TextType], v)
	}
}

func NewWType() WType {
	return WType{
		BlockByID:    make(map[string]*Block),
		BlockByName:  make(map[string]*Block),
		BlockByType:  make(map[string]*Block),
		BlocksByName: make(map[string][]*Block),
		BlocksByType: make(map[string][]*Block),
	}
}
