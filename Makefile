.PHONY: protoc build deps

PROTOS = $(patsubst ./%,%,$(shell find . -name "*.proto"))

build:
	go build -o v2ray ./main

deps: $(GOPATH)/bin/protoc-gen-go $(GOPATH)/bin/protoc-gen-go-grpc
	go mod download

protoc:
	@ protoc --go_out=. --go-grpc_out=. $(PROTOS)

$(GOPATH)/bin/protoc-gen-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.1

$(GOPATH)/bin/protoc-gen-go-grpc:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
