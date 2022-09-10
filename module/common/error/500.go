package error

import (
	"fmt"
	"net/http"
)

type Err500 struct {}

func (e Err500) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	fmt.Fprintln(w, "system error")
}