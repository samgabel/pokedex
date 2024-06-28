.SILENT:

run:
	go run .

build:
	go build .

test:
	go test . -test.v

format:
	gofmt -w -d .
