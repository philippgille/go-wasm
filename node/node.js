globalThis.crypto = require("crypto"); // required in wasm_exec.js

const fs = require('fs');
require('./wasm_exec.js');

const go = new Go();

const buf = fs.readFileSync('./go-wasm.wasm');
WebAssembly.instantiate(buf, go.importObject).then((inst) => {
    go.run(inst.instance);
});
