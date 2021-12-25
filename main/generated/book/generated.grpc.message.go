// Copyright (C) Subhajit DasGupta 2021

package book

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/lf8r/example-data/pkg/data1"
)

//
// CreateBookRequest
//

type CreateBookRequest struct {
	Body *data1.Book
}

func NewCreateBookRequest(b *data1.Book) (*CreateBookRequest, error) {
	return &CreateBookRequest{
		Body: b,
	}, nil
}

func ParseCreateBookRequest(req *CreateBookRequest) (*data1.Book, error) {
	return req.Body, nil
}

var _ proto.Message = (*CreateBookRequest)(nil)

func (p *CreateBookRequest) Reset() {}

func (p *CreateBookRequest) String() string {
	return ""
}

func (p *CreateBookRequest) ProtoMessage() {}

//
// CreateBookResponse
//

type CreateBookResponse struct {
	Body *data1.Book
	Err  string
}

func ParseCreateBookResponse(resp *CreateBookResponse) (*data1.Book, error) {
	if resp.Err != "" {
		return nil, fmt.Errorf("remote error in CreateBookResponse %s", resp.Err)
	}

	return resp.Body, nil
}

var _ proto.Message = (*CreateBookResponse)(nil)

func (p *CreateBookResponse) Reset() {}

func (p *CreateBookResponse) String() string {
	return ""
}

func (p *CreateBookResponse) ProtoMessage() {}

//
// DeleteBookRequest.
//
type DeleteBookRequest struct {
	Body *data1.Book
}

func NewDeleteBookRequest(b *data1.Book) (*DeleteBookRequest, error) {
	return &DeleteBookRequest{
		Body: b,
	}, nil
}

func ParseDeleteBookRequest(req *DeleteBookRequest) (*data1.Book, error) {
	return req.Body, nil
}

var _ proto.Message = (*DeleteBookRequest)(nil)

func (p *DeleteBookRequest) Reset() {}

func (p *DeleteBookRequest) String() string {
	return ""
}

func (p *DeleteBookRequest) ProtoMessage() {}

//
// DeleteBookResponse.
//

type DeleteBookResponse struct {
	Err string
}

func ParseDeleteBookResponse(resp *DeleteBookResponse) error {
	if resp.Err != "" {
		return fmt.Errorf("remote error in DeleteBookResponse %s", resp.Err)
	}

	return nil
}

var _ proto.Message = (*DeleteBookResponse)(nil)

func (p *DeleteBookResponse) Reset() {}

func (p *DeleteBookResponse) String() string {
	return ""
}

func (p *DeleteBookResponse) ProtoMessage() {}

//
// DeleteByIDBookRequest.
//
type DeleteByIDBookRequest struct {
	Body string
}

func NewDeleteByIDBookRequest(b *data1.Book) (*DeleteByIDBookRequest, error) {
	return &DeleteByIDBookRequest{
		Body: b.ID,
	}, nil
}

func ParseDeleteByIDBookRequest(req *DeleteByIDBookRequest) (string, error) {
	return req.Body, nil
}

var _ proto.Message = (*DeleteByIDBookRequest)(nil)

func (p *DeleteByIDBookRequest) Reset() {}

func (p *DeleteByIDBookRequest) String() string {
	return ""
}

func (p *DeleteByIDBookRequest) ProtoMessage() {}

//
// DeleteByIDBookResponse.
//

type DeleteByIDBookResponse struct {
	Err string
}

func ParseDeleteByIDBookResponse(resp *DeleteByIDBookResponse) error {
	if resp.Err != "" {
		return fmt.Errorf("remote error in DeleteByIDBookResponse %s", resp.Err)
	}

	return nil
}

var _ proto.Message = (*DeleteByIDBookResponse)(nil)

func (p *DeleteByIDBookResponse) Reset() {}

func (p *DeleteByIDBookResponse) String() string {
	return ""
}

//
// GetByIDBookRequest.
//
type GetByIDBookRequest struct {
	Body string
}

func NewGetByIDBookRequest(id string) (*GetByIDBookRequest, error) {
	return &GetByIDBookRequest{
		Body: id,
	}, nil
}

func ParseGetByIDBookRequest(req *GetByIDBookRequest) (string, error) {
	return req.Body, nil
}

var _ proto.Message = (*GetByIDBookRequest)(nil)

func (p *GetByIDBookRequest) Reset() {}

func (p *GetByIDBookRequest) String() string {
	return ""
}

func (p *GetByIDBookRequest) ProtoMessage() {}

//
// GetByIDBookResponse.
//
type GetByIDBookResponse struct {
	Body *data1.Book
	Err  string
}

func ParseGetByIDBookResponse(resp *GetByIDBookResponse) (*data1.Book, error) {
	if resp.Err != "" {
		return nil, fmt.Errorf("remote error in GetByIDBookResponse %s", resp.Err)
	}

	return resp.Body, nil
}

var _ proto.Message = (*GetByIDBookResponse)(nil)

func (p *GetByIDBookResponse) Reset() {}

func (p *GetByIDBookResponse) String() string {
	return ""
}

func (p *GetByIDBookResponse) ProtoMessage() {}

func (p *DeleteByIDBookResponse) ProtoMessage() {}

//
// ListBookRequest.
//

type ListBookRequest struct {
	Query map[string]interface{}
}

func NewListBookRequest(query map[string]interface{}) (*ListBookRequest, error) {
	return &ListBookRequest{
		Query: query,
	}, nil
}

func ParseListBookRequest(req *ListBookRequest) (map[string]interface{}, error) {
	return req.Query, nil
}

var _ proto.Message = (*ListBookRequest)(nil)

func (p *ListBookRequest) Reset() {}

func (p *ListBookRequest) String() string {
	return ""
}

func (p *ListBookRequest) ProtoMessage() {}

//
// ListBookResponse.
//

type ListBookResponse struct {
	Body []data1.Book
	Err  string
}

func ParseListBookResponse(resp *ListBookResponse) ([]data1.Book, error) {
	if resp.Err != "" {
		return nil, fmt.Errorf("remote error in ListBookResponse %s", resp.Err)
	}

	return resp.Body, nil
}

var _ proto.Message = (*ListBookResponse)(nil)

func (p *ListBookResponse) Reset() {}

func (p *ListBookResponse) String() string {
	return ""
}

func (p *ListBookResponse) ProtoMessage() {}

//
// UpdateBookRequest.
//
type UpdateBookRequest struct {
	Body *data1.Book
}

func NewUpdateBookRequest(b *data1.Book) (*UpdateBookRequest, error) {
	return &UpdateBookRequest{
		Body: b,
	}, nil
}

func ParseUpdateBookRequest(req *UpdateBookRequest) (*data1.Book, error) {
	return req.Body, nil
}

var _ proto.Message = (*UpdateBookRequest)(nil)

func (p *UpdateBookRequest) Reset() {}

func (p *UpdateBookRequest) String() string {
	return ""
}

func (p *UpdateBookRequest) ProtoMessage() {}

//
// UpdateBookResponse.
//
type UpdateBookResponse struct {
	Body *data1.Book
	Err  string
}

func ParseUpdateBookResponse(resp *UpdateBookResponse) (*data1.Book, error) {
	if resp.Err != "" {
		return nil, fmt.Errorf("remote error in UpdateBookResponse %s", resp.Err)
	}

	return resp.Body, nil
}

var _ proto.Message = (*UpdateBookResponse)(nil)

func (p *UpdateBookResponse) Reset() {}

func (p *UpdateBookResponse) String() string {
	return ""
}

func (p *UpdateBookResponse) ProtoMessage() {}

//
// PartialUpdateBookRequest.
//

type PartialUpdateBookRequest struct {
	Id   string
	Body map[string]interface{}
}

func NewPartialUpdateBookRequest(id string, b map[string]interface{}) (*PartialUpdateBookRequest, error) {
	return &PartialUpdateBookRequest{
		Body: b,
		Id:   id,
	}, nil
}

func ParsePartialUpdateBookRequest(req *PartialUpdateBookRequest) (string, map[string]interface{}, error) {
	return req.Id, req.Body, nil
}

var _ proto.Message = (*PartialUpdateBookRequest)(nil)

func (p *PartialUpdateBookRequest) Reset() {}

func (p *PartialUpdateBookRequest) String() string {
	return ""
}

func (p *PartialUpdateBookRequest) ProtoMessage() {}

//
// PartialUpdateBookResponse.
//

type PartialUpdateBookResponse struct {
	Body *data1.Book
	Err  string
}

func ParsePartialUpdateBookResponse(resp *PartialUpdateBookResponse) (*data1.Book, error) {
	if resp.Err != "" {
		return nil, fmt.Errorf("remote error in PartialUpdateBookResponse %s", resp.Err)
	}

	return resp.Body, nil
}

var _ proto.Message = (*PartialUpdateBookResponse)(nil)

func (p *PartialUpdateBookResponse) Reset() {}

func (p *PartialUpdateBookResponse) String() string {
	return ""
}

func (p *PartialUpdateBookResponse) ProtoMessage() {}
