package wtype

import "regexp"

// BlockDef字段
const (
	DefFieldID     = "id"
	DefFieldName   = "name"
	DefFieldType   = "type"
	DefFieldStruct = "struct"
	DefFieldFile   = "file"
)

// 解释器类型
const (
	TextTypeText   = "text"
	TextTypeTabrow = "tabrow"
	TextTypeJson   = "json"
	TextTypeYaml   = "yaml"
	TextTypeTmpl   = "tmpl"
)

var (
	BlockDefReg   = regexp.MustCompile(`^@=(\w+:\w*)?(,\w+:\w*)*$`)
	BlockDefKVReg = regexp.MustCompile(`\w+:\w+`)
)
