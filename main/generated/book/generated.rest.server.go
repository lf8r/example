// Copyright (C) Subhajit DasGupta 2021

package book

import (
	context "context"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"path"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/lf8r/example-data/pkg/data1"
	"github.com/lf8r/example/main/generated/bookdao"
	"github.com/lf8r/example/main/generated/persondao"

	"gopkg.in/yaml.v2"
)

// Warning - This is generated code. It is overwritten on each build.

type baseContextKey string

const (
	baseScheme baseContextKey = "basePath"
)

//
// Common REST scaffolding.
//

// Types returns a map in which the key for each entry is a REST path and the
// value is a persisted struct (e.g. "/rest/persons").
func types() map[string]string {
	return map[string]string{

		"/rest/books": "data1.Book",
	}
}

// Handler handles incoming REST calls.
func Handler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	if urlPath == "/" {
		w.Header().Add("Content-Type", "application/text")
		// Return a list of resource names.
		w.WriteHeader(http.StatusOK)

		for k := range types() {
			w.Write([]byte(fmt.Sprintf("%s\n", k)))
		}

		return
	}

	// Find the type of data to be served from the URLpath.
	dtype := ""
	basePath := ""

	for k, v := range types() {
		if k == urlPath || strings.HasPrefix(urlPath, k+"/") {
			dtype = v
			basePath = k

			break
		}
	}

	if dtype == "" {
		// Remove any traling "/" from urlPath.
		if strings.HasSuffix(urlPath, "/") {
			urlPath = path.Dir(urlPath)
		}

		// No exact match found for what the user was looking for,
		// so be helpful about what's available.
		matchingPaths := make([]string, 0)
		for k := range types() {
			if path.Dir(k) == urlPath {
				matchingPaths = append(matchingPaths, k)
			}
		}

		if len(matchingPaths) != 0 {
			w.Header().Add("Content-Type", "application/text")
			w.WriteHeader(http.StatusOK)

			for k := range types() {
				w.Write([]byte(fmt.Sprintf("%s\n", k)))
			}

			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid path %s", urlPath)))

		return
	}

	// Put the base path in the context.
	r = r.WithContext(context.WithValue(r.Context(), baseScheme, basePath))

	// Start a transaction for use by downstream calls.
	ctx, err := bookdao.BeginTx(r.Context())
	if err != nil {
		w.Header().Add("Content-Type", "application/text")

		log.Printf("Could not start DB transaction: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not start DB transaction"))

		return
	}

	// Put the context with the transaction into the request.
	r = r.WithContext(ctx)

	if dtype == bookdao.BookTypeID {
		handleBook(w, r)
	}
	// Per expected behavior, commitTx may fail (return an error) if the
	// transaction has been rolled back or committed by preceding code.
	defer bookdao.CommitTx(ctx)

}

// WriteResults writes the given results to the response writer.
func WriteResults(w http.ResponseWriter, accepts []string, results interface{}) {
	outputFormat := "application/json"

	if len(accepts) != 0 {
		for _, accept := range accepts {
			if accept == "application/yml" || accept == "application/yaml" {
				outputFormat = "application/yml"

				break
			}
		}
	}

	switch outputFormat {
	case "application/json":
		content, err := jsoniter.MarshalIndent(results, "", "    ")
		if err != nil {
			WriteInternalError(w, fmt.Sprintf("JSON marshaling error: %v", err))

			return
		}

		if _, err := w.Write(content); err != nil {
			log.Printf("Error writing results: %v\n", err)

			return
		}

	case "application/yml":
		content, err := yaml.Marshal(results)
		if err != nil {
			WriteInternalError(w, fmt.Sprintf("YAML marshaling error: %v", err))

			return
		}

		if _, err := w.Write(content); err != nil {
			log.Printf("Error writing results: %v\n", err)

			return
		}
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

//
// Entity specific handlers.
//

// handleBook handles requests for Book.
func handleBook(w http.ResponseWriter, r *http.Request) {
	c := bookdao.BookService(r.Context())

	switch r.Method {
	case http.MethodGet:
		handleBookGet(c, w, r)

		return

	case http.MethodPost:
		handleBookPost(c, w, r)

		return

	case http.MethodPut:
		handleBookPut(c, w, r)

		return

	case http.MethodDelete:
		handleBookDelete(c, w, r)

		return

	case http.MethodPatch:
		handleBookPatch(c, w, r)

		return
	}

	// Unhandled HTTP methods.
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(fmt.Sprintf("%s HTTP method not supported", r.Method)))

	if err := bookdao.RollbackTx(r.Context()); err != nil {
		log.Printf("RollbackTx error: %v\n", err)
	}
}

// handleBookGet handles HTTP/GET for Book.
func handleBookGet(c bookdao.BookClient, w http.ResponseWriter, r *http.Request) {
	// Get the base path schemeIf from the request context.
	schemeIf := r.Context().Value(baseScheme)
	if schemeIf == nil {
		log.Printf("Expected %s in context: \n", baseScheme)

		if err := bookdao.RollbackTx(r.Context()); err != nil {
			log.Printf("RollbackTx error: %v\n", err)
		}

		return
	}

	// Get the scheme (base URL) for the service.
	scheme, ok := schemeIf.(string)
	if !ok {
		log.Printf("Unexpected type of scheme in context, expected string.\n")

		if err := bookdao.RollbackTx(r.Context()); err != nil {
			log.Printf("RollbackTx error: %v\n", err)
		}

		return
	}

	urlPath := r.URL.Path

	var result interface{}

	// The GET REST API is used both for queries and for retrieval of structs by ID.
	if urlPath == scheme {
		// Query
		queryStr := r.URL.Query().Get("query")

		var queryM map[string]interface{}

		if queryStr != "" {
			queryM = make(map[string]interface{})
			if err := jsoniter.Unmarshal([]byte(queryStr), &queryM); err != nil {
				WriteInternalError(w, fmt.Sprintf("Internal error in list: %v", err))

				if err := persondao.RollbackTx(r.Context()); err != nil {
					log.Printf("RollbackTx error: %v\n", err)
				}

				return
			}
		}

		results, err := c.List(queryM)
		if err != nil {
			WriteInternalError(w, fmt.Sprintf("Internal error in list: %v", err))

			if err := bookdao.RollbackTx(r.Context()); err != nil {
				log.Printf("RollbackTx error: %v\n", err)
			}

			return
		}

		result = results
	} else {
		if path.Dir(urlPath) == scheme {
			// Retrieve by id.
			id := path.Base(urlPath)
			val, err := c.GetByID(id)
			if err != nil {
				WriteInternalError(w, fmt.Sprintf("Internal error in get %s: %v", id, err))

				if err := bookdao.RollbackTx(r.Context()); err != nil {
					log.Printf("RollbackTx error: %v\n", err)
				}

				return
			}

			result = val
		} else {
			// Can't process this URL since we have a sub-resource under the
			// base URL for this type of resource.
			WriteNotFound(w, fmt.Sprintf("Not found: %s\n", urlPath))

			if err := bookdao.RollbackTx(r.Context()); err != nil {
				log.Printf("RollbackTx error: %v\n", err)
			}

			return
		}
	}

	WriteResults(w, r.Header.Values("Accept"), result)
}

// handleBookPost handles HTTP/POST for Book.
func handleBookPost(c bookdao.BookClient, w http.ResponseWriter, r *http.Request) {
	val, err := readBookFromRequest(r)
	if err != nil {
		WriteInternalError(w, fmt.Sprintf("Internal error in post: %v", err))

		if err = bookdao.RollbackTx(r.Context()); err != nil {
			log.Printf("RollbackTx error: %v\n", err)
		}

		return
	}

	ret, err := c.Create(val)

	if err != nil {
		WriteInternalError(w, fmt.Sprintf("Internal error in post: %v", err))

		if err = bookdao.RollbackTx(r.Context()); err != nil {
			log.Printf("RollbackTx error: %v\n", err)
		}

		return
	}

	WriteResults(w, r.Header.Values("Accept"), ret)
}

// handleBookDelete handles HTTP/DELETE for Book.
func handleBookDelete(c bookdao.BookClient, w http.ResponseWriter, r *http.Request) {
	val, err := readBookFromRequest(r)
	if err != nil {
		WriteInternalError(w, fmt.Sprintf("Internal error in delete: %v", err))

		if err = bookdao.RollbackTx(r.Context()); err != nil {
			log.Printf("RollbackTx error: %v\n", err)
		}

		return
	}

	if err := c.Delete(val); err != nil {
		WriteInternalError(w, fmt.Sprintf("Internal error in delete: %v", err))

		if err = bookdao.RollbackTx(r.Context()); err != nil {
			log.Printf("RollbackTx error: %v\n", err)
		}

		return
	}

	WriteResults(w, r.Header.Values("Accept"), val)
}

// handleBookPut handles HTTP/PUT for Book.
func handleBookPut(c bookdao.BookClient, w http.ResponseWriter, r *http.Request) {
	val, err := readBookFromRequest(r)
	if err != nil {
		WriteInternalError(w, fmt.Sprintf("Internal error in put: %v", err))

		if err = bookdao.RollbackTx(r.Context()); err != nil {
			log.Printf("RollbackTx error: %v\n", err)
		}

		return
	}

	ret, err := c.Update(val)
	if err != nil {
		WriteInternalError(w, fmt.Sprintf("Internal error in update: %v", err))

		if err = bookdao.RollbackTx(r.Context()); err != nil {
			log.Printf("RollbackTx error: %v\n", err)
		}

		return
	}

	WriteResults(w, r.Header.Values("Accept"), ret)
}

// handleBookPatch handles HTTP/PATCH for Book.
func handleBookPatch(c bookdao.BookClient, w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteInternalError(w, fmt.Sprintf("Internal error in patch: %v", err))

		if err = bookdao.RollbackTx(r.Context()); err != nil {
			log.Printf("RollbackTx error: %v\n", err)
		}

		return
	}

	id := r.URL.Query().Get("id")

	val := make(map[string]interface{})
	if err = jsoniter.Unmarshal(content, &val); err != nil {
		WriteInternalError(w, fmt.Sprintf("Internal error in patch: %v", err))

		if err = bookdao.RollbackTx(r.Context()); err != nil {
			log.Printf("RollbackTx error: %v\n", err)
		}

		return
	}

	ret, err := c.PartialUpdate(id, val)
	if err != nil {
		WriteInternalError(w, fmt.Sprintf("Internal error in patch: %v", err))

		if err = bookdao.RollbackTx(r.Context()); err != nil {
			log.Printf("RollbackTx error: %v\n", err)
		}

		return
	}

	WriteResults(w, r.Header.Values("Accept"), ret)
}

// readBookFromRequest reads a Book from the given request. It returns an
// error if a Book couldn't be read.
func readBookFromRequest(r *http.Request) (*data1.Book, error) {
	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/json"
	}

	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, fmt.Errorf("parse media type %s: %w", contentType, err)
	}

	if mediaType != "" {
		contentType = mediaType
	}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("read request body: %w", err)
	}

	val := data1.Book{}

	switch contentType {
	case "application/json":
		if err := jsoniter.Unmarshal(content, &val); err != nil {
			return nil, fmt.Errorf("unmarshal json: %w", err)
		}

	case "application/yml", "application/yaml":
		if err := yaml.Unmarshal(content, &val); err != nil {
			return nil, fmt.Errorf("unmarshal yaml: %w", err)
		}

	default:
		return nil, fmt.Errorf("unhandled content type in request: %s", contentType)
	}

	return &val, nil
}
