package parse

import "text/template"

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"IsInStringSlice": IsInStringSlice,
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
