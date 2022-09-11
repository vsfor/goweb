package res

import (
	"html/template"
	"net/http"
)

type TplData struct {
	Info1 string
	Data2 string
	Dbool bool
	Dhtml string
}

func ShowTpl(w http.ResponseWriter, r *http.Request) {
	t,_ := template.ParseFiles("tpl.html", "t2.html")
	d := TplData{
		Info1: "str",
		Data2: "Hello world",
		Dbool: false,
		Dhtml: "<b>strong text</b>",
	}
	_ =  d

	tstr := "abcd<b>asdfase</b>"

	//t.ExecuteTemplate(w, "t2.html", d)
	t.Execute(w, template.HTML(tstr))
}