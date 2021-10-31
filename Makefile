lint:
	go fmt ./...

test:
	go test -race ./...

cover:
	go test -coverprofile cover.out ./... && go tool cover -html cover.out -o cover.html

generate:
	go generate ./...

all: generate lint test


