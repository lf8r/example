// Copyright (C) Subhajit DasGupta 2021

package bookdao

import (
	"context"
	"fmt"

	"github.com/lf8r/example-data/pkg/data1"
	"github.com/lf8r/example/pkg/book"
)

// Warning - This is generated code. It is overwritten on each build.

type bookServiceImpl struct {
	ctx context.Context
}

// BookService is a service level implementation of the BookClient
// interface, which calls back user defined validation functionality.
func BookService(ctx context.Context) BookClient {
	return &bookServiceImpl{
		ctx: ctx,
	}
}

var _ BookClient = (*bookServiceImpl)(nil)

// Create creates the given Book and returns the created instance and any
// error.
func (p *bookServiceImpl) Create(value *data1.Book) (*data1.Book, error) {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	p.ctx = ctx

	client, ctx, err := Book(ctx)
	if err != nil {
		return nil, fmt.Errorf("create db client: %w", err)
	}

	if err := book.CreateBusinessLogic(p.ctx, value); err != nil {
		return nil, fmt.Errorf("create business logic: %w", err)
	}

	val, err := client.Create(value)
	if err != nil {
		return nil, fmt.Errorf("db create: %w", err)
	}

	if err := CommitTx(ctx); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
	}

	return val, nil
}

// Delete deletes the given Book and returns any error.
func (p *bookServiceImpl) Delete(value *data1.Book) error {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	p.ctx = ctx

	client, ctx, err := Book(ctx)
	if err != nil {
		return fmt.Errorf("create db client: %w", err)
	}

	if err := book.DeleteBusinessLogic(p.ctx, value); err != nil {
		return fmt.Errorf("delete business logic: %w", err)
	}

	if err := client.Delete(value); err != nil {
		return fmt.Errorf("db delete: %w", err)
	}

	if err := CommitTx(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

// DeleteByID deletes the Book with the given id and returns any error.
func (p *bookServiceImpl) DeleteByID(id string) error {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	client, ctx, err := Book(ctx)
	if err != nil {
		return fmt.Errorf("create db client: %w", err)
	}

	value, err := client.GetByID(id)
	if err != nil {
		return fmt.Errorf("get by id failed %s: %w", id, err)
	}

	if value == nil {
		return nil
	}

	p.ctx = ctx

	if err := client.Delete(value); err != nil {
		return fmt.Errorf("db delete: %w", err)
	}

	if err := CommitTx(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

// Update updates the given Book and returns the updated instance and any
// error.
func (p *bookServiceImpl) Update(value *data1.Book) (*data1.Book, error) {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	p.ctx = ctx

	client, ctx, err := Book(ctx)
	if err != nil {
		return nil, fmt.Errorf("create db client: %w", err)
	}

	if err = book.UpdateBusinessLogic(p.ctx, value); err != nil {
		return nil, fmt.Errorf("update business logic: %w", err)
	}

	val, err := client.Update(value)
	if err != nil {
		return nil, fmt.Errorf("db update: %w", err)
	}

	if err := CommitTx(ctx); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
	}

	return val, nil
}

// List returns a slice of Books satisfying the given query and any error.
func (p *bookServiceImpl) List(query map[string]interface{}) ([]data1.Book, error) {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	client, ctx, err := Book(ctx)
	if err != nil {
		return nil, fmt.Errorf("create db client: %w", err)
	}

	p.ctx = ctx

	if err = book.ListBusinessLogic(p.ctx, &data1.Book{}, query); err != nil {
		return nil, fmt.Errorf("list business logic: %w", err)
	}

	val, err := client.List(query)
	if err != nil {
		return nil, fmt.Errorf("db list: %w", err)
	}

	if err := CommitTx(ctx); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
	}

	return val, nil
}

// GetByID returns the Book with the given id and any error.
func (p *bookServiceImpl) GetByID(id string) (*data1.Book, error) {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	client, ctx, err := Book(ctx)
	if err != nil {
		return nil, fmt.Errorf("create db client: %w", err)
	}

	p.ctx = ctx

	if err = book.GetByIDBusinessLogic(p.ctx, &data1.Book{}, id); err != nil {
		return nil, fmt.Errorf("get-by-id business logic: %w", err)
	}

	val, err := client.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("db get by id %s: %w", id, err)
	}

	if err := CommitTx(ctx); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
	}

	return val, nil
}

// PartialUpdate performs a partial update of the data1.Book with the given id using
// the given values.
func (p *bookServiceImpl) PartialUpdate(id string, values map[string]interface{}) (*data1.Book, error) {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	p.ctx = ctx

	client, ctx, err := Book(ctx)
	if err != nil {
		return nil, fmt.Errorf("create db client: %w", err)
	}

	val, err := client.PartialUpdate(id, values)
	if err != nil {
		return nil, fmt.Errorf("db partial update: %w", err)
	}

	if err := CommitTx(ctx); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
	}

	return val, nil
}
