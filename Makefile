BINPATH ?= build

.PHONY: all
all: test build

.PHONY: build
build:
	go build -o $(BINPATH)/go-helloworld-api

.PHONY: debug
debug:
	go build -tags 'debug' -o $(BINPATH)/go-helloworld-api
	HUMAN_LOG=1 DEBUG=1 $(BINPATH)/go-helloworld-api

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: convey
convey:
	goconvey ./...

