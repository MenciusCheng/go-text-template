package handler

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/history":
		ShowReqHistoryHandler(w, r)
		return
	case "/favicon.ico":
		FaviconHandler(w, r)
		return
	}

	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Printf("DumpRequest error: %v\n", err)
	}
	log.Printf("request:\n%s\n", reqDump)

	reqEcho := ReqEcho{
		CreatedAt: time.Now(),
		Request:   string(reqDump),
	}
	ReqEchoList = append(ReqEchoList, reqEcho)

	w.Write(reqDump)
}

func ShowReqHistoryHandler(w http.ResponseWriter, r *http.Request) {
	sb := strings.Builder{}

	if len(ReqEchoList) == 0 {

		sb.WriteString("History is empty.\n")

		sb.WriteString("\nPlease Visit:\n\nhttp://localhost:8080/mock\n")
	}

	for i := len(ReqEchoList) - 1; i >= 0; i-- {
		sb.WriteString("----------------------------------\n")
		sb.WriteString(fmt.Sprintf("No.%d time: %s\n", i+1, ReqEchoList[i].CreatedAt.Format("2006-01-02 15:04:05.999")))
		sb.WriteString(fmt.Sprintf("request:\n%s\n", ReqEchoList[i].Request))
	}

	w.Write([]byte(sb.String()))
}

//go:embed favicon.ico
var favicon []byte

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(favicon)
}
