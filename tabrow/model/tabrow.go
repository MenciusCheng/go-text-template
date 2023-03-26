package model

type TabRow struct {
	Rows []Row `json:"rows"`
}

type Row struct {
	Origin string   `json:"origin"` // 原始内容
	Cells  []string `json:"cells"`  // 以\t分隔的每一单元格
}
