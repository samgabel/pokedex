.SILENT:

.PHONY: run build test-repl test-pokecache format

run:
	go run .

build:
	go build .

test:
	go test ./...

test-repl:
	go test -v .

test-pokecache:
	go test -v ./internal/pokecache

format:
	gofmt -w -d .
