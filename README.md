# go-wasm

This repo contains examples of how to work with WebAssembly and WASI in the Go ecosystem.

Regular WASM is meant to be executed inside a browser (like Chrome or Firefox) or JavaScript runtime (like Node, Deno or Bun). As such it only has access to [web APIs](https://developer.mozilla.org/en-US/docs/Web/API) or runtime-specific APIs.

WASI on the other hand, as the name "**W**ebAssembly **S**ystem **I**nterface" suggests, provides more direct access to the host system, like to the filesystem or network sockets. Similar to how the browser executes JavaScript and WebAssembly in a sandbox, WASI runtimes also execute WASI programs in a sandbox, requiring explicit permissions for things like file access.

## WASM

This works with Go versions *prior* to 1.21, as it's "regular" WASM, not WASI.

Compile the Go program to WebAssembly: `GOOS=js GOARCH=wasm go build -o go-wasm.wasm`

Go WASM execution requires a Go-specific wrapper, `wasm_exec.js`. This can then be executed by any JavaScript runtime, in or outside the browser. The Go project provides this in `$(go env GOROOT)/misc/wasm/wasm_exec.js`.

The non-browser runtimes differ in how files are read from the host's filesystem, so for reading the `wasm_exec.js` file, in addition to this file itself, we need to have another runtime-specific wrapper.

### Browser

For running the WASM program in the browser, you can serve the files with any web server that supports the `application/wasm` MIME type. You can use a regular Go server (`http.ListenAndServe(...)`), Caddy, or others.

```bash
cp go-wasm.wasm browser
cp $(go env GOROOT)/misc/wasm/wasm_exec.js browser/wasm_exec.js
cd browser
caddy file-server --listen :2015
```

Then in your browser visit <http://localhost:2015>.

> Tested with Firefox 116.0.3

### Node

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

> Tested with Node v18.17.1
