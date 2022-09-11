package storage

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func FileStore(w http.ResponseWriter, r *http.Request) {
	data := []byte("hello world\n")
	err := ioutil.WriteFile("data1", data, 0777)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Fprintln(w, string(read1))

	file1, _ := os.Create("data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Fprintf(w, "Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Fprintf(w, "Read %d bytes from file\n", bytes)
	fmt.Fprintln(w, string(read2))
}
