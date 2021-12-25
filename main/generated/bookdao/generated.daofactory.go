// Copyright (C) Subhajit DasGupta 2021

package bookdao

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
)

var Db *sql.DB

// ClientFactory is a factory for data access clients.
type ClientFactory struct {
	sync.Mutex
}

var clientFactory *ClientFactory = &ClientFactory{}

const (
	// BookTypeID is the typeID for data1.Book
	BookTypeID = "data1.Book"
)

// BeginTx starts a transaction if the given context does not already contain
// one, sets it into the given context and returns the context. If the given
// context already contains a transaction, it returns the context with a nil
// error. If there is an error starting the transaction, it returns the given
// context and the error.
func BeginTx(ctx context.Context) (context.Context, error) {
	tx := ctx.Value("tx")

	if tx != nil {
		return ctx, nil
	}

	txv, err := Db.Begin()
	if err != nil {
		return ctx, fmt.Errorf("begin tx: %w", err)
	}

	tx = txv
	ctx = context.WithValue(ctx, "tx", tx)

	return ctx, nil
}

// CommitTx commits the transaction in the given context.
func CommitTx(ctx context.Context) error {
	// Commit the transaction.
	txv := ctx.Value("tx")
	tx, ok := txv.(*sql.Tx)

	if !ok {
		return fmt.Errorf("ctx.tx is not *sql.Tx")
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

// RollbackTx rolls back the transaction in the given context.
func RollbackTx(ctx context.Context) error {
	// Commit the transaction.
	txv := ctx.Value("tx")
	tx, ok := txv.(*sql.Tx)

	if !ok {
		return fmt.Errorf("ctx.tx is not *sql.Tx")
	}

	if err := tx.Rollback(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

// Book returns a BookClient, a context containing transaction information
// and any error.
func Book(ctx context.Context) (BookClient, context.Context, error) {
	return clientFactory.newBookClient(ctx)
}

// newBookClient returns a BookClient.
func (cf *ClientFactory) newBookClient(ctx context.Context) (BookClient, context.Context, error) {
	cf.Lock()
	defer cf.Unlock()

	if Db == nil {
		return nil, nil, fmt.Errorf("no sql db")
	}

	ctx, err := BeginTx(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("begin tx: %w", err)
	}

	tx := ctx.Value("tx")

	return &BookClientDBImpl{
		Tx: tx.(*sql.Tx),
	}, ctx, nil
}

// Client returns a DB access client corresponding to the given typeID, using
// the given ctx.
func Client(ctx context.Context, typeID string) (interface{}, context.Context, error) {
	switch typeID {
	case BookTypeID:
		val, ctxn, err := Book(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf("client with type-id %s: %w", typeID, err)
		}

		return val, ctxn, nil

	default:
		return nil, nil, fmt.Errorf("client with unknown type-id %s", typeID)
	}
}
