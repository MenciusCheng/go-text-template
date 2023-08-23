package main

import (
	"github.com/MenciusCheng/go-text-template/http/handler"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
)

func main() {
	go func() {
		// pprof
		http.ListenAndServe(":6060", nil)
	}()
	go func() {
		// echo server
		muxEcho := http.NewServeMux()
		muxEcho.HandleFunc("/", handler.EchoHandler)
		http.ListenAndServe(":8888", muxEcho)
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handler.Ping)
	mux.HandleFunc("/gen", handler.GenHandler)
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
