// create table test_mysql (id int key, code int);
// insert into test_mysql values (1, 200);

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	CONN_STR = "dbu:dbp@tcp(xmysql:3306)/test?charset=utf8"
)

func main() {
	f2()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func f2() {
	db, err := sql.Open("mysql", CONN_STR)

	rows, err := db.Query("SELECT * FROM test_mysql")
	checkErr(err)

	for rows.Next() {
		id := -1
		code := -1
		err = rows.Scan(&id, &code)
		checkErr(err)
		fmt.Printf("Got row with id: %d, and code: %d \n", id, code)
	}
}

func f1() {
	db, err := sql.Open("mysql", CONN_STR)

	stmt, err := db.Prepare("INSERT INTO test_mysql VALUES (?, ?)")
	checkErr(err)

	res, err := stmt.Exec(2, 204)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("Inserted id:", id)

	checkErr(err)
}
