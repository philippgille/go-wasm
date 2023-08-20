package main

import (
	"context"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func main() {
	ctx := context.Background()

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	buf, err := os.ReadFile("go-wasi.wasm")
	if err != nil {
		panic(err)
	}

	// By default I/O streams are discarded so we have to enable them.
	// Same would be for file system access, which would require a `.WithFS()` with a Go stdlib `io/fs.FS`.
	config := wazero.NewModuleConfig().
		WithStdout(os.Stdout).WithStderr(os.Stderr)

	// Instantiate WASI, which implements system I/O such as console output.
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	// Instantiate() returns the module, whose exported functions we could later call.
	// But it also runs the `_start` function directly, which a Go program's `main` function is compiled to.
	// In our case the Go program only has the `main` function, so we don't need the module.
	_, err = r.InstantiateWithConfig(ctx, buf, config.WithArgs("wasi"))
	if err != nil {
		panic(err)
	}
}
