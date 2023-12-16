# go-protoc-gen-grpc

go-protoc-gen-grpc is a distribution of the official gRPC protoc generation plugins from [grpc/grpc][1]
(this is not to be confused with protoc-gen-grpc-go, the compiler for generating Go gRPC stubs, which also
happens to be written in Go). It does not actually reimplement any functionality of gRPC in Go, instead compiling
the original source code to WebAssembly, and executing with the pure Go Wasm runtime [wazero][2].
This means that `go install` or `go run` can be used to execute it, with no need to rely on external
package managers such as Homebrew, on any platform that Go supports.

## Installation

Install the plugin you want using `go install`.

```bash
$ go install github.com/wasilibs/go-protoc-gen-grpc/cmd/protoc-gen-grpc_python@latest
```

As long as `$GOPATH/bin`, e.g. `~/go/bin` is on the `PATH`, you can use it with protoc as normal.

```bash
$ protoc --grpc_python_out=out/python -Iprotos protos/helloworld.proto
```

Note that the filenames of binaries in this repository match protoc conventions so `--plugin` is not needed.

For [buf][3] users, it can be convenient to use `go run` in `buf.gen.yaml`.

```yaml
version: v1
plugins:
  - plugin: grpc_python
    out: out/python
    path: ["go", "run", "github.com/wasilibs/go-protoc-gen-grpc/cmd/protoc-gen-grpc_python@latest"]
```

If also using [go-protoc][4] for `protoc_path` when generating the non-gRPC protobuf stubs, and invoking
`buf` with `go run`, it is possible to have full protobuf/gRPC generation with no installation of tools,
besides Go itself, on any platform that Go supports. The above examples use `@latest`, but it is
recommended to specify a version, in which case all of the developers on your codebase will use the
same version of the tool with no special steps.

[1]: https://github.com/grpc/grpc
[2]: https://wazero.io/
[3]: https://buf.build/
[4]: https://github.com/wasilibs/go-protoc
