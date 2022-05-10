SHELL = /bin/bash
.DEFAULT_GOAL := all

.PHONY: all dep format vet test bench

all: dep format vet test bench

dep:
	go get -d ./...
	go mod tidy -v

format:
	gofmt -s -w -e -l .

vet:
	go vet ./...

test:
	go test -race -cover ./...

bench:
	go test -run=XXX -benchmem -bench=. ./...
