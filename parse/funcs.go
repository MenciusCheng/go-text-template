package parse

import "text/template"

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"IsInStringSlice": IsInStringSlice,
		"IfElse":          IfElse,
	}
}

func IsInStringSlice(el string, arr []interface{}) bool {
	for _, s := range arr {
		if el == s {
			return true
		}
	}
	return false
}

func IfElse(a bool, b, c interface{}) interface{} {
	if a {
		return b
	}
	return c
}
