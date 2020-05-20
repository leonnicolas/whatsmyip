package main

import (
	"net/http"
	"strings"
)

func ipHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(strings.Split(req.RemoteAddr, ":")[0] + "\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ipHandler)
	http.ListenAndServe(":8080", mux)
}
