package error

import (
	"fmt"
	"net/http"
)

type Err404 struct {}

func (e Err404) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintln(w, "page not found")
}