package storage

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlStore(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	CheckErr(err)
	defer db.Close()

	sql := "select `id`,`name` from `base_city`"
	result, err := db.Query(sql)

	CheckErr(err)

	for result.Next() {
		var id int64
		var name string
		//需与查询字段对应，否则会报错
		err = result.Scan(&id, &name)
		CheckErr(err)
		fmt.Fprintln(w, "ID:", id, " Name:", name)
	}
	fmt.Fprintln(w, "this is mysql test")
}
