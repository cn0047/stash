package main

import (
	"fmt"
	"time"
	"database/sql"
	_ "github.com/lib/pq"
)

type Record struct {
	Id int
	Sha1 string
	Path string
}

func (r *Record) Print() {
	fmt.Printf("%d, %s, %s \n", r.Id, r.Sha1, r.Path)
}

func main() {
	now := time.Now()
	nanos := now.UnixNano()

	action()

	now2 := time.Now()
	nanos2 := now2.UnixNano()

	fmt.Printf("Took: %d microseconds", (nanos2 - nanos) / 1000)
}

func action() {
	connStr := "host=xpostgres port=5432 user=dbu password=dbp dbname=test sslmode=disable"
	db, _ := sql.Open("postgres", connStr)
	defer db.Close()

	query := `
		SELECT df.id, df.sha1, dfs.path
		FROM MOCK_device_file df
		JOIN MOCK_device_file_sandbox dfs ON df.id = dfs.device_file_id
		WHERE df.count_devices_for_file_sandbox > 0
		ORDER by id DESC, path ASC
		OFFSET 100 LIMIT 10
	`
	rows, _ := db.Query(query)
	defer rows.Close()

	r := &Record{}
	for rows.Next() {
		rows.Scan(&r.Id, &r.Sha1, &r.Path)
		r.Print()
	}
}
