package res

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	User string
	Threads []string
}

func JsonRes(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "John Doe",
		Threads: []string{"first", "second", "third"},
	}
	jsonStr,_ := json.Marshal(post)
	w.Write(jsonStr)
}
