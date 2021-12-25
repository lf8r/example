# Example application using generated data bindings

## Introduction
This example demonstrates how to "bind" externally defined data structs into arbitrary Go applications. This example binds the following externally defined types:
 - github.com/lf8r/example-data/pkg/data/Person
 - github.com/lf8r/example-data/pkg/data1/Book

The "bindings" (generated code) consists of:
 - A Go package named "book" containing: 
    - gRPC skeleton server and client, 
    - Protobuf bindings for data1.Book, 
    - REST handler and client for books (served at "/rest/books"), supporting "json", "yaml" and "text" formatted data representations. The handler supports HTTP/GET ("list", as well as "get by id"), POST, PUT, DELETE and PATCH (for partial updates).
 - A Go package named "bookdao" containing:
    - Posgres/JSONB bindings for persisting data1.Book objects.
    - A Client that adds transaction and business logic support over the raw Postgres/JSONB bindings.
 - Go packages named "person" and "persondao" containing similar artifacts for data.Person objects.