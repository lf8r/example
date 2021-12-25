// Copyright (C) // Copyright (C) Subhajit DasGupta 2021

package bookdao

import (
	"github.com/lf8r/example-data/pkg/data1"
)

// Warning - This is generated code. It is overwritten on each build.

// BookClient is a generic interface for manipulating data1.Book
type BookClient interface {
	// Create creates a new data1.Book.
	Create(value *data1.Book) (*data1.Book, error)

	// Update updates the given data1.Book.
	Update(value *data1.Book) (*data1.Book, error)

	// Delete deletes the given data1.Book.
	Delete(value *data1.Book) error

	// DeleteByID deletes the data1.Book with the given ID.
	DeleteByID(id string) error

	// GetByID returns the data1.Book with the given ID.
	GetByID(id string) (*data1.Book, error)

	// List returns all data1.Book items satisfying the given query.
	List(query map[string]interface{}) ([]data1.Book, error)

	// PartialUpdate performs the updates given in values to the data1.Book
	// with the given id, and returns the updated data1.Book.
	PartialUpdate(id string, values map[string]interface{}) (*data1.Book, error)
}
