package main

import (
	"net/http"
	"strings"
)

func ipHandler(w http.ResponseWriter, req *http.Request) {
	addrs := strings.Split(strings.TrimSpace(req.Header.Get("X-Real-IP")), ",")
	if len(addrs) == 1 && !strings.Contains(addrs[0], ".") {
		w.Write([]byte(strings.Split(req.RemoteAddr, ":")[0] + "\n"))
		return
	}
	w.Write([]byte(addrs[len(addrs)-1] + "\n"))
	return
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ipHandler)
	http.ListenAndServe(":8080", mux)
}
