package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"vsfor/goweb/module/learn/res"
	"vsfor/goweb/module/learn/storage"
	"vsfor/goweb/module/tm/tc"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello world, %s ~", request.URL.Path[1:])
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("handler function called - " + name)
		h(w, r)
	}
}

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("User-Agent")
	q := r.URL.RawQuery
	fmt.Fprintln(w, h, "<br>", q)
}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.PostForm)
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head>
<title>Go Web Programming</title>
</head> 
<body>
<h1>Hello World</h1>
</body>
</html>`
	w.Write([]byte(str))
}

func main() {
	fmt.Println("start server...")
	handler := MyHandler{}
	hello := HelloHandler{}
	world := WorldHandler{}
	writeHeader := tc.TaHandler{}

	fmt.Println("init...")
	server := http.Server{
		Addr: "0.0.0.0:8181",
	}

	fmt.Println("add routes...")
	http.HandleFunc("/db-mysql", storage.MysqlStore)
	http.HandleFunc("/db-sqlite", storage.SqliteStore)
	http.HandleFunc("/db-gob", storage.GobStore)
	http.HandleFunc("/db-csv", storage.CsvStore)
	http.HandleFunc("/db-file", storage.FileStore)
	http.HandleFunc("/db-memo", storage.MemStore)
	http.HandleFunc("/layout", res.ShowLayout)
	http.HandleFunc("/tpl", res.ShowTpl)
	http.HandleFunc("/redirect", res.Location302)
	http.HandleFunc("/json", res.JsonRes)
	http.HandleFunc("/cookie", res.CookieSet)

	http.Handle("/writeHeaders", log(tc.WriteHeader))
	http.Handle("/writeHeader", log(writeHeader.ServeHTTP))
	http.Handle("/write", log(writeExample))
	http.Handle("/headers", log(headers))
	http.Handle("/process", log(process))
	http.Handle("/hello", log(hello.ServeHTTP))
	http.Handle("/world", log(world.ServeHTTP))
	http.Handle("/", log(handler.ServeHTTP))

	fmt.Println("start listen...")
	server.ListenAndServe()
}
