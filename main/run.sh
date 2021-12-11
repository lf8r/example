#!/bin/bash
set -e
launcher -dbport=3030 -types=github.com/lf8r/example-data/pkg/data1/Book,github.com/lf8r/dbgen/pkg/data/Person -count=100 -cmd=./example1
