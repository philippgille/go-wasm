# go-wasm

This repo contains examples of how to work with WebAssembly and WASI in the Go ecosystem.

Regular WASM is meant to be executed inside a browser (like Chrome or Firefox) or JavaScript runtime (like [Node](https://nodejs.org), [Deno](https://deno.land/) or [Bun](https://bun.sh/)). As such it only has access to [web APIs](https://developer.mozilla.org/en-US/docs/Web/API) or runtime-specific APIs.

WASI on the other hand, as the name "**W**eb**A**ssembly **S**ystem **I**nterface" suggests, provides more direct access to the host system, like to the filesystem or network sockets. Similar to how the browser executes JavaScript and WebAssembly in a sandbox, WASI runtimes also execute WASI programs in a sandbox, requiring explicit permissions for things like file access.

## Table of Contents

- [WASM](#wasm)
  - [Execute WASM directly](#execute-wasm-directly)
    - [Browser](#browser)
    - [Node](#node)
    - [Deno](#deno)
    - [Bun](#bun)
  - [Embed WASM in Go](#embed-wasm-in-go)
- [WASI](#wasi)
  - [Execute WASI directly](#execute-wasi-directly)
  - [Embed WASI in Go](#embed-wasi-in-go)
    - [wazero](#wazero)
    - [wasmtime](#wasmtime)
    - [Wasmer](#wasmer)

## WASM

This works with Go versions *prior* to 1.21, as it's "regular" WASM, not WASI.

Compile the Go program to WebAssembly: `GOOS=js GOARCH=wasm go build -o go-wasm.wasm`

### Execute WASM directly

Go WASM execution requires a Go-specific wrapper, `wasm_exec.js`. This can then be executed by any JavaScript runtime, in or outside the browser. The Go project provides this in `$(go env GOROOT)/misc/wasm/wasm_exec.js`.

The non-browser runtimes differ in how files are read from the host's filesystem, so for reading the `wasm_exec.js` file, in addition to this file itself, we need to have another runtime-specific wrapper.

#### Browser

For running the WASM program in the browser, you can serve the files with any web server that supports the `application/wasm` MIME type. You can use a regular Go server (`http.ListenAndServe(...)`), [Caddy](https://caddyserver.com/), or others.

```bash
cp go-wasm.wasm browser
cp $(go env GOROOT)/misc/wasm/wasm_exec.js browser/wasm_exec.js
cd browser
caddy file-server --listen :2015
```

Then in your browser visit <http://localhost:2015>.

> Tested with Firefox 116.0.3

#### Node

```bash
cp go-wasm.wasm node
cp $(go env GOROOT)/misc/wasm/wasm_exec.js deno/wasm_exec.js
cd node
node node.js
```

The Go authors also provide a convenience script for this, so instead of copying the `wasm_exec.js` and writing our own `node.js` wrapper script, we can simply run:

```bash
$(go env GOROOT)/misc/wasm/go_js_wasm_exec go-wasm.wasm
```

> Tested with [Node v18.17.1](https://github.com/nodejs/node/releases/tag/v18.17.1)

#### Deno

```bash
cp go-wasm.wasm deno
cp $(go env GOROOT)/misc/wasm/wasm_exec.js deno/wasm_exec.js
cd deno
deno run --allow-read=./go-wasm.wasm deno.js
```

ℹ️: Deno runs programs in a sandbox by default, so for reading the WASM file we need to grant the permission accordingly.

> Tested with [Deno 1.36.1](https://github.com/denoland/deno/releases/tag/v1.36.1)

#### Bun

```bash
cp go-wasm.wasm bun
cp $(go env GOROOT)/misc/wasm/wasm_exec.js deno/wasm_exec.js
cd bun
bun bun.js
```

> Tested with [Bun 0.7.3](https://github.com/oven-sh/bun/releases/tag/bun-v0.7.3)

### Embed WASM in Go

🚧 TODO

## WASI

Go 1.21 has official WASI *preview1* support, so we can compile Go programs to WASM that execute outside of a browser / JavaScript runtime.

Compile Go program to WASI: `GOOS=wasip1 GOARCH=wasm go build -o go-wasi.wasm`

### Execute WASI directly

Any WASM runtime with WASI support, on any OS, can then execute this file:

- [wasmtime](https://wasmtime.dev/): `wasmtime go-wasi.wasm` (tested with [wasmtime CLI 11.0.0](https://github.com/bytecodealliance/wasmtime/releases/tag/v11.0.0))
- [Wasmer](https://wasmer.io/): `wasmer run go-wasi.wasm` (tested with [Wasmer 4.1.1](https://github.com/wasmerio/wasmer/releases/tag/v4.1.1))
- [wazero](https://wazero.io/): `wazero run go-wasi.wasm` (tested with [wazero 1.4.0](https://github.com/tetratelabs/wazero/releases/tag/v1.4.0))

### Embed WASI in Go

Or we run the WASI program with an embeddable WASM runtime from within another Go program. Most of the ⬆️ listed runtimes also have Go libraries for embedding.

#### wazero

```bash
cp go-wasi.wasm embed-wasi/wazero
cd embed-wasi/wazero
go run .
```

#### wasmtime

🚧 TODO: <https://github.com/bytecodealliance/wasmtime-go>

#### Wasmer

🚧 TODO: <https://github.com/wasmerio/wasmer-go>
