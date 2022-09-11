package res

import (
	"html/template"
	"net/http"
)

func ShowLayout(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html","t2.html")
	t.ExecuteTemplate(w, "layout", "")
}