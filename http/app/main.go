package main

import (
	"encoding/json"
	"fmt"
	"github.com/MenciusCheng/go-text-template/parse"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
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
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/gen", genHandler)
	mux.HandleFunc("/ping", Ping)
	InitPProf(mux)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

func InitPProf(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
}
