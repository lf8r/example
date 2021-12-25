// Copyright (C) Subhajit DasGupta 2021

package persondao

import (
	"context"
	"fmt"

	"github.com/lf8r/example-data/pkg/data"
	"github.com/lf8r/example/pkg/person"
)

// Warning - This is generated code. It is overwritten on each build.

type personServiceImpl struct {
	ctx context.Context
}

// PersonService is a service level implementation of the PersonClient
// interface, which calls back user defined validation functionality.
func PersonService(ctx context.Context) PersonClient {
	return &personServiceImpl{
		ctx: ctx,
	}
}

var _ PersonClient = (*personServiceImpl)(nil)

// Create creates the given Person and returns the created instance and any
// error.
func (p *personServiceImpl) Create(value *data.Person) (*data.Person, error) {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	p.ctx = ctx

	client, ctx, err := Person(ctx)
	if err != nil {
		return nil, fmt.Errorf("create db client: %w", err)
	}

	if err := person.CreateBusinessLogic(p.ctx, value); err != nil {
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

// Delete deletes the given Person and returns any error.
func (p *personServiceImpl) Delete(value *data.Person) error {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	p.ctx = ctx

	client, ctx, err := Person(ctx)
	if err != nil {
		return fmt.Errorf("create db client: %w", err)
	}

	if err := person.DeleteBusinessLogic(p.ctx, value); err != nil {
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

// DeleteByID deletes the Person with the given id and returns any error.
func (p *personServiceImpl) DeleteByID(id string) error {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	client, ctx, err := Person(ctx)
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

// Update updates the given Person and returns the updated instance and any
// error.
func (p *personServiceImpl) Update(value *data.Person) (*data.Person, error) {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	p.ctx = ctx

	client, ctx, err := Person(ctx)
	if err != nil {
		return nil, fmt.Errorf("create db client: %w", err)
	}

	if err = person.UpdateBusinessLogic(p.ctx, value); err != nil {
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

// List returns a slice of Persons satisfying the given query and any error.
func (p *personServiceImpl) List(query map[string]interface{}) ([]data.Person, error) {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	client, ctx, err := Person(ctx)
	if err != nil {
		return nil, fmt.Errorf("create db client: %w", err)
	}

	p.ctx = ctx

	if err = person.ListBusinessLogic(p.ctx, &data.Person{}, query); err != nil {
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

// GetByID returns the Person with the given id and any error.
func (p *personServiceImpl) GetByID(id string) (*data.Person, error) {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	client, ctx, err := Person(ctx)
	if err != nil {
		return nil, fmt.Errorf("create db client: %w", err)
	}

	p.ctx = ctx

	if err = person.GetByIDBusinessLogic(p.ctx, &data.Person{}, id); err != nil {
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

// PartialUpdate performs a partial update of the data.Person with the given id using
// the given values.
func (p *personServiceImpl) PartialUpdate(id string, values map[string]interface{}) (*data.Person, error) {
	ctx, err := BeginTx(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}

	defer RollbackTx(ctx)

	p.ctx = ctx

	client, ctx, err := Person(ctx)
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
