package storage

import (
	"fmt"
	"net/http"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func MemStore(w http.ResponseWriter, r *http.Request) {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{1, "Hello world", "john doe"}
	post2 := Post{2, "bonjour monde", "pirre"}
	post3 := Post{3, "hola mundo", "pedro"}
	post4 := Post{4, "Greetings", "john doe"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Fprintln(w, PostById[1])
	fmt.Fprintln(w, PostById[2])

	for _, post := range PostsByAuthor["john doe"] {
		fmt.Fprintln(w, post)
	}
	for _, post := range PostsByAuthor["pedro"] {
		fmt.Fprintln(w, post)
	}
}
