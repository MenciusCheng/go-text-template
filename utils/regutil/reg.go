package regutil

import (
	"encoding/json"
	"regexp"
)

type RegRes struct {
	Expr string `json:"expr"`
	S    string `json:"s"`

	MatchString                bool       `json:"matchString"`
	FindString                 string     `json:"findString"`
	FindStringIndex            []int      `json:"findStringIndex"`
	FindStringSubmatch         []string   `json:"findStringSubmatch"`
	FindStringSubmatchIndex    []int      `json:"findStringSubmatchIndex"`
	FindAllString              []string   `json:"findAllString"`
	FindAllStringIndex         [][]int    `json:"findAllStringIndex"`
	FindAllStringSubmatch      [][]string `json:"findAllStringSubmatch"`
	FindAllStringSubmatchIndex [][]int    `json:"findAllStringSubmatchIndex"`
}

func (r *RegRes) String() string {
	bs, _ := json.MarshalIndent(r, "", "\t")
	return string(bs)
}

func ShowRegRes(expr string, s string) (RegRes, error) {

	reg, err := regexp.Compile(expr)
	if err != nil {
		return RegRes{}, err
	}
	//reg := regexp.MustCompile(expr)

	res := RegRes{
		Expr:                       expr,
		S:                          s,
		MatchString:                reg.MatchString(s),
		FindString:                 reg.FindString(s), // 匹配一次
		FindStringIndex:            reg.FindStringIndex(s),
		FindStringSubmatch:         reg.FindStringSubmatch(s), // 匹配一次和括号中
		FindStringSubmatchIndex:    reg.FindStringSubmatchIndex(s),
		FindAllString:              reg.FindAllString(s, -1), // 全局匹配
		FindAllStringIndex:         reg.FindAllStringIndex(s, -1),
		FindAllStringSubmatch:      reg.FindAllStringSubmatch(s, -1), // 全局匹配和括号中
		FindAllStringSubmatchIndex: reg.FindAllStringSubmatchIndex(s, -1),
	}
	return res, nil
}
