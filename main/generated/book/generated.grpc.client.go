// Copyright (C) Subhajit DasGupta 2021

package book

import (
	"github.com/lf8r/example-data/pkg/data1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client interface {
	// Create creates a new data1.Book.
	Create(ctx context.Context, value *data1.Book) (*data1.Book, error)

	// Update updates the given data1.Book.
	Update(ctx context.Context, value *data1.Book) (*data1.Book, error)

	// Delete deletes the given data1.Book.
	Delete(ctx context.Context, value *data1.Book) error

	// DeleteByID deletes the data1.Book with the given ID.
	DeleteByID(ctx context.Context, id string) error

	// GetByID returns the data1.Book with the given ID.
	GetByID(ctx context.Context, id string) (*data1.Book, error)

	// List returns all data1.Book items fulfilling the given query.
	List(ctx context.Context, query map[string]interface{}) ([]data1.Book, error)

	// PartialUpdate performs the updates given in values to the data1.Book
	// with the given id, and returns the updated data1.Book.
	PartialUpdate(ctx context.Context, id string, values map[string]interface{}) (*data1.Book, error)
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

func (c *client) Create(ctx context.Context, val *data1.Book) (*data1.Book, error) {
	req, err := NewCreateBookRequest(val)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.CreateBook(ctx, req)
	if err != nil {
		return nil, err
	}

	return ParseCreateBookResponse(resp)
}

func (c *client) Update(ctx context.Context, value *data1.Book) (*data1.Book, error) {
	req, err := NewUpdateBookRequest(value)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.UpdateBook(ctx, req)
	if err != nil {
		return nil, err
	}

	return ParseUpdateBookResponse(resp)
}

func (c *client) Delete(ctx context.Context, value *data1.Book) error {
	req, err := NewDeleteBookRequest(value)
	if err != nil {
		return err
	}

	_, err = c.client.DeleteBook(ctx, req)

	return err
}

func (c *client) DeleteByID(ctx context.Context, id string) error {
	req := DeleteByIDBookRequest{
		Body: id,
	}

	_, err := c.client.DeleteByIDBook(ctx, &req)

	return err
}

func (c *client) GetByID(ctx context.Context, id string) (*data1.Book, error) {
	req := GetByIDBookRequest{
		Body: id,
	}

	resp, err := c.client.GetByIDBook(ctx, &req)
	if err != nil {
		return nil, err
	}

	return ParseGetByIDBookResponse(resp)
}

func (c *client) List(ctx context.Context, query map[string]interface{}) ([]data1.Book, error) {
	req := ListBookRequest{
		Query: query,
	}

	resp, err := c.client.ListBook(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (c *client) PartialUpdate(ctx context.Context, id string, update map[string]interface{}) (*data1.Book, error) {
	req := PartialUpdateBookRequest{
		Id:   id,
		Body: update,
	}

	resp, err := c.client.PartialUpdateBook(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
