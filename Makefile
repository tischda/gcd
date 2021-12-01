# ---------------------------------------------------------------------------
# Makefile for CLI utilities
# ---------------------------------------------------------------------------

BUILD_TAG=$(shell git describe --tags 2>/dev/null || echo undefined)
LDFLAGS=-ldflags=all="-X main.version=${BUILD_TAG} -s -w"

all: build

build:
	go build ${LDFLAGS}

test:	clean
	go test -v -cover

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

install:
	go install ${LDFLAGS} ./...

update:
	go get -u
	go mod tidy
	# https://github.com/golang/go/issues/45161
	go mod vendor

snapshot:
	goreleaser --snapshot --skip-publish --rm-dist

release: clean
	# you must escape '#' and '$' characters
	@sed '1,/\#\#.*${BUILD_TAG}/d;/\#\#/Q' CHANGELOG.md | sed '/^$$/d' > releaseinfo
	goreleaser release --rm-dist --release-notes=releaseinfo
	@rm -f releaseinfo

clean:
	go clean
	rm -f releaseinfo
