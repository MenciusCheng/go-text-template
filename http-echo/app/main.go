package main

import (
	"fmt"
	"github.com/MenciusCheng/go-text-template/http-echo/handler"
	"github.com/MenciusCheng/go-text-template/utils/urlutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := "8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.EchoHandler)

	go func() {
		// service connections
		if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	openUrl := fmt.Sprintf("http://localhost:%s/history", port)
	echoUrl := fmt.Sprintf("http://localhost:%s/echo?name=cat", port)

	go func() {
		time.Sleep(time.Second)
		urlutil.OpenBrowser(openUrl)
	}()

	log.Printf("server started:\n%s\nechoUrl:\n%s\n", openUrl, echoUrl)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("server exit")
}
