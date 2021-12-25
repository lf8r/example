#!/bin/bash
set -e
./clean.sh

# Remove and recreate the temp directory, since we want our grpc generator main
# function to be generated there.
rm -rf generated
mkdir generated
rm -f generated/grpc.proto
rm -rf temp
mkdir temp
mkdir temp/main

rm -rf book/ person/
basegen "github.com/lf8r/example-data/pkg/data1/Book[*:copyright=// Copyright (C) Subhajit DasGupta 2021|*:dir=generated|*:package=book|dao:package=bookdao|mw:file=../pkg/book/business_logic.go|rest:endpoint=/rest/books]"
basegen "github.com/lf8r/example-data/pkg/data/Person[*:copyright=// Copyright (C) Subhajit DasGupta 2021|*:dir=generated|*:package=person|dao:package=persondao|mw:file=../pkg/person/business_logic.go|rest:endpoint=/rest/persons]"
restfunnel "package=resthandler,saveDir=generated,saveFile=generated.rest.handler.go,types=/rest/persons:data.Person|/rest/books:data1.Book"
find . -name \*.go -exec goimports -w {} \;

# Generate the per-data-type proto bindings. Note that since main is being
# executed from temp/main/., we set the save dir to ../../generated/protobuf.
# Also, we want all of the code to be generated under a "common" package to
# avoid duplicating common.pb.go, which is used by both Person and Book (for
# common.Resource).
protogen github.com/lf8r/example-data/pkg/data1/Book ../../generated/protobuf common 
protogen github.com/lf8r/example-data/pkg/data/Person ../../generated/protobuf common 

rm -rf temp

# Build main to ensure no syntax errors in generated files.
rm -f example
CGO_ENABLED=0 go build -o example

