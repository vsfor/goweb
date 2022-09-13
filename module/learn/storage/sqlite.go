package storage

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func SqliteStore(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "post.db")
	CheckErr(err)
	defer db.Close()

	createSql := `
create table if not exists posts (
	id integer primary key autoincrement,
	content varchar(64) null,
	author varchar(64) null
);
`

	db.Exec(createSql)

	addSql := `
insert into posts(content, author) values (?,?)
`
	stmt, err := db.Prepare(addSql)
	CheckErr(err)

	res, err := stmt.Exec("hello sqlite", "john doe")
	CheckErr(err)

	id, err := res.LastInsertId()
	CheckErr(err)

	fmt.Fprintln(w, id)
	fmt.Fprintln(w, "this is sqlite test")
}
