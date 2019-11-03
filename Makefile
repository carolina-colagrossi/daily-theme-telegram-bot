GO ?= go

all: build

fmt:
	$(GO) fmt ./...

test:
	$(GO) test ./pkg/...

build: fmt
build: test
build:
	$(GO) build -o bin/theme-bot ./cmd

clean:
	rm -r ./bin

.PHONY: clean
