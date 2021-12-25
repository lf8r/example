// Copyright (C) Subhajit DasGupta 2021

package person

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/lf8r/example-data/pkg/data"
	"github.com/lf8r/example/main/generated/persondao"
)

// Warning - This is generated code. It is overwritten on each build.

type personRestClient struct {
	// baseAddr is of the form http://localhost:8080/rest/persons
	baseAddr string
	client   *http.Client
}

var _ persondao.PersonClient = (*personRestClient)(nil)

// NewPersonRestClient creates a new PersonClient REST client.
func NewPersonRestClient(baseAddr string) persondao.PersonClient {
	return &personRestClient{
		baseAddr: baseAddr,
		client:   &http.Client{},
	}
}

// Create creates the given Person person and
// returns the created person and an error if there's any.
func (c *personRestClient) Create(p *data.Person) (*data.Person, error) {
	content, err := jsoniter.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("marshal for %s: %w", c.baseAddr, err)
	}

	req, err := http.NewRequest("POST", c.baseAddr, bytes.NewBuffer(content))
	if err != nil {
		return nil, fmt.Errorf("new post request for %s: %w", c.baseAddr, err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http/post response from %s: %w", c.baseAddr, err)
	}

	r := resp.Body
	defer resp.Body.Close()

	content, err = ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read http/get response body from %s: %w", c.baseAddr, err)
	}

	ret := data.Person{}
	if err := jsoniter.Unmarshal(content, &ret); err != nil {
		return nil, fmt.Errorf("unmarshal http/get response body from %s: %w", c.baseAddr, err)
	}

	return &ret, nil
}

// Delete deletes the given person and returns any error.
func (c *personRestClient) Delete(p *data.Person) error {
	req, err := http.NewRequest("DELETE", c.baseAddr, nil)
	if err != nil {
		return fmt.Errorf("new delete request for %s: %w", c.baseAddr, err)
	}

	if _, err = c.client.Do(req); err != nil {
		return fmt.Errorf("http/delete response from %s: %w", c.baseAddr, err)
	}

	return nil
}

// DeleteByID deletes the person with the given and returns any error.
func (c *personRestClient) DeleteByID(id string) error {
	val, err := c.GetByID(id)
	if err != nil {
		return fmt.Errorf("delete by id for %s: %w", c.baseAddr, err)
	}

	return c.Delete(val)
}

// GetByID gets the person with the given id and returns any error.
func (c *personRestClient) GetByID(id string) (*data.Person, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.baseAddr, id), nil)
	if err != nil {
		return nil, fmt.Errorf("new get by id request for %s: %w", c.baseAddr, err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http/get by id response from %s: %w", c.baseAddr, err)
	}

	r := resp.Body
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read http/get response body from %s: %w", c.baseAddr, err)
	}

	ret := data.Person{}
	if err := jsoniter.Unmarshal(content, &ret); err != nil {
		return nil, fmt.Errorf("unmarshal http/get response body from %s: %w", c.baseAddr, err)
	}

	return &ret, nil
}

// Update updates the given person and returns
// the updated person and any errors.
func (c *personRestClient) Update(p *data.Person) (*data.Person, error) {
	content, err := jsoniter.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("marshal for %s: %w", c.baseAddr, err)
	}

	req, err := http.NewRequest("PUT", c.baseAddr, bytes.NewBuffer(content))
	if err != nil {
		return nil, fmt.Errorf("new put request for %s: %w", c.baseAddr, err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http/put response from %s: %w", c.baseAddr, err)
	}

	r := resp.Body
	defer resp.Body.Close()

	content, err = ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read http/put response body from %s: %w", c.baseAddr, err)
	}

	ret := data.Person{}
	if err := jsoniter.Unmarshal(content, &ret); err != nil {
		return nil, fmt.Errorf("unmarshal http/put response body from %s: %w", c.baseAddr, err)
	}

	return &ret, nil
}

// List runs an optional query and returns the results and any errors.
func (c *personRestClient) List(q map[string]interface{}) ([]data.Person, error) {
	var query string

	if q != nil {
		queryStr, err := jsoniter.Marshal(q)
		if err != nil {
			return nil, fmt.Errorf("internal error marshaling q: %w", err)
		}

		query = string(queryStr)
	}

	if query != "" {
		query = url.QueryEscape(query)
	}

	getURL := c.baseAddr

	if query != "" {
		getURL = fmt.Sprintf("%s?query=%s", c.baseAddr, query)
	}

	req, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new get request for %s: %w", getURL, err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http/get by id response from %s: %w", getURL, err)
	}

	r := resp.Body
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read http/get response body from %s: %w", getURL, err)
	}

	ret := make([]data.Person, 0)
	if err := jsoniter.Unmarshal(content, &ret); err != nil {
		return nil, fmt.Errorf("unmarshal http/get response body from %s: %w", getURL, err)
	}

	return ret, nil
}

// PartialUpdate performs a partial update of the data.Person with the given id using
// the given values.
func (c *personRestClient) PartialUpdate(id string, values map[string]interface{}) (*data.Person, error) {
	content, err := jsoniter.Marshal(values)
	if err != nil {
		return nil, fmt.Errorf("marshal for %s: %w", c.baseAddr, err)
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s?id=%s", c.baseAddr, id), bytes.NewBuffer(content))
	if err != nil {
		return nil, fmt.Errorf("new patch request for %s: %w", c.baseAddr, err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http/patch response from %s: %w", c.baseAddr, err)
	}

	r := resp.Body
	defer resp.Body.Close()

	content, err = ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read http/patch response body from %s: %w", c.baseAddr, err)
	}

	ret := data.Person{}
	if err := jsoniter.Unmarshal(content, &ret); err != nil {
		return nil, fmt.Errorf("unmarshal http/put response body from %s: %w", c.baseAddr, err)
	}

	return &ret, nil
}
