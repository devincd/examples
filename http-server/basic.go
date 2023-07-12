package main

import (
	"log"
	"net/http"

	"github.com/k0kubun/pp/v3"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", Ping)
	server := &http.Server{
		Addr:    "0.0.0.0:9080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to listen and server http server:%v", err)
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	pp.Println(r)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Fatalf("failed to listen and server http server:%v", err)
	}
}

/*
源码链路

*/
