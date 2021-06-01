// create table testj (id character varying(128) not null primary key, data json not null);
// insert into testj values ('1', '{ "name": "jt",  "brand": "White" }')

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Row struct {
	n string
	c int
}

type RowJ struct {
	ID   string
	Data json.RawMessage
}

type RowJData struct {
	Name  string
	Brand string
}

func main() {
	connStr := "host=xpostgres port=5432 user=dbu password=dbp dbname=test sslmode=disable"
	// connStr := "host=localhost port=5432 user=cws password=yalloThere dbname=cws sslmode=disable"
	db, err1 := sql.Open("postgres", connStr)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer db.Close()

	f1(db)
	//f2(db)
}

func f1(db *sql.DB) {
	result := &Row{}
	row := db.QueryRow("SELECT NOW() AS n, $1 AS c", "200")
	err := row.Scan(&result.n, &result.c)
	if err == sql.ErrNoRows {
		log.Fatal("ErrNoRows", err)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", *result)
}

func f2(db *sql.DB) {
	result := &RowJ{}
	row := db.QueryRow("SELECT id, data FROM testj LIMIT 1")
	err := row.Scan(&result.ID, &result.Data)
	if err == sql.ErrNoRows {
		log.Fatal("ErrNoRows", err)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", *result)

	data := RowJData{}
	json.Unmarshal(result.Data, &data)
	fmt.Printf("%#v", data)
}

func in() {
	// db.Query(&result, `SELECT id FROM bar WHERE code_id IN (?)`, pg.In(params.Codes))
	// db.Query(&result, `SELECT id FROM bar WHERE code_id = ANY (?)`, pg.Array(params.Codes))
}
