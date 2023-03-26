package wtype

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
	"sync/atomic"
)

// 内容块
type Block struct {
	Def     BlockDef    `json:"def"`     // 格式定义
	Lines   []string    `json:"lines"`   // 读取的每一行内容
	Content interface{} `json:"content"` // 解析后的内容
}

type BlockDef struct {
	ID         string `json:"id"`          // 内容块唯一ID
	Name       string `json:"name"`        // 内容块名称，可以重复
	TextType   string `json:"text_type"`   // 文本格式类型
	StructType string `json:"struct_type"` // 结构体类型
	File       string `json:"file"`        // 文件名
}

type BlockScanner struct {
	Def      BlockDef `json:"def"`      // 格式定义
	Lines    []string `json:"lines"`    // 读取的每一行内容
	Buffered bool     `json:"buffered"` // 是否有未输出的内容
}

func (s *BlockScanner) ReadLine(line string) {
	s.Buffered = true
	s.Lines = append(s.Lines, line)
}

func (s *BlockScanner) String() string {
	return strings.Join(s.Lines, "\n")
}

func (s *BlockScanner) ToBlock() Block {
	if !s.Buffered {
		return Block{}
	}

	block := Block{
		Def:   s.Def,
		Lines: s.Lines,
	}

	switch strings.ToLower(s.Def.TextType) {
	case TextTypeText:
		block.Content = s.String()
	case TextTypeTabrow:
		str := strings.TrimSpace(s.String())
		block.Content = ParseToTabrow(str)
	case TextTypeJson:
		content := make(map[string]interface{})
		err := json.Unmarshal([]byte(s.String()), &content)
		if err != nil {
			panic(err)
		}
		block.Content = content
	case TextTypeYaml:
		content := make(map[string]interface{})
		err := yaml.Unmarshal([]byte(s.String()), &content)
		if err != nil {
			panic(err)
		}
		block.Content = content
	case TextTypeTmpl:
		block.Content = s.String()
	}

	return block
}

var blockCount int32

func NewBlockScanner(line string) *BlockScanner {
	id := fmt.Sprintf("%d", atomic.AddInt32(&blockCount, 1))

	bs := &BlockScanner{
		Def: BlockDef{
			ID:       id, // 默认为自增id，可以在配置中指定
			TextType: TextTypeText,
		},
	}

	allKV := BlockDefKVReg.FindAllString(line, -1)
	for _, kv := range allKV {
		str := kv
		split := strings.Split(str, ":")
		if len(split) < 2 {
			continue
		}
		key := strings.ToLower(split[0])
		value := split[1]
		switch key {
		case DefFieldID:
			bs.Def.ID = value
		case DefFieldName:
			bs.Def.Name = value
		case DefFieldType:
			bs.Def.TextType = value
		case DefFieldStruct:
			bs.Def.StructType = value
		case DefFieldFile:
			bs.Def.File = value
		}
	}

	return bs
}

func IsBlockDefLine(line string) bool {
	return BlockDefReg.Match([]byte(line))
}

func DefaultBlockScanner() *BlockScanner {
	return NewBlockScanner("")
}
