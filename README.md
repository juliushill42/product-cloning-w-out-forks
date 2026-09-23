# Product Cloning Without Forks

Builds an independent runnable product from a behavior/interface blueprint instead of copying or forking an upstream repository. A blueprint defines pages, visible blocks, navigation targets, API contracts, status codes, response bodies, and initial state. The generator emits a standalone Go product with its own source, manifest, UI, API behavior, README, and watermark.

```bash
make all
./dist/clone validate -in examples/board-lite.json
./dist/clone generate -in examples/board-lite.json -out dist/generated
cd dist/generated && go run .
```

The build never requires an upstream Git repository or source tree. Input is explicit behavior and interface information supplied to the blueprint.

Owner: Julius Cameron Hill / Titan Universal AI LLC  
Watermark: `":"`
