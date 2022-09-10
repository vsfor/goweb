package res

import "net/http"

func Location302(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(302)
}
