package generate_method

import (
	"bytes"
	"fmt"
	"sort"
)

func WriteControllerRegister(beanMap map[string][]BeanColumn) {

	var (
		updateStr string
		keySlice  []string
	)

	for s := range beanMap {
		keySlice = append(keySlice, s)
	}

	sort.Strings(keySlice)

	for _, key := range keySlice {

		if _, ok := generateControllerTableMap[key]; !ok {
			continue
		}

		updateStr = updateStr + fmt.Sprintf(`"%s": &%s{"%s"}, `,
			RemoveUnderscoreCapitalize(key),
			RemoveUnderscoreCapitalize(key)+`Controller`,
			RemoveUnderscoreCapitalize(key),
		)

		updateStr = updateStr + "\n"
	}

	var template string

	template = fmt.Sprintf(`
package dbcontroller

var Registry = map[string]interface{}{
	%s
}
`,
		updateStr,
	)

	var b bytes.Buffer

	b.WriteString(template)

	writeFile(`../dbcontroller/`, `main_controller_generate.go`, b.Bytes(), true)

}
