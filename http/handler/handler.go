package handler

import (
	"encoding/json"
	"fmt"
	"github.com/MenciusCheng/go-text-template/parse"
	"net/http"
	"net/http/httputil"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

type GenObj struct {
	Template string `json:"template"`
	Data     string `json:"data"`
}

// 生成代码
func GenHandler(w http.ResponseWriter, r *http.Request) {
	var req GenObj
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := parse.StringToMap(req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := parse.JsonMap(req.Template, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", res)
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Printf("DumpRequest error: %v\n", err)
	}

	fmt.Printf("%s request:\n%s\n", time.Now().Format("2006-01-02 15:04:05"), reqDump)

	w.Write(reqDump)
}
