// // Copyright (C) Subhajit DasGupta 2022

package bookdao

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"

	"github.com/lf8r/dbgen-common/pkg/common"
	"github.com/lf8r/example-data/pkg/data1"
	_ "github.com/lib/pq" // Load the Posgres driver.
)

// Warning - This is generated code. It is overwritten on each build.

// BookClientDBImpl implements BookClient.
type BookClientDBImpl struct {
	Tx *sql.Tx
}

var _ BookClient = (*BookClientDBImpl)(nil)

// Create creates a new data1.Book instance.
func (c *BookClientDBImpl) Create(value *data1.Book) (*data1.Book, error) {
	value.ID = uuid.NewString()

	if value.Name == "" {
		return nil, fmt.Errorf("missing \"Name\" field")
	}

	value.Created = common.Now()
	value.Modified = common.Now()

	content, err := jsoniter.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}

	if _, err := c.Tx.Exec("insert into Book values($1,$2,$3,$4,$5)", value.ID, value.Name, value.Created.Time, value.Modified.Time, string(content)); err != nil {
		return nil, fmt.Errorf("insert: %w", err)
	}

	return value, nil
}

// Delete deletes the given data1.Book instance.
func (c *BookClientDBImpl) Delete(value *data1.Book) error {
	if value.ID == "" {
		return fmt.Errorf("missing id")
	}

	if _, err := c.Tx.Exec("delete from Book where ID=$1", value.ID); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// DeleteByID deletes the data1.Book with the given ID, returning without error
// if the ID doesn't exist.
func (c *BookClientDBImpl) DeleteByID(id string) error {
	if id == "" {
		return fmt.Errorf("missing id")
	}

	if _, err := c.Tx.Exec("delete from Book where ID=$1", id); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// GetByID gets the data1.Book with the given id.
func (c *BookClientDBImpl) GetByID(id string) (*data1.Book, error) {
	if id == "" {
		return nil, fmt.Errorf("missing id")
	}

	rows, err := c.Tx.Query("select * from Book where ID=$1", id)
	if err != nil {
		return nil, fmt.Errorf("select * from Book with id %s: %w", id, err)
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var name string

	var created, modified time.Time

	var content string

	if err := rows.Scan(&id, &name, &created, &modified, &content); err != nil {
		return nil, fmt.Errorf("row scan: %w", err)
	}

	obj := data1.Book{}
	if err := jsoniter.Unmarshal([]byte(content), &obj); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	obj.ID = id
	obj.Name = name
	obj.Created = common.Time{Time: created}
	obj.Modified = common.Time{Time: modified}

	return &obj, nil
}

// List returns a slice of data1.Book instances fulfilling the query (if given).
// The caller must check the returned error (if any) before using the returned
// items.
//
// An example query is m["Name"] = value:
//  m := make(map[string]interface{})
//  m["Name"] = value
//
// A query may be nested by using map values (instead of strings) in the input.
// For example, if the type has a field named "Address" whih itself has a field
// named "Street1", use:
//  m := make(map[string]interface{})
//  m["Address"] = map[string]string{"Street1": v.Address.Street1}
//
// A query may contain more than one entry. For example, the following query
// specifies a "Name" and a street address:
//  m := make(map[string]interface{})
//  m["Name"] = v.Name
//  m["Address"] = map[string]string{"Street1": v.Address.Street1}
func (c *BookClientDBImpl) List(query map[string]interface{}) ([]data1.Book, error) {
	queryStr := "select * from Book"

	// Check if a query is specified by looking for a map[string]interface{}
	// in args.
	var m map[string]interface{}
	if query != nil {
		m = query
	}

	// If m isn't nil and has at least one item, append a suitable WHERE clause.
	if m != nil && len(m) != 0 {
		queryStr += " WHERE data @> '"

		content, err := jsoniter.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("marshal query: %w", err)
		}

		queryStr += string(content) + "'"
	}

	rows, err := c.Tx.Query(queryStr)
	if err != nil {
		return nil, fmt.Errorf("select *: %w", err)
	}

	defer rows.Close()

	ret := make([]data1.Book, 0)

	for {
		if !rows.Next() {
			break
		}

		var id, name string

		var created, modified time.Time

		var content string

		if err := rows.Scan(&id, &name, &created, &modified, &content); err != nil {
			return nil, fmt.Errorf("row scan: %w", err)
		}

		obj := data1.Book{}
		if err := jsoniter.Unmarshal([]byte(content), &obj); err != nil {
			return nil, fmt.Errorf("unmarshal: %w", err)
		}

		obj.ID = id
		obj.Name = name
		obj.Created = common.Time{Time: created}
		obj.Modified = common.Time{Time: modified}

		ret = append(ret, obj)
	}

	return ret, nil
}

// Update updates the given data1.Book instance.
func (c *BookClientDBImpl) Update(value *data1.Book) (*data1.Book, error) {
	saved, err := c.GetByID(value.ID)
	if err != nil {
		return nil, fmt.Errorf("get by id %s: %w", value.ID, err)
	}

	if saved == nil {
		return nil, fmt.Errorf("no data with id %s: %w", value.ID, err)
	}

	content, err := jsoniter.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}

	if _, err = c.Tx.Exec("Update Book set Name=$2, Created=$3, Modified=$4, Data=$5 where ID=$1", value.ID, value.Name, value.Created.Time, time.Now(), string(content)); err != nil {
		return nil, fmt.Errorf("exec update %s: %w", value.ID, err)
	}

	saved, err = c.GetByID(value.ID)
	if err != nil {
		return nil, fmt.Errorf("get by id %s: %w", value.ID, err)
	}

	return saved, nil
}

// PartialUpdate applies changes represented by values to the data1.Book with the
// given id.
func (c *BookClientDBImpl) PartialUpdate(id string, values map[string]interface{}) (*data1.Book, error) {
	value, err := c.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("get by id %s: %w", id, err)
	}

	if value == nil {
		return nil, fmt.Errorf("no data with id %s", id)
	}

	if err = patchVal(value, values); err != nil {
		return nil, fmt.Errorf("patch value: %w", err)
	}

	return c.Update(value)
}

// patchVal patches the given val with the given patch.
func patchVal(val interface{}, patch map[string]interface{}) error {
	content, err := jsoniter.Marshal(val)
	if err != nil {
		return fmt.Errorf("marshap to bytes: %w", err)
	}

	m := make(map[string]interface{})
	err = jsoniter.Unmarshal(content, &m)

	if err != nil {
		return fmt.Errorf("unmarshal to map: %w", err)
	}

	for k, v := range patch {
		if _, ok := m[k]; !ok {
			continue
		}

		m[k] = v
	}

	content, err = jsoniter.Marshal(m)
	if err != nil {
		return fmt.Errorf("marshal patched map to bytes: %w", err)
	}

	err = jsoniter.Unmarshal(content, val)
	if err != nil {
		return fmt.Errorf("unmarshal to val: %w", err)
	}

	return nil
}
