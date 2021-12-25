// Copyright (C) Subhajit DasGupta 2021

package person

import (
	context "context"
	"fmt"

	"github.com/lf8r/example/main/generated/persondao"
)

type Server struct {
	UnimplementedServiceServer
}

var _ ServiceServer = (*Server)(nil)

func (g Server) CreatePerson(ctx context.Context, req *CreatePersonRequest) (*CreatePersonResponse, error) {
	resp := CreatePersonResponse{}

	val, err := ParseCreatePersonRequest(req)
	if err != nil {
		resp.Err = fmt.Errorf("parse CreatePerson request: %w", err).Error()

		return &resp, nil
	}

	ctx, err = persondao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx CreatePerson: %w", err).Error()

		return &resp, nil
	}

	defer persondao.CommitTx(ctx)

	client := persondao.PersonService(ctx)

	val1, err := client.Create(val)
	if err != nil {
		resp.Err = fmt.Errorf("CreatePerson: %w", err).Error()

		if err := persondao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	resp.Body = val1

	return &resp, nil
}

func (g Server) DeletePerson(ctx context.Context, req *DeletePersonRequest) (*DeletePersonResponse, error) {
	resp := DeletePersonResponse{}

	val, err := ParseDeletePersonRequest(req)
	if err != nil {
		resp.Err = fmt.Errorf("parse DeletePerson request: %w", err).Error()

		return &resp, nil
	}

	ctx, err = persondao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx DeletePerson: %w", err).Error()

		return &resp, nil
	}

	defer persondao.CommitTx(ctx)

	client := persondao.PersonService(ctx)

	err = client.Delete(val)
	if err != nil {
		resp.Err = fmt.Errorf("DeletePerson: %w", err).Error()

		if err := persondao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	return &resp, nil
}

func (g Server) DeleteByIDPerson(ctx context.Context, req *DeleteByIDPersonRequest) (*DeleteByIDPersonResponse, error) {
	resp := DeleteByIDPersonResponse{}

	id, err := ParseDeleteByIDPersonRequest(req)
	if err != nil {
		resp.Err = fmt.Errorf("parse DeleteByIDPerson request: %w", err).Error()

		return &resp, nil
	}

	ctx, err = persondao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx DeleteByIDPerson: %w", err).Error()

		return &resp, nil
	}

	defer persondao.CommitTx(ctx)

	client := persondao.PersonService(ctx)

	err = client.DeleteByID(id)
	if err != nil {
		resp.Err = fmt.Errorf("DeleteByIDPerson: %w", err).Error()

		if err := persondao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	return &resp, nil
}

func (g Server) GetByIDPerson(ctx context.Context, req *GetByIDPersonRequest) (*GetByIDPersonResponse, error) {
	resp := GetByIDPersonResponse{}

	id, err := ParseGetByIDPersonRequest(req)
	if err != nil {
		resp.Err = fmt.Errorf("parse GetByIDPerson request: %w", err).Error()

		return &resp, nil
	}

	ctx, err = persondao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx GetByIDPerson: %w", err).Error()

		return &resp, nil
	}

	defer persondao.CommitTx(ctx)

	client := persondao.PersonService(ctx)

	val, err := client.GetByID(id)
	if err != nil {
		resp.Err = fmt.Errorf("GetByIDPerson: %w", err).Error()

		if err := persondao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	resp.Body = val

	return &resp, nil
}

func (g Server) UpdatePerson(ctx context.Context, req *UpdatePersonRequest) (*UpdatePersonResponse, error) {
	resp := UpdatePersonResponse{}

	val, err := ParseUpdatePersonRequest(req)
	if err != nil {
		resp.Err = fmt.Errorf("parse UpdatePerson request: %w", err).Error()

		return &resp, nil
	}

	ctx, err = persondao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx UpdatePerson: %w", err).Error()

		return &resp, nil
	}

	defer persondao.CommitTx(ctx)

	client := persondao.PersonService(ctx)

	val1, err := client.Update(val)
	if err != nil {
		resp.Err = fmt.Errorf("UpdatePerson: %w", err).Error()

		if err := persondao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	resp.Body = val1

	return &resp, nil
}

func (g Server) ListPerson(ctx context.Context, req *ListPersonRequest) (*ListPersonResponse, error) {
	resp := ListPersonResponse{}

	query := req.Query

	ctx, err := persondao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx ListPerson: %w", err).Error()

		return &resp, nil
	}

	defer persondao.CommitTx(ctx)

	client := persondao.PersonService(ctx)

	val, err := client.List(query)
	if err != nil {
		resp.Err = fmt.Errorf("ListPerson: %w", err).Error()

		if err := persondao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	resp.Body = val

	return &resp, nil
}

func (g Server) PartialUpdatePerson(ctx context.Context, req *PartialUpdatePersonRequest) (*PartialUpdatePersonResponse, error) {
	resp := PartialUpdatePersonResponse{}

	id := req.Id
	values := req.Body

	ctx, err := persondao.BeginTx(ctx)
	if err != nil {
		resp.Err = fmt.Errorf("begin tx PartialUpdatePerson: %w", err).Error()

		return &resp, nil
	}

	defer persondao.CommitTx(ctx)

	client := persondao.PersonService(ctx)

	val, err := client.PartialUpdate(id, values)
	if err != nil {
		resp.Err = fmt.Errorf("PartialUpdatePerson: %w", err).Error()

		if err := persondao.RollbackTx(ctx); err != nil {
			return &resp, err
		}

		return &resp, nil
	}

	resp.Body = val

	return &resp, nil
}
