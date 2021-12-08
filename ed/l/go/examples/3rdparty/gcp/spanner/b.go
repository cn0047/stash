package main

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/spanner"
)

const DB = "projects/$p/instances/$i/databases/$db"

func main() {

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
