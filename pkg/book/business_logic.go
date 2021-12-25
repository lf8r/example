// Copyright (C) Subhajit DasGupta 2021
package book

import (
	"context"
)

// Note that this code is generated once, after which it is not overwritten since
// it's intended to contain user modifications for business logic.

// CreateBusinessLogic is called back by the generated Book Service just before
// creating the given struct p in the persistent store. If this function returns
// an error, the struct is not created. This function may use downcasts to
// determine the struct type being called back for and adjust the business logic
// suitably as shown below.
//
//  //Show the type of object we're being called to delete (p's type).
//	input := ""
//
//	switch p.(type) {
//	case data.Person:
//		input = "person"
//
//	case *data.Person:
//		input = "*person"
//
// 	case data1.Book:
//		input = "book"
//
//	case *data1.Book:
//		input = "*book"
func CreateBusinessLogic(ctx context.Context, p interface{}) error {
	return nil
}

// UpdateBusinessLogic is called back by the generated Book Service just
// before updating the given struct p in the persistent store. If this function
// returns an error, the struct is not updated. This function may use downcasts
// to determine the struct type being called back for and adjust the business
// logic suitably. See CreateBusinessLogic for example code.
func UpdateBusinessLogic(ctx context.Context, p interface{}) error {
	return nil
}

// ListBusinessLogic is called back by the generated Book Service just
// before running the query against the given struct type in the persistent
// store. If this function returns an error, the query is not run. This function
// may use downcasts to determine the struct type being called back for and
// adjust the business logic suitably. See CreateBusinessLogic for example code.
func ListBusinessLogic(ctx context.Context, p interface{}, args ...interface{}) error {
	return nil
}

// DeleteBusinessLogic is called back by the generated Book Service just
// before deleting the given struct from the persistent store. If this function
// returns an error, the struct is not deleted. This function may use downcasts
// to determine the struct type being called back for and adjust the business
// logic suitably. See CreateBusinessLogic for example code.
func DeleteBusinessLogic(ctx context.Context, p interface{}) error {
	return nil
}

// GetByIDBusinessLogic is called back by the generated Book Service
// just before retrieving the struct with the given id from the persistent
// store. If this function returns an error, the struct is not retrieved. This
// function may use downcasts to determine the struct type being called back for
// and adjust the business logic suitably. See CreateBusinessLogic for example code.
func GetByIDBusinessLogic(ctx context.Context, p interface{}, id string) error {
	return nil
}
