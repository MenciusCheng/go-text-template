package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func MapToJsonToMap(data map[string]interface{}) map[string]interface{} {
	dataJson, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Marshal err", err)
		return nil
	}

	res := make(map[string]interface{})
	err = json.Unmarshal(dataJson, &res)
	if err != nil {
		fmt.Println("Unmarshal err", err)
		return nil
	}
	return res
}

func PrintMapToJson(data map[string]interface{}) string {
	dataJson, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Marshal err", err)
		return ""
	}

	var out bytes.Buffer
	json.Indent(&out, dataJson, "", "\t")
	// out.WriteTo(os.Stdout)
	return out.String()
}
