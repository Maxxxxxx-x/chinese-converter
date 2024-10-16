#!make

binary_name = convert
main_path = ./cmd/convert/main.go

.PHONY: tidy
tidy:
	go mod tidy
	go fmt ./...

.PHONY: build
build:
	go build -o=./bin/${binary_name} ${main_path}
