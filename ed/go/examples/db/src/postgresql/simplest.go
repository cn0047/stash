package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"log"
)

type Row struct {
	n string
	c int
}

func main() {
	connStr := "host=xpostgres port=5432 user=dbu password=dbp dbname=test sslmode=disable"
	// connStr := "host=localhost port=5432 user=cws password=yalloThere dbname=cws sslmode=disable"
	db, err1 := sql.Open("postgres", connStr)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer db.Close()

	result := &Row{}
	row := db.QueryRow("SELECT NOW() AS n, $1 AS c", "200")
	err := row.Scan(&result.n, &result.c)
	if err == sql.ErrNoRows {
		log.Fatal("ErrNoRows", err)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", *result)
}

func in() {
	// db.Query(&result, `SELECT id FROM bar WHERE code_id IN (?)`, pg.In(params.Codes))
	// db.Query(&result, `SELECT id FROM bar WHERE code_id = ANY (?)`, pg.Array(params.Codes))
}
