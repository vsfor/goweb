package storage

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"net/http"
)

func StoreGob(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0777)
	if err != nil {
		panic(err)
	}
}

func LoadGob(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func GobStore(w http.ResponseWriter, r *http.Request) {
	post := Post{1, "hello world~", "johnDoe"}
	StoreGob(post, "post1gob")
	var postRead Post
	LoadGob(&postRead, "post1gob")
	fmt.Fprintln(w, postRead)
}
