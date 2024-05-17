export PATH:=$(PATH):$(shell go env GOPATH)/bin

include .env

start-dev: go run main.go

test:
	go test ./...

runbuild:
	go build -o bin/products-api main.go

start-prod: runbuild
	./bin/products-api start