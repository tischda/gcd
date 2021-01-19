# ---------------------------------------------------------------------------
# Makefile for GO utilities
# ---------------------------------------------------------------------------

PROJECT_DIR=$(notdir $(shell pwd))

all: build

build: 
	go build

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

fmt:
	go fmt

dist: clean build

clean:
	go clean
