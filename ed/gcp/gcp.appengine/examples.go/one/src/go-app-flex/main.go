package main

import (
  "database/sql"
  "fmt"
  "net/http"

  _ "github.com/lib/pq"
  "google.golang.org/appengine"
)

type Row struct {
  n string
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("ok"))
  })
  http.HandleFunc("/x", x)
  appengine.Main()
}

func x(w http.ResponseWriter, r *http.Request) {
  datastoreName := "sslmode=disable dbname=d user=u password=p host=/cloudsql/kint-dev:us-central1:productsdb port=5432"

  db, err := sql.Open("postgres", datastoreName)
  if err != nil {
    w.Write([]byte("[ERR]" + err.Error()))
    return
  }

  row := db.QueryRow("SELECT NOW() AS n")
  result := &Row{}
  er := row.Scan(&result.n)
  if er != nil {
    w.Write([]byte("[ERR]" + er.Error()))
    return
  }

  w.Write([]byte(fmt.Sprintf("[RES] %#v", *result)))
}
