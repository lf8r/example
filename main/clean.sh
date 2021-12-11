#!/bin/bash
set -e
rm -rf generated/*
rm -rf book/*
rm -rf bookdao/*
rm -rf person/*
rm -rf persondao/*

# Remove and recreate the temp directory, since we want our grpc generator main
# function to be generated there.
rm -f generated/grpc.proto
rm -rf temp

# Remove all proto files.
rm -f *.proto

# Remove all .pb.go files
rm -rf *.pb.go

# Remove the generated grpc.proto file since we no longer need it.
rm -f generated/grpc.proto

# Init grpc tooling by downloading protoc-gen-go and protoc-gen-go-grpc.
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# Move these files to a directory in the path (~/bin in this example).
mv ~/go/bin/protoc-gen-go ~/bin/.
mv ~/go/bin/protoc-gen-go-grpc ~/bin/.

rm -f example1

rm -f *.pprof
rm -f *.pdf
