import * as _ from "./wasm_exec.js";

const go = new window.Go();

const buf = await Deno.readFile("./go-wasm.wasm")
const inst = await WebAssembly.instantiate(buf, go.importObject);

go.run(inst.instance);
