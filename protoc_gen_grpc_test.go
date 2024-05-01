package protoc_gen_grpc

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuf(t *testing.T) {
	if err := os.RemoveAll(filepath.Join("out", "buf")); err != nil {
		t.Fatalf("failed to remove build directory: %v", err)
	}

	output := bytes.Buffer{}
	cmd := exec.Command("go", "run", "github.com/bufbuild/buf/cmd/buf@v1.30.0", "generate")
	cmd.Stderr = &output
	cmd.Stdout = &output
	cmd.Dir = "testdata"
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to run buf: %v\n%v", err, output.String())
	}

	for _, path := range []string{
		filepath.Join("out", "buf", "cpp", "helloworld.grpc.pb.cc"),
		filepath.Join("out", "buf", "csharp", "HelloworldGrpc.cs"),
		filepath.Join("out", "buf", "node", "helloworld_grpc_pb.js"),
		filepath.Join("out", "buf", "objective_c", "Helloworld.pbrpc.m"),
		filepath.Join("out", "buf", "php", "Helloworld", "GreeterClient.php"),
		filepath.Join("out", "buf", "python", "helloworld_pb2_grpc.py"),
		filepath.Join("out", "buf", "ruby", "helloworld_services_pb.rb"),
	} {
		if _, err := os.Stat(path); err != nil {
			t.Errorf("failed to stat %v: %v", path, err)
		}
	}
}

func TestProtoc(t *testing.T) {
	if _, err := exec.LookPath("protoc"); err != nil {
		t.Skip("protoc not found")
	}

	outDir := filepath.Join("out", "protoc")
	if err := os.RemoveAll(outDir); err != nil {
		t.Fatalf("failed to remove build directory: %v", err)
	}
	if err := os.RemoveAll(filepath.Join("out", "plugins")); err != nil {
		t.Fatalf("failed to remove build directory: %v", err)
	}

	plugins := []string{"cpp", "csharp", "node", "objective_c", "php", "python", "ruby"}
	for _, plugin := range plugins {
		output := bytes.Buffer{}
		cmd := exec.Command("go", "build", "-o", filepath.Join("build", "plugins", "protoc-gen-grpc_"+plugin), "./cmd/protoc-gen-grpc_"+plugin)
		cmd.Stderr = &output
		cmd.Stdout = &output
		if err := cmd.Run(); err != nil {
			t.Fatalf("failed to build plugin %v: %v\n%v", plugin, err, output.String())
		}

		if err := os.MkdirAll(filepath.Join(outDir, plugin), 0o755); err != nil {
			t.Fatalf("failed to create directory %v: %v", filepath.Join(outDir, plugin), err)
		}
	}
	output := bytes.Buffer{}
	env := os.Environ()
	for i, val := range env {
		if strings.HasPrefix(val, "PATH=") {
			env[i] = "PATH=" + filepath.Join("build", "plugins") + string(os.PathListSeparator) + val[len("PATH="):]
		}
	}
	cmd := exec.Command(
		"protoc",
		"--grpc_cpp_out="+filepath.Join(outDir, "cpp"),
		"--grpc_csharp_out="+filepath.Join(outDir, "csharp"),
		"--grpc_node_out="+filepath.Join(outDir, "node"),
		"--grpc_objective_c_out="+filepath.Join(outDir, "objective_c"),
		"--grpc_php_out="+filepath.Join(outDir, "php"),
		"--grpc_python_out="+filepath.Join(outDir, "python"),
		"--grpc_ruby_out="+filepath.Join(outDir, "ruby"),
		"-I"+filepath.Join("testdata", "protos"),
		filepath.Join("testdata", "protos", "helloworld.proto"),
	)
	cmd.Stderr = &output
	cmd.Stdout = &output
	cmd.Env = env
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to run protoc: %v\n%v", err, output.String())
	}

	for _, path := range []string{
		filepath.Join("out", "protoc", "cpp", "helloworld.grpc.pb.cc"),
		filepath.Join("out", "protoc", "csharp", "HelloworldGrpc.cs"),
		filepath.Join("out", "protoc", "node", "helloworld_grpc_pb.js"),
		filepath.Join("out", "protoc", "objective_c", "Helloworld.pbrpc.m"),
		filepath.Join("out", "protoc", "php", "Helloworld", "GreeterClient.php"),
		filepath.Join("out", "protoc", "python", "helloworld_pb2_grpc.py"),
		filepath.Join("out", "protoc", "ruby", "helloworld_services_pb.rb"),
	} {
		if _, err := os.Stat(path); err != nil {
			t.Errorf("failed to stat %v: %v", path, err)
		}
	}
}
