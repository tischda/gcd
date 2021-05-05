# ---------------------------------------------------------------------------
# Makefile for GO utilities
# ---------------------------------------------------------------------------

PROJECT_DIR=$(notdir $(shell pwd))

all: build

build: 
	go build -ldflags "all=-s -w"


test:
	go test -v

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

fmt:
	go fmt

dist: clean build

clean:
	go clean
