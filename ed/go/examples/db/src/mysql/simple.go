package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	CONN_STR = "dbu:dbp@tcp(172.17.0.5:3306)/test?charset=utf8" // xmysql
)

func main() {
	f0()
	f1()
	f2()
	web()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func f4() {
	db, err := sql.Open("mysql", CONN_STR)

	rows, err := db.Query("CALL getCountry(?)", "ua")
	checkErr(err)

	for rows.Next() {
		id := -1
		name := ""
		err = rows.Scan(&id, &name)
		checkErr(err)
		fmt.Printf("Got country with code: %v and name: %v \n", id, name)
	}
}

func f3() {
	db, err := sql.Open("mysql", CONN_STR)
	checkErr(err)

	rows, err := db.Query("SELECT ?", 204)
	checkErr(err)

	for rows.Next() {
		v := -1
		err = rows.Scan(&v)
		checkErr(err)
		fmt.Printf("Got value: %d \n", v)
	}
}

func f2() {
	db, err := sql.Open("mysql", CONN_STR)
	checkErr(err)

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
	checkErr(err)

	stmt, err := db.Prepare("INSERT IGNORE INTO test_mysql VALUES (?, ?)")
	checkErr(err)

	res, err := stmt.Exec(2, 204)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("Inserted id:", id)
}

func f0() {
	db, err := sql.Open("mysql", CONN_STR)
	checkErr(err)

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS test_mysql (id int key, code int)")
	checkErr(err)

	_, er2 := stmt.Exec()
	checkErr(er2)

	stmt2, err := db.Prepare("REPLACE test_mysql VALUES (?, ?)")
	checkErr(err)

	_, err3 := stmt2.Exec(1, 200)
	checkErr(err3)
}

func web() {
	db, err := sql.Open("mysql", CONN_STR)
	checkErr(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM test_mysql")
		checkErr(err)

		w.Write([]byte("<table>"))
		w.Write([]byte("<tr><td>id</td><td>code</td></tr>"))
		for rows.Next() {
			id := -1
			code := -1
			err = rows.Scan(&id, &code)
			checkErr(err)
			w.Write([]byte(fmt.Sprintf("<tr> <td>%d</td> <td>%d</td> </tr>", id, code)))
		}
		w.Write([]byte("</table>"))
	})

	p := ":8080"
	fmt.Printf("Open: http://localhost%s\n", p)
	http.ListenAndServe(p, nil)
}
