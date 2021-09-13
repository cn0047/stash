package bigtable

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigtable"
)

func Run(projectID, bigTableInstanceID, SAFilePath string) {
	ctx := context.Background()
	//c := getAdminClient(ctx, projectID, bigTableInstanceID)
	c := getClient(ctx, projectID, bigTableInstanceID)

	//printTables(ctx, c)
	table := ""
	//createTable(ctx, c, table)
	//printTableInfo(ctx, c, table)
	scanTable(ctx, c, table)
}

func getClient(ctx context.Context, projectID string, bigTableInstanceID string) *bigtable.Client {
	c, err := bigtable.NewClient(ctx, projectID, bigTableInstanceID)
	if err != nil {
		log.Fatalf("failed to create admin client: %v", err)
	}

	return c
}

func getAdminClient(ctx context.Context, projectID string, bigTableInstanceID string) *bigtable.AdminClient {
	c, err := bigtable.NewAdminClient(ctx, projectID, bigTableInstanceID)
	if err != nil {
		log.Fatalf("failed to create admin client: %v", err)
	}

	return c
}

func printTables(ctx context.Context, c *bigtable.AdminClient) {
	tables, err := c.Tables(ctx)
	if err != nil {
		log.Fatalf("failed to fetch table list: %v", err)
	}

	fmt.Printf("Tables in the instance: \n")
	for _, t := range tables {
		fmt.Printf("\t %+v \n", t)
	}

	if len(tables) == 0 {
		return
	}
	printTableInfo(ctx, c, tables[0])
}

func printTableInfo(ctx context.Context, c *bigtable.AdminClient, table string) {
	ti, err := c.TableInfo(ctx, table)
	if err != nil {
		log.Fatalf("failed to read info for table %s: %v", table, err)
	}

	fmt.Printf("Info about table %+v: %+v \n", table, ti)
}

func createTable(ctx context.Context, c *bigtable.AdminClient, table string) {
	if err := c.CreateTable(ctx, table); err != nil {
		log.Fatalf("failed to create table %s: %v", table, err)
	}
}

func scanTable(ctx context.Context, c *bigtable.Client, table string) {
	tbl := c.Open(table)

	err := tbl.ReadRows(ctx, bigtable.PrefixRange(""), func(row bigtable.Row) bool {
		fmt.Printf("Row: %+v \n", row)
		return false
	}, bigtable.RowFilter(bigtable.ColumnFilter("")))
	if err != nil {
		log.Fatalf("failed to read table rows, err: %v", err)
	}
}
