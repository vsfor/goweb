package main

import (
	"fmt"
	"net/http"
)

func index (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "nothing ...")
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr: "0.0.0.0:8181",
		Handler: mux,
	}
	server.ListenAndServe()
}
