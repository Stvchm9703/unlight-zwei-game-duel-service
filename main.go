package main

import (
	// _ "github.com/gogo/googleapis"
	"ULZGameDuelService/pkg/cmd"

	_ "github.com/gogo/protobuf/proto"
	// _ "github.com/golang/protobuf/protoc-gen-go"
	// _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	// _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/mwitkow/go-proto-validators"
	// _ "github.com/grpc-ecosystem/grpc-gateway" // ???
)

func main() {
	cmd.Execute()
}
