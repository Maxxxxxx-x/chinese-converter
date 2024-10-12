#!make

binary_name = convert
main_path = ./cmd/convert/main.go

.PHONY: tidy
tidy:
	go mod tidy
	go fmt ./...

.PHONY: build
build:
	go build ${main_path} -o=${binary_name}

.PHONY: run
run: build
	 ./tmp/${binary_name}

