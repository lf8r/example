#!/bin/bash
set -e
#./generate.sh
CGO_ENABLED=0 go build -o example
