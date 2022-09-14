package main

import (
	"encoding/json"
	"fmt"
	"github.com/MenciusCheng/go-text-template/parse"
	"net/http"
)

type GenObj struct {
	Template string `json:"template"`
	Data     string `json:"data"`
}

func genHandler(w http.ResponseWriter, r *http.Request) {
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

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/gen", genHandler)
	http.HandleFunc("/ping", Ping)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
