#!/bin/bash

set -e
set -x

go install github.com/GeertJohan/aerospike-udf-go/udfgo

udfgo --verbose build

# go build -buildmode c-shared -o hello.go.so

# echo "register module 'hello.go.so'" | aql

