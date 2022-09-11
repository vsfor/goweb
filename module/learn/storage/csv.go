package storage

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func CsvStore(w http.ResponseWriter, r *http.Request) {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	post1 := Post{1, "Hello world", "john doe"}
	post2 := Post{2, "bonjour monde", "pirre"}
	post3 := Post{3, "hola mundo", "pedro"}
	post4 := Post{4, "Greetings", "john doe"}

	allPosts := []Post{
		post1, post2, post3, post4,
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{
			strconv.Itoa(post.Id),
			post.Content,
			post.Author,
		}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{int(id), item[1], item[2]}
		posts = append(posts, post)
	}

	fmt.Fprintln(w, posts[0].Id)
	fmt.Fprintln(w, posts[0].Content)
	fmt.Fprintln(w, posts[0].Author)

	fmt.Fprintln(w, posts)
}
