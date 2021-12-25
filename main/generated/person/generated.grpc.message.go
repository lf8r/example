// Copyright (C) Subhajit DasGupta 2021

package person

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/lf8r/example-data/pkg/data"
)

//
// CreatePersonRequest
//

type CreatePersonRequest struct {
	Body *data.Person
}

func NewCreatePersonRequest(b *data.Person) (*CreatePersonRequest, error) {
	return &CreatePersonRequest{
		Body: b,
	}, nil
}

func ParseCreatePersonRequest(req *CreatePersonRequest) (*data.Person, error) {
	return req.Body, nil
}

var _ proto.Message = (*CreatePersonRequest)(nil)

func (p *CreatePersonRequest) Reset() {}

func (p *CreatePersonRequest) String() string {
	return ""
}

func (p *CreatePersonRequest) ProtoMessage() {}

//
// CreatePersonResponse
//

type CreatePersonResponse struct {
	Body *data.Person
	Err  string
}

func ParseCreatePersonResponse(resp *CreatePersonResponse) (*data.Person, error) {
	if resp.Err != "" {
		return nil, fmt.Errorf("remote error in CreatePersonResponse %s", resp.Err)
	}

	return resp.Body, nil
}

var _ proto.Message = (*CreatePersonResponse)(nil)

func (p *CreatePersonResponse) Reset() {}

func (p *CreatePersonResponse) String() string {
	return ""
}

func (p *CreatePersonResponse) ProtoMessage() {}

//
// DeletePersonRequest.
//
type DeletePersonRequest struct {
	Body *data.Person
}

func NewDeletePersonRequest(b *data.Person) (*DeletePersonRequest, error) {
	return &DeletePersonRequest{
		Body: b,
	}, nil
}

func ParseDeletePersonRequest(req *DeletePersonRequest) (*data.Person, error) {
	return req.Body, nil
}

var _ proto.Message = (*DeletePersonRequest)(nil)

func (p *DeletePersonRequest) Reset() {}

func (p *DeletePersonRequest) String() string {
	return ""
}

func (p *DeletePersonRequest) ProtoMessage() {}

//
// DeletePersonResponse.
//

type DeletePersonResponse struct {
	Err string
}

func ParseDeletePersonResponse(resp *DeletePersonResponse) error {
	if resp.Err != "" {
		return fmt.Errorf("remote error in DeletePersonResponse %s", resp.Err)
	}

	return nil
}

var _ proto.Message = (*DeletePersonResponse)(nil)

func (p *DeletePersonResponse) Reset() {}

func (p *DeletePersonResponse) String() string {
	return ""
}

func (p *DeletePersonResponse) ProtoMessage() {}

//
// DeleteByIDPersonRequest.
//
type DeleteByIDPersonRequest struct {
	Body string
}

func NewDeleteByIDPersonRequest(b *data.Person) (*DeleteByIDPersonRequest, error) {
	return &DeleteByIDPersonRequest{
		Body: b.ID,
	}, nil
}

func ParseDeleteByIDPersonRequest(req *DeleteByIDPersonRequest) (string, error) {
	return req.Body, nil
}

var _ proto.Message = (*DeleteByIDPersonRequest)(nil)

func (p *DeleteByIDPersonRequest) Reset() {}

func (p *DeleteByIDPersonRequest) String() string {
	return ""
}

func (p *DeleteByIDPersonRequest) ProtoMessage() {}

//
// DeleteByIDPersonResponse.
//

type DeleteByIDPersonResponse struct {
	Err string
}

func ParseDeleteByIDPersonResponse(resp *DeleteByIDPersonResponse) error {
	if resp.Err != "" {
		return fmt.Errorf("remote error in DeleteByIDPersonResponse %s", resp.Err)
	}

	return nil
}

var _ proto.Message = (*DeleteByIDPersonResponse)(nil)

func (p *DeleteByIDPersonResponse) Reset() {}

func (p *DeleteByIDPersonResponse) String() string {
	return ""
}

//
// GetByIDPersonRequest.
//
type GetByIDPersonRequest struct {
	Body string
}

func NewGetByIDPersonRequest(id string) (*GetByIDPersonRequest, error) {
	return &GetByIDPersonRequest{
		Body: id,
	}, nil
}

func ParseGetByIDPersonRequest(req *GetByIDPersonRequest) (string, error) {
	return req.Body, nil
}

var _ proto.Message = (*GetByIDPersonRequest)(nil)

func (p *GetByIDPersonRequest) Reset() {}

func (p *GetByIDPersonRequest) String() string {
	return ""
}

func (p *GetByIDPersonRequest) ProtoMessage() {}

//
// GetByIDPersonResponse.
//
type GetByIDPersonResponse struct {
	Body *data.Person
	Err  string
}

func ParseGetByIDPersonResponse(resp *GetByIDPersonResponse) (*data.Person, error) {
	if resp.Err != "" {
		return nil, fmt.Errorf("remote error in GetByIDPersonResponse %s", resp.Err)
	}

	return resp.Body, nil
}

var _ proto.Message = (*GetByIDPersonResponse)(nil)

func (p *GetByIDPersonResponse) Reset() {}

func (p *GetByIDPersonResponse) String() string {
	return ""
}

func (p *GetByIDPersonResponse) ProtoMessage() {}

func (p *DeleteByIDPersonResponse) ProtoMessage() {}

//
// ListPersonRequest.
//

type ListPersonRequest struct {
	Query map[string]interface{}
}

func NewListPersonRequest(query map[string]interface{}) (*ListPersonRequest, error) {
	return &ListPersonRequest{
		Query: query,
	}, nil
}

func ParseListPersonRequest(req *ListPersonRequest) (map[string]interface{}, error) {
	return req.Query, nil
}

var _ proto.Message = (*ListPersonRequest)(nil)

func (p *ListPersonRequest) Reset() {}

func (p *ListPersonRequest) String() string {
	return ""
}

func (p *ListPersonRequest) ProtoMessage() {}

//
// ListPersonResponse.
//

type ListPersonResponse struct {
	Body []data.Person
	Err  string
}

func ParseListPersonResponse(resp *ListPersonResponse) ([]data.Person, error) {
	if resp.Err != "" {
		return nil, fmt.Errorf("remote error in ListPersonResponse %s", resp.Err)
	}

	return resp.Body, nil
}

var _ proto.Message = (*ListPersonResponse)(nil)

func (p *ListPersonResponse) Reset() {}

func (p *ListPersonResponse) String() string {
	return ""
}

func (p *ListPersonResponse) ProtoMessage() {}

//
// UpdatePersonRequest.
//
type UpdatePersonRequest struct {
	Body *data.Person
}

func NewUpdatePersonRequest(b *data.Person) (*UpdatePersonRequest, error) {
	return &UpdatePersonRequest{
		Body: b,
	}, nil
}

func ParseUpdatePersonRequest(req *UpdatePersonRequest) (*data.Person, error) {
	return req.Body, nil
}

var _ proto.Message = (*UpdatePersonRequest)(nil)

func (p *UpdatePersonRequest) Reset() {}

func (p *UpdatePersonRequest) String() string {
	return ""
}

func (p *UpdatePersonRequest) ProtoMessage() {}

//
// UpdatePersonResponse.
//
type UpdatePersonResponse struct {
	Body *data.Person
	Err  string
}

func ParseUpdatePersonResponse(resp *UpdatePersonResponse) (*data.Person, error) {
	if resp.Err != "" {
		return nil, fmt.Errorf("remote error in UpdatePersonResponse %s", resp.Err)
	}

	return resp.Body, nil
}

var _ proto.Message = (*UpdatePersonResponse)(nil)

func (p *UpdatePersonResponse) Reset() {}

func (p *UpdatePersonResponse) String() string {
	return ""
}

func (p *UpdatePersonResponse) ProtoMessage() {}

//
// PartialUpdatePersonRequest.
//

type PartialUpdatePersonRequest struct {
	Id   string
	Body map[string]interface{}
}

func NewPartialUpdatePersonRequest(id string, b map[string]interface{}) (*PartialUpdatePersonRequest, error) {
	return &PartialUpdatePersonRequest{
		Body: b,
		Id:   id,
	}, nil
}

func ParsePartialUpdatePersonRequest(req *PartialUpdatePersonRequest) (string, map[string]interface{}, error) {
	return req.Id, req.Body, nil
}

var _ proto.Message = (*PartialUpdatePersonRequest)(nil)

func (p *PartialUpdatePersonRequest) Reset() {}

func (p *PartialUpdatePersonRequest) String() string {
	return ""
}

func (p *PartialUpdatePersonRequest) ProtoMessage() {}

//
// PartialUpdatePersonResponse.
//

type PartialUpdatePersonResponse struct {
	Body *data.Person
	Err  string
}

func ParsePartialUpdatePersonResponse(resp *PartialUpdatePersonResponse) (*data.Person, error) {
	if resp.Err != "" {
		return nil, fmt.Errorf("remote error in PartialUpdatePersonResponse %s", resp.Err)
	}

	return resp.Body, nil
}

var _ proto.Message = (*PartialUpdatePersonResponse)(nil)

func (p *PartialUpdatePersonResponse) Reset() {}

func (p *PartialUpdatePersonResponse) String() string {
	return ""
}

func (p *PartialUpdatePersonResponse) ProtoMessage() {}
