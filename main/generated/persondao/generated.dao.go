// Copyright (C) // Copyright (C) Subhajit DasGupta 2022

package persondao

import (
	"github.com/lf8r/example-data/pkg/data"
)

// Warning - This is generated code. It is overwritten on each build.

// PersonClient is a generic interface for manipulating data.Person
type PersonClient interface {
	// Create creates a new data.Person.
	Create(value *data.Person) (*data.Person, error)

	// Update updates the given data.Person.
	Update(value *data.Person) (*data.Person, error)

	// Delete deletes the given data.Person.
	Delete(value *data.Person) error

	// DeleteByID deletes the data.Person with the given ID.
	DeleteByID(id string) error

	// GetByID returns the data.Person with the given ID.
	GetByID(id string) (*data.Person, error)

	// List returns all data.Person items satisfying the given query.
	List(query map[string]interface{}) ([]data.Person, error)

	// PartialUpdate performs the updates given in values to the data.Person
	// with the given id, and returns the updated data.Person.
	PartialUpdate(id string, values map[string]interface{}) (*data.Person, error)
}
