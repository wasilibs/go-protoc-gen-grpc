package wasm

import _ "embed"

//go:embed grpc_cpp_plugin.wasm
var GRPCCPPPlugin []byte

//go:embed grpc_csharp_plugin.wasm
var GRPCCSharpPlugin []byte

//go:embed grpc_node_plugin.wasm
var GRPCNodePlugin []byte

//go:embed grpc_objective_c_plugin.wasm
var GRPCObjectiveCPlugin []byte

//go:embed grpc_php_plugin.wasm
var GRPCPHPPlugin []byte

//go:embed grpc_python_plugin.wasm
var GRPCPythonPlugin []byte

//go:embed grpc_ruby_plugin.wasm
var GRPCRubyPlugin []byte

//go:embed memory.wasm
var Memory []byte
