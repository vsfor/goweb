package tc

import (
	"fmt"
	"net/http"
)

type TaHandler struct {}

func (h TaHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "no such service,by handler")
}

func WriteHeader(w http.ResponseWriter,r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "no such service,by fun")
}