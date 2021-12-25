// Copyright (C) Subhajit DasGupta 2021

package person

import (
	"github.com/lf8r/example-data/pkg/data"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client interface {
	// Create creates a new data.Person.
	Create(ctx context.Context, value *data.Person) (*data.Person, error)

	// Update updates the given data.Person.
	Update(ctx context.Context, value *data.Person) (*data.Person, error)

	// Delete deletes the given data.Person.
	Delete(ctx context.Context, value *data.Person) error

	// DeleteByID deletes the data.Person with the given ID.
	DeleteByID(ctx context.Context, id string) error

	// GetByID returns the data.Person with the given ID.
	GetByID(ctx context.Context, id string) (*data.Person, error)

	// List returns all data.Person items fulfilling the given query.
	List(ctx context.Context, query map[string]interface{}) ([]data.Person, error)

	// PartialUpdate performs the updates given in values to the data.Person
	// with the given id, and returns the updated data.Person.
	PartialUpdate(ctx context.Context, id string, values map[string]interface{}) (*data.Person, error)
}

func NewClient(conn *grpc.ClientConn) Client {
	cl := client{
		conn: conn,
	}

	cl.client = NewServiceClient(conn)

	return &cl
}

type client struct {
	conn   *grpc.ClientConn
	client ServiceClient
}

var _ Client = (*client)(nil)

func (c *client) Create(ctx context.Context, val *data.Person) (*data.Person, error) {
	req, err := NewCreatePersonRequest(val)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.CreatePerson(ctx, req)
	if err != nil {
		return nil, err
	}

	return ParseCreatePersonResponse(resp)
}

func (c *client) Update(ctx context.Context, value *data.Person) (*data.Person, error) {
	req, err := NewUpdatePersonRequest(value)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.UpdatePerson(ctx, req)
	if err != nil {
		return nil, err
	}

	return ParseUpdatePersonResponse(resp)
}

func (c *client) Delete(ctx context.Context, value *data.Person) error {
	req, err := NewDeletePersonRequest(value)
	if err != nil {
		return err
	}

	_, err = c.client.DeletePerson(ctx, req)

	return err
}

func (c *client) DeleteByID(ctx context.Context, id string) error {
	req := DeleteByIDPersonRequest{
		Body: id,
	}

	_, err := c.client.DeleteByIDPerson(ctx, &req)

	return err
}

func (c *client) GetByID(ctx context.Context, id string) (*data.Person, error) {
	req := GetByIDPersonRequest{
		Body: id,
	}

	resp, err := c.client.GetByIDPerson(ctx, &req)
	if err != nil {
		return nil, err
	}

	return ParseGetByIDPersonResponse(resp)
}

func (c *client) List(ctx context.Context, query map[string]interface{}) ([]data.Person, error) {
	req := ListPersonRequest{
		Query: query,
	}

	resp, err := c.client.ListPerson(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (c *client) PartialUpdate(ctx context.Context, id string, update map[string]interface{}) (*data.Person, error) {
	req := PartialUpdatePersonRequest{
		Id:   id,
		Body: update,
	}

	resp, err := c.client.PartialUpdatePerson(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
