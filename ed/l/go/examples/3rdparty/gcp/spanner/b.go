package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/big"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

const DB = "projects/test-project/instances/test-instance/databases/test-db"

func main() {
	var err error

	c, err := getClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	//err = selectTestRow1(ctx, c) // Sanity check.
	//err = insertJSON(ctx, c)
	//err = insertItoFieldsTest(ctx, c)
	err = readFromFieldsTest(ctx, c)
	//err = selectTestRows(ctx, c)
	//err = deleteTestRow1(ctx, c)
	//err = deleteTestRow1v2(ctx, c)
	//err = upsertUsingMutation(ctx, c)
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

type FieldsTestRowWrite struct {
	ID1 string           `spanner:"id1"`
	ID2 string           `spanner:"id2"`
	B1  bool             `spanner:"b1"`
	I1  int64            `spanner:"i1"`
	F1  float64          `spanner:"f1"`
	N1  big.Rat          `spanner:"n1"`
	S1  string           `spanner:"s1"`
	J1  spanner.NullJSON `spanner:"j1"`
	BT1 string           `spanner:"bt1"`
	D1  string           `spanner:"d1"`
	T1  time.Time        `spanner:"t1"`
	T2  time.Time        `spanner:"t2"`
	A1  []int64          `spanner:"a1"`
	A2  []string         `spanner:"a2"`
}

type FieldsTestRowRead struct {
	ID1 string           `spanner:"id1"`
	ID2 string           `spanner:"id2"`
	B1  bool             `spanner:"b1"`
	I1  int64            `spanner:"i1"`
	F1  float64          `spanner:"f1"`
	N1  big.Rat          `spanner:"n1"`
	S1  string           `spanner:"s1"`
	J1  spanner.NullJSON `spanner:"j1"`
	BT1 []byte           `spanner:"bt1"`
	D1  spanner.NullDate `spanner:"d1"`
	T1  time.Time        `spanner:"t1"`
	T2  time.Time        `spanner:"t2"`
	A1  []int64          `spanner:"a1"`
	A2  []string         `spanner:"a2"`
}

func insertItoFieldsTest(ctx context.Context, c *spanner.Client) error {
	location, err := time.LoadLocation("Europe/Kyiv")
	if err != nil {
		return fmt.Errorf("failed to create location, err: %w", err)
	}
	now := time.Now()
	nowInKyiv := now.In(location)

	input := FieldsTestRowWrite{
		ID1: "a",
		ID2: "1",
		D1:  "2022-11-03",
		T1:  nowInKyiv, // In Spanner value will be saved in UTC.
	}
	m1, err := spanner.InsertStruct("fields_test", input)
	if err != nil {
		return fmt.Errorf("failed to create insert mutation, err: %w", err)
	}
	mutations := []*spanner.Mutation{m1}

	cb := func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		err := tx.BufferWrite(mutations)
		if err != nil {
			return fmt.Errorf("failed to perform BufferWrite, err: %w", err)
		}
		return nil
	}
	_, err = c.ReadWriteTransaction(ctx, cb)
	if err != nil {
		return fmt.Errorf("failed to perform ReadWriteTransaction, err: %w", err)
	}

	return nil
}

func readFromFieldsTest(ctx context.Context, c *spanner.Client) error {
	fields := []string{"a1", "a2", "b1", "bt1", "d1", "f1", "i1", "id1", "id2", "j1", "n1", "s1", "t1", "t2"}
	row, err := c.Single().ReadRow(ctx, "fields_test", spanner.Key{"a", "1"}, fields)
	if err != nil {
		return fmt.Errorf("failed to perform ReadRow, err: %w", err)
	}

	r := FieldsTestRowRead{}
	err = row.ToStruct(&r)
	if err != nil {
		return fmt.Errorf("failed convert row to struct, err: %w", err)
	}
	fmt.Printf("result row: %+v \n", r)

	return nil
}

func insertJSON(ctx context.Context, c *spanner.Client) error {
	cb := func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{
			SQL: `INSERT INTO test (id, msg, data) VALUES (@id, @msg, @data)`,
			Params: map[string]interface{}{
				"id":   11,
				"msg":  "New message generated at: " + time.Now().Format(time.Kitchen),
				"data": spanner.NullJSON{Value: map[string]string{"foo": "bar"}, Valid: true},
			},
		}
		_, err := tx.Update(ctx, stmt)
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
	ID   int64            `json:"id" spanner:"id"`
	Msg  string           `json:"msg" spanner:"msg"`
	Data spanner.NullJSON `json:"data" spanner:"data"`
}

func deleteTestRow1v2(ctx context.Context, c *spanner.Client) error {
	cb := func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{
			SQL:    `DELETE FROM test WHERE id = @id`,
			Params: map[string]interface{}{"id": 1},
		}
		_, err := tx.Update(ctx, stmt)
		if err != nil {
			return fmt.Errorf("failed to delete row from spanner, err: %w", err)
		}
		return nil
	}
	_, err := c.ReadWriteTransaction(ctx, cb)
	if err != nil {
		return fmt.Errorf("failed to perform transaction, err: %w", err)
	}

	return nil
}

func deleteTestRow1(ctx context.Context, c *spanner.Client) error {
	m := []*spanner.Mutation{
		spanner.Delete("test", spanner.Key{1}),
	}
	_, err := c.Apply(ctx, m)
	if err != nil {
		return fmt.Errorf("failed to delete row from spanner, err: %w", err)
	}

	return nil
}

func selectTestRows(ctx context.Context, c *spanner.Client) error {
	stmt := spanner.Statement{
		SQL:    `SELECT id, msg FROM test WHERE id IN UNNEST(@ids)`,
		Params: map[string]interface{}{"ids": []int64{1, 11}},
	}
	iter := c.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to get next, err: %w", err)
		}
		r := TestRow{}
		if err := row.Columns(&r.ID, &r.Msg); err != nil {
			return fmt.Errorf("failed to get columns, err: %w", err)
		}
		fmt.Printf("test row with ID *: %+v \n", r)
	}

	return nil
}

func selectTestRow1WithQuery(ctx context.Context, c *spanner.Client) error {
	stmt := spanner.Statement{
		SQL:    `SELECT id, msg FROM test WHERE id = @id`,
		Params: map[string]interface{}{"id": 1},
	}
	iter := c.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to get next, err: %w", err)
		}
		r := TestRow{}
		if err := row.Columns(&r.ID, &r.Msg); err != nil {
			return fmt.Errorf("failed to get columns, err: %w", err)
		}
		fmt.Printf("test row with ID *: %+v \n", r)
	}

	return nil
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

	cb := func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		var n = Name{"Timothy", "Campbell"}
		stmt := spanner.Statement{
			SQL: `
				Update Users Set LastName = 'Grant'
				WHERE STRUCT<FirstName String, LastName String>(Firstname, LastName) = @name
			`,
			Params: map[string]interface{}{"name": n},
		}
		rowCount, err := tx.Update(ctx, stmt)
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(w, "%d record(s) updated.\n", rowCount)
		return nil
	}
	_, err = client.ReadWriteTransaction(ctx, cb)

	return err
}

func upsertUsingMutation(ctx context.Context, c *spanner.Client) error {
	table := "test"
	columns := []string{"id", "msg"}
	m := []*spanner.Mutation{
		spanner.InsertOrUpdate(table, columns, []interface{}{5, "upsert"}),
	}
	_, err := c.Apply(ctx, m)
	if err != nil {
		return fmt.Errorf("failed to upsert row into spanner, err: %w", err)
	}

	return nil
}
