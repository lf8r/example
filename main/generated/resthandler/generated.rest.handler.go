// Copyright (C) Subhajit DasGupta 2021

package resthandler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/lf8r/example/main/generated/book"
	"github.com/lf8r/example/main/generated/person"
)

// Warning - This is generated code. It is overwritten on each build.

//
// Common REST scaffolding.
//

// Types returns a map in which the key for each entry is a REST path and the
// value is a persisted struct (e.g. "/rest/persons").
func types() map[string]string {
	return map[string]string{

		"/rest/persons": "data.Person",
		"/rest/books":   "data1.Book",
	}
}

// Handler handles incoming REST calls.
// Handler handles incoming REST calls.
func Handler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/" || r.URL.Path == "/rest":
		w.Header().Add("Content-Type", "application/text")
		// Return a list of resource names.
		w.WriteHeader(http.StatusOK)

		for k := range types() {
			w.Write([]byte(fmt.Sprintf("%s\n", k)))
		}

		return

	case strings.HasPrefix(r.URL.Path, "/rest/persons"):
		person.Handler(w, r)

		return

	case strings.HasPrefix(r.URL.Path, "/rest/books"):
		book.Handler(w, r)

		return

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid path %s", r.URL.Path)))

		return
	}
}

// WriteInternalError writes an http internal error code and the given message
// to the response writer.
func WriteInternalError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(msg))
}

// WriteNotFound writes an http not found error code and the given message to
// the response writer.
func WriteNotFound(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(msg))
}
