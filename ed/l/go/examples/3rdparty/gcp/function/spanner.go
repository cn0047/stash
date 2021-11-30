package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

const DB = "projects/$p/instances/$i/databases/$db"

var (
	client         *spanner.Client
	initClientOnce sync.Once
)

func init() {
	initSpannerClient()
}

func initSpannerClient() {
	initClientOnce.Do(func() {
		c, err := spanner.NewClient(context.Background(), DB)
		if err != nil {
			log.Printf("spanner.NewClient: %v", err)
			return
		}
		client = c
	})
}

func ReadFromDB(w http.ResponseWriter, r *http.Request) {
	initSpannerClient()

	stmt := spanner.Statement{SQL: `SELECT id, msg FROM test`}
	itr := client.Single().Query(r.Context(), stmt)
	defer itr.Stop()
	for {
		row, err := itr.Next()
		if err == iterator.Done {
			return
		}
		if err != nil {
			http.Error(w, "failed to get next row", http.StatusInternalServerError)
			log.Printf("failed to get next row, err: %v", err)
			return
		}

		var id int64
		var msg string
		err = row.Columns(&id, &msg)
		if err != nil {
			http.Error(w, "failed to parse row", http.StatusInternalServerError)
			log.Printf("failed to parse row, err: %v", err)
			return
		}
		_, _ = fmt.Fprintf(w, "Row id:%d, msg:%s\n", id, msg)
	}
}
