package main

import (
	"github.com/wasilibs/go-protoc-gen-grpc/internal/runner"
	"github.com/wasilibs/go-protoc-gen-grpc/internal/wasm"
)

func main() {
	runner.Run("protoc-gen-grpc_node", wasm.GRPCNodePlugin)
}
