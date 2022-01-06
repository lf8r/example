#!/bin/bash
set -e
go mod tidy
cd main
ab -output=build.log
