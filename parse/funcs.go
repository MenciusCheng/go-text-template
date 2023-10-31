package parse

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"unicode"
)

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"IsInStringSlice":  IsInStringSlice,
		"IfElse":           IfElse,
		"UpperFirst":       UpperFirst,
		"UpperFirst2":      UpperFirst2,
		"LowerFirst":       LowerFirst,
		"SnakeCase":        SnakeCase,
		"SnakeToCamel":     SnakeToCamel,
		"FloatToIntString": FloatToIntString,
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

func UpperFirst(s string) string {
	if len(s) > 0 {
		return fmt.Sprintf("%s%s", strings.ToUpper(string(s[0])), s[1:])
	}
	return s
}

func UpperFirst2(s string) string {
	if s == "id" {
		return "ID"
	}
	return UpperFirst(s)
}

func LowerFirst(s string) string {
	if len(s) > 0 {
		return fmt.Sprintf("%s%s", strings.ToLower(string(s[0])), s[1:])
	}
	return s
}

func SnakeCase(str string) string {
	d := byte('a' - 'A')
	sb := strings.Builder{}
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			if i > 0 {
				sb.WriteByte('_')
			}
			sb.WriteByte(str[i] + d)
		} else {
			sb.WriteByte(str[i])
		}
	}
	return sb.String()
}

func SnakeToCamel(s string) string {
	var buf bytes.Buffer
	upNext := false

	for i, c := range s {
		if i == 0 {
			buf.WriteRune(unicode.ToUpper(c))
		} else {
			if c == '_' {
				upNext = true
			} else if upNext {
				buf.WriteRune(unicode.ToUpper(c))
				upNext = false
			} else {
				buf.WriteRune(c)
			}
		}
	}

	return buf.String()
}

func FloatToIntString(f float64) string {
	return fmt.Sprintf("%d", int(f))
}
