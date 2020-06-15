# tla protobuf definitions

## generate gRPC stub (.pb.go)
```
protoc --go_out=plugins=grpc:. story.proto
protoc --go_out=plugins=grpc:. location.proto
protoc --go_out=plugins=grpc:. vote.proto
```
