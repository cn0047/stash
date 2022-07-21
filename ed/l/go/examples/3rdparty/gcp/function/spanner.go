// This is 1 gen func.

package p

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

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

func main() {
	err := InsertIntoTestTable(context.Background())
	if err != nil {
		log.Printf("failed create new spanner client: %v", err)
		return
	}

	CLIReadFromDB()
}

func initSpannerClient() {
	initClientOnce.Do(func() {
		c, err := spanner.NewClient(context.Background(), DB)
		if err != nil {
			log.Printf("failed create new spanner client: %v", err)
			return
		}
		client = c
	})
}

func CLIReadFromDB() {
	data, err := SelectFromTestTable(context.Background())
	if err != nil {
		log.Printf("failed to ReadFromTestTable, err: %v", err)
		return
	}

	for _, row := range data {
		_, _ = fmt.Printf("%s\n", row)
	}
}

func HTTPReadFromDB(w http.ResponseWriter, r *http.Request) {
	if err := InsertIntoTestTable; err != nil {
		http.Error(w, "failed to InsertIntoTestTable", http.StatusInternalServerError)
		log.Printf("failed to InsertIntoTestTable, err: %v", err)
		return
	}

	data, err := SelectFromTestTable(r.Context())
	if err != nil {
		http.Error(w, "failed to ReadFromTestTable", http.StatusInternalServerError)
		log.Printf("failed to ReadFromTestTable, err: %v", err)
		return
	}

	for _, row := range data {
		_, _ = fmt.Fprintf(w, "Rows:\n%s\n", row)
	}
}

func SelectFromTestTable(ctx context.Context) ([]string, error) {
	data := make([]string, 0)

	stmt := spanner.Statement{SQL: `SELECT id, msg FROM test`}
	itr := client.Single().Query(ctx, stmt)
	defer itr.Stop()
	for {
		row, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get next row, err: %w", err)
		}

		var id int64
		var msg string
		err = row.Columns(&id, &msg)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row, err: %w", err)
		}

		data = append(data, fmt.Sprintf("Row id:%d, msg:%s", id, msg))
	}

	return data, nil
}

func InsertIntoTestTable(ctx context.Context) error {
	err1 := InsertIntoTestTable1(ctx)
	err2 := InsertIntoTestTable2(ctx)
	err3 := InsertIntoTestTable3(ctx)

	if err1 == nil && err2 == nil && err3 == nil {
		return nil
	}

	return fmt.Errorf("err1: %v, err2: %v, err3: %v", err1, err2, err3)
}

// InsertIntoTestTable1 - Approach #1.
func InsertIntoTestTable1(ctx context.Context) error {
	insertMutation := spanner.Insert(
		"test",
		[]string{"id", "msg"},
		[]interface{}{
			time.Now().UTC().UnixMilli(),
			"New message #1 generated at: " + time.Now().Format(time.Kitchen),
		},
	)
	_, err := client.Apply(ctx, []*spanner.Mutation{insertMutation})
	if err != nil {
		return fmt.Errorf("failed to perform apply, err: %w", err)
	}

	return nil
}

// InsertIntoTestTable2 - Approach #2.
func InsertIntoTestTable2(ctx context.Context) error {
	cb := func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{
			SQL: `INSERT INTO test (id, msg) VALUES (@id, @msg)`,
			Params: map[string]interface{}{
				"id":  time.Now().UTC().UnixMilli(),
				"msg": "New message #2 generated at: " + time.Now().Format(time.Kitchen),
			},
		}
		_, err := tx.Update(ctx, stmt)
		if err != nil {
			return fmt.Errorf("failed to perform insert, err: %w", err)
		}
		return nil
	}
	_, err := client.ReadWriteTransaction(ctx, cb)
	if err != nil {
		return fmt.Errorf("failed to perform ReadWriteTransaction, err: %w", err)
	}

	return nil
}

// InsertIntoTestTable3 - Approach #3.
func InsertIntoTestTable3(ctx context.Context) error {
	cb := func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{
			SQL: `
				INSERT INTO test (id, msg)
				VALUES ((SELECT MAX(id) FROM test)+1, @msg)
			`,
			Params: map[string]interface{}{
				"id":  time.Now().UTC().UnixMilli(),
				"msg": "New message #3 with sub query generated at: " + time.Now().Format(time.Kitchen),
			},
		}
		_, err := tx.Update(ctx, stmt)
		if err != nil {
			return fmt.Errorf("failed to perform insert, err: %w", err)
		}
		return nil
	}
	_, err := client.ReadWriteTransaction(ctx, cb)
	if err != nil {
		return fmt.Errorf("failed to perform ReadWriteTransaction, err: %w", err)
	}

	return nil
}
