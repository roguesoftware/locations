#!/bin/bash

# go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
# go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

# Generate gRPC stub (.pb.go)
# --go-grpc_out=
protoc --go_out=plugins=grpc:. story.proto
protoc --go_out=plugins=grpc:. location.proto
protoc --go_out=plugins=grpc:. vote.proto
