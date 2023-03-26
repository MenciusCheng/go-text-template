package parser

import (
	"strings"
)

type TabRow struct {
	Rows []Row `json:"rows"`
}

type Row struct {
	Origin string   `json:"origin"` // 原始内容
	Cells  []string `json:"cells"`  // 以\t分隔的每一单元格
}

func ParseToTabRow(str string) TabRow {
	rows := make([]Row, 0)
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		row := Row{Origin: line}

		subs := strings.Split(strings.TrimSpace(line), "\t")
		cells := make([]string, 0, len(subs))
		for _, sub := range subs {
			cells = append(cells, strings.TrimSpace(sub))
		}
		row.Cells = cells

		rows = append(rows, row)
	}
	return TabRow{Rows: rows}
}
