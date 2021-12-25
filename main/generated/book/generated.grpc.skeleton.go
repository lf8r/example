// Copyright (C) Subhajit DasGupta 2021

package book

import (
	context "context"
	"fmt"

	"github.com/lf8r/example/main/generated/bookdao"
)

type Server struct {
	UnimplementedServiceServer
}

var _ ServiceServer = (*Server)(nil)

func (g Server) CreateBook(ctx context.Context, req *CreateBookRequest) (*CreateBookResponse, error) {
	resp := CreateBookResponse{}

	val, err := ParseCreateBookRequest(req)
	if err != nil {
		resp.Err = fmt.Errorf("parse CreateBook request: %w", err).Error()

		return &resp, nil
	}

	ctx, err = bookdao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx CreateBook: %w", err).Error()

		return &resp, nil
	}

	defer bookdao.CommitTx(ctx)

	client := bookdao.BookService(ctx)

	val1, err := client.Create(val)
	if err != nil {
		resp.Err = fmt.Errorf("CreateBook: %w", err).Error()

		if err := bookdao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	resp.Body = val1

	return &resp, nil
}

func (g Server) DeleteBook(ctx context.Context, req *DeleteBookRequest) (*DeleteBookResponse, error) {
	resp := DeleteBookResponse{}

	val, err := ParseDeleteBookRequest(req)
	if err != nil {
		resp.Err = fmt.Errorf("parse DeleteBook request: %w", err).Error()

		return &resp, nil
	}

	ctx, err = bookdao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx DeleteBook: %w", err).Error()

		return &resp, nil
	}

	defer bookdao.CommitTx(ctx)

	client := bookdao.BookService(ctx)

	err = client.Delete(val)
	if err != nil {
		resp.Err = fmt.Errorf("DeleteBook: %w", err).Error()

		if err := bookdao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	return &resp, nil
}

func (g Server) DeleteByIDBook(ctx context.Context, req *DeleteByIDBookRequest) (*DeleteByIDBookResponse, error) {
	resp := DeleteByIDBookResponse{}

	id, err := ParseDeleteByIDBookRequest(req)
	if err != nil {
		resp.Err = fmt.Errorf("parse DeleteByIDBook request: %w", err).Error()

		return &resp, nil
	}

	ctx, err = bookdao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx DeleteByIDBook: %w", err).Error()

		return &resp, nil
	}

	defer bookdao.CommitTx(ctx)

	client := bookdao.BookService(ctx)

	err = client.DeleteByID(id)
	if err != nil {
		resp.Err = fmt.Errorf("DeleteByIDBook: %w", err).Error()

		if err := bookdao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	return &resp, nil
}

func (g Server) GetByIDBook(ctx context.Context, req *GetByIDBookRequest) (*GetByIDBookResponse, error) {
	resp := GetByIDBookResponse{}

	id, err := ParseGetByIDBookRequest(req)
	if err != nil {
		resp.Err = fmt.Errorf("parse GetByIDBook request: %w", err).Error()

		return &resp, nil
	}

	ctx, err = bookdao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx GetByIDBook: %w", err).Error()

		return &resp, nil
	}

	defer bookdao.CommitTx(ctx)

	client := bookdao.BookService(ctx)

	val, err := client.GetByID(id)
	if err != nil {
		resp.Err = fmt.Errorf("GetByIDBook: %w", err).Error()

		if err := bookdao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	resp.Body = val

	return &resp, nil
}

func (g Server) UpdateBook(ctx context.Context, req *UpdateBookRequest) (*UpdateBookResponse, error) {
	resp := UpdateBookResponse{}

	val, err := ParseUpdateBookRequest(req)
	if err != nil {
		resp.Err = fmt.Errorf("parse UpdateBook request: %w", err).Error()

		return &resp, nil
	}

	ctx, err = bookdao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx UpdateBook: %w", err).Error()

		return &resp, nil
	}

	defer bookdao.CommitTx(ctx)

	client := bookdao.BookService(ctx)

	val1, err := client.Update(val)
	if err != nil {
		resp.Err = fmt.Errorf("UpdateBook: %w", err).Error()

		if err := bookdao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	resp.Body = val1

	return &resp, nil
}

func (g Server) ListBook(ctx context.Context, req *ListBookRequest) (*ListBookResponse, error) {
	resp := ListBookResponse{}

	query := req.Query

	ctx, err := bookdao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx ListBook: %w", err).Error()

		return &resp, nil
	}

	defer bookdao.CommitTx(ctx)

	client := bookdao.BookService(ctx)

	val, err := client.List(query)
	if err != nil {
		resp.Err = fmt.Errorf("ListBook: %w", err).Error()

		if err := bookdao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	resp.Body = val

	return &resp, nil
}

func (g Server) PartialUpdateBook(ctx context.Context, req *PartialUpdateBookRequest) (*PartialUpdateBookResponse, error) {
	resp := PartialUpdateBookResponse{}

	id := req.Id
	values := req.Body

	ctx, err := bookdao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx PartialUpdateBook: %w", err).Error()

		return &resp, nil
	}

	defer bookdao.CommitTx(ctx)

	client := bookdao.BookService(ctx)

	val, err := client.PartialUpdate(id, values)
	if err != nil {
		resp.Err = fmt.Errorf("PartialUpdateBook: %w", err).Error()

		if err := bookdao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	resp.Body = val

	return &resp, nil
}
