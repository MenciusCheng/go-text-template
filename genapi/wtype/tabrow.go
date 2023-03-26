package wtype

import "strings"

type Tabrow [][]string

func ParseToTabrow(str string) Tabrow {
	rows := Tabrow{}
	lines := strings.Split(strings.TrimSpace(str), "\n")
	for _, line := range lines {
		row := []string{line}
		subs := strings.Split(strings.TrimSpace(line), "\t")
		for _, sub := range subs {
			row = append(row, strings.TrimSpace(sub))
		}
		rows = append(rows, row)
	}
	return rows
}
