package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "protoc_gen_grpc",
		LibraryRepo: "grpc/grpc",
		GoReleaser:  true,
	})
	boot.Main()
}
