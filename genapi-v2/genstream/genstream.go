package genstream

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/MenciusCheng/go-text-template/parse"
	"github.com/MenciusCheng/go-text-template/utils/fileutil"
	"github.com/elliotchance/pie/v2"
	"log"
	"os"
	"strings"
	"text/template"
)

// GenStream 流式生成器
type GenStream struct {
	Data [][]string
}

func NewGenStream() *GenStream {
	g := &GenStream{}
	return g
}

// 读取文本
func (g *GenStream) Read(text string) *GenStream {
	lines := strings.Split(text, "\n")
	g.readLines(lines)
	return g
}

// 读取文件
func (g *GenStream) ReadFile(filename string) *GenStream {
	lines, err := fileutil.ReadFileByLine(filename)
	if err != nil {
		log.Fatal(err)
	}
	g.readLines(lines)
	return g
}

func (g *GenStream) readLines(lines []string) {
	records := make([][]string, 0)
	for _, line := range lines {
		records = append(records, []string{line})
	}
	g.Data = records
}

func (g *GenStream) writeLines() []string {
	return pie.Flat(g.Data)
}

func (g *GenStream) Write() string {
	return strings.Join(g.writeLines(), "\n")
}

func (g *GenStream) writeFile(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *GenStream) WriteFile(filename string) {
	g.writeFile(filename, g.Write())
}

// 读取CSV文本
func (g *GenStream) ReadCsv(text string) {
	r := csv.NewReader(strings.NewReader(text))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	g.Data = records
}

// 读取CSV文件
func (g *GenStream) ReadFileCsv(filename string) {
	records, err := fileutil.ReadCsvFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	g.Data = records
}

// 写入CSV文本
func (g *GenStream) WriteCsv() string {
	var sb strings.Builder
	w := csv.NewWriter(&sb)
	for _, record := range g.Data {
		err := w.Write(record)
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
	return sb.String()
}

// 写入CSV文件
func (g *GenStream) WriteFileCsv(filename string) {
	g.writeFile(filename, g.WriteCsv())
}

func (g *GenStream) PrintJson() {
	fmt.Println("g.Data:")
	for _, item := range g.Data {
		dataJson, _ := json.Marshal(item)
		fmt.Printf("%s\n", string(dataJson))
	}
}

func (g *GenStream) WriteTemp(tmpl string) string {
	t, err := template.New("").Funcs(parse.GetFuncMap()).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer
	err = t.Execute(&b, g)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return b.String()
}

func (g *GenStream) WriteFileTemp(tmpl string, filename string) {
	g.writeFile(filename, g.WriteTemp(tmpl))
}

func (g *GenStream) Split(sep string) *GenStream {
	records := make([][]string, 0)
	for _, item := range g.Data {
		records = append(records, strings.Split(item[0], sep))
	}
	g.Data = records
	return g
}

func (g *GenStream) Join(sep string) *GenStream {
	records := make([][]string, 0)
	for _, item := range g.Data {
		records = append(records, []string{strings.Join(item, sep)})
	}
	g.Data = records
	return g
}

func (g *GenStream) Chunk(chunkLength int) *GenStream {
	// 先转为一维数组，再转为二维数组
	g.Data = pie.Chunk(pie.Flat(g.Data), chunkLength)
	return g
}

func (g *GenStream) Filter(condition func(item []string) bool) *GenStream {
	records := make([][]string, 0)
	for _, item := range g.Data {
		if condition(item) {
			records = append(records, item)
		}
	}
	g.Data = records
	return g
}

func (g *GenStream) FilterBlank() *GenStream {
	return g.Filter(func(item []string) bool {
		for _, s := range item {
			if len(s) > 0 {
				return true
			}
		}
		return false
	})
}
