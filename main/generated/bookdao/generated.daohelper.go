// Copyright (C) Subhajit DasGupta 2021

package bookdao

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/lf8r/dbgen/pkg/ddlgen"
	"github.com/lf8r/dbgen/pkg/reflects"
	"github.com/lf8r/example-data/pkg/data1"
)

// SetupBookData sets up developtime time and tables and example data for the data1.Book type.
func SetupBookData(db *sql.DB) error {
	// Define the Book type in the Dev DB.
	createStatements, dropStatements, err := ddlgen.GenerateJSONDDL(reflect.TypeOf(data1.Book{}))
	if err != nil {
		return fmt.Errorf("generate DDL to create Book: %w", err)
	}

	conn, err := db.Conn(context.Background())
	if err != nil {
		return fmt.Errorf("get connection: %w", err)
	}

	for _, drop := range dropStatements {
		conn.ExecContext(context.Background(), drop)
	}

	for _, create := range createStatements {
		if _, err = conn.ExecContext(context.Background(), create); err != nil {
			return fmt.Errorf("run DDL %s: %w", create, err)
		}
	}

	clientFactory = &ClientFactory{}

	// Use the client factory to create clients for all the data types being
	// served and enter some dev data.
	ctx := context.Background()
	cl, ctx, err := Book(ctx)

	if err != nil {
		return fmt.Errorf("create BookClient: %w", err)
	}

	const count = 10
	if err := InsertBookData(cl, count); err != nil {
		return fmt.Errorf("insert Book: %w", err)
	}

	// Commit the transaction.
	if err := CommitTx(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

// InsertBookData inserts the given number of Book objects using the
// given client. This is useful to setup integration tests.
func InsertBookData(client BookClient, count int) error {
	for i := 0; i < count; i++ {
		p := data1.Book{}
		reflects.Set(&p)

		_, err := client.Create(&p)
		if err != nil {
			return fmt.Errorf("insert Book: %w", err)
		}
	}

	return nil
}
