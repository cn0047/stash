package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"cloud.google.com/go/spanner"
)

const DB = "projects/$p/instances/$i/databases/$db"

func main() {
	var err error

	c, err := getClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	//err = insertJSON(ctx, c)
	err = selectTestRow1(ctx, c)
	if err != nil {
		log.Fatal(err)
	}
}

func getClient() (*spanner.Client, error) {
	c, err := spanner.NewClient(context.Background(), DB)
	if err != nil {
		return nil, fmt.Errorf("failed create new spanner client: %w", err)
	}

	return c, nil
}

func insertJSON(ctx context.Context, c *spanner.Client) error {
	cb := func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{
			SQL: `INSERT INTO test (id, msg, data) VALUES (@id, @msg, @data)`,
			Params: map[string]interface{}{
				"id":   1,
				"msg":  "New message generated at: " + time.Now().Format(time.Kitchen),
				"data": spanner.NullJSON{Value: map[string]string{"foo": "bar"}, Valid: true},
			},
		}
		_, err := txn.Update(ctx, stmt)
		if err != nil {
			return fmt.Errorf("failed to perform insert, err: %w", err)
		}
		return nil
	}
	_, err := c.ReadWriteTransaction(ctx, cb)
	if err != nil {
		return fmt.Errorf("failed to perform ReadWriteTransaction, err: %w", err)
	}

	return nil
}

type TestRow struct {
	ID   int64            `json:"id"`
	Msg  string           `json:"msg"`
	Data spanner.NullJSON `json:"data"`
}

func selectTestRow1(ctx context.Context, c *spanner.Client) error {
	row, err := c.Single().ReadRow(ctx, "test", spanner.Key{1}, []string{"id", "msg", "data"})
	if err != nil {
		return fmt.Errorf("failed to perform ReadRow, err: %w", err)
	}

	r := TestRow{}
	err = row.ToStruct(&r)
	if err != nil {
		return fmt.Errorf("failed convert row to struct, err: %w", err)
	}
	fmt.Printf("test row with ID 1: %+v \n", r)

	return nil
}

type Name struct {
	FirstName string
	LastName  string
}

func updateUsingDMLStruct(w io.Writer, db string) error {
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		return err
	}
	defer client.Close()

	cb := func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		var n = Name{"Timothy", "Campbell"}
		stmt := spanner.Statement{
			SQL: `
				Update Users Set LastName = 'Grant'
				WHERE STRUCT<FirstName String, LastName String>(Firstname, LastName) = @name
			`,
			Params: map[string]interface{}{"name": n},
		}
		rowCount, err := txn.Update(ctx, stmt)
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(w, "%d record(s) updated.\n", rowCount)
		return nil
	}
	_, err = client.ReadWriteTransaction(ctx, cb)

	return err
}
