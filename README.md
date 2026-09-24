# Product Cloning Without Forks

A Go toolchain for generating an independent runnable product from an explicit behavior/interface blueprint instead of copying or forking an upstream source tree.

The input describes behavior. The generator emits its own product.

## Implemented

A blueprint can define:

- pages
- visible blocks
- navigation targets
- API contracts
- HTTP status codes
- response bodies
- initial state

The generator emits a standalone Go product with its own:

- source
- manifest
- UI
- API behavior
- README
- watermark

The build does not require an upstream Git repository or source checkout.

## Build and verify

```bash
make all
```

`make all` runs:

```text
go test ./...
go vet ./...
CGO_ENABLED=0 go build ...
./scripts/smoke.sh
```

The resulting CLI is:

```text
dist/clone
```

## Validate a blueprint

```bash
./dist/clone validate -in examples/board-lite.json
```

## Generate a product

```bash
./dist/clone generate \
  -in examples/board-lite.json \
  -out dist/generated
```

Then run the generated product:

```bash
cd dist/generated
go run .
```

## Development commands

```bash
make test
make build
make smoke
make run
```

## Repository layout

```text
cmd/          CLI entrypoint
internal/     generator/runtime implementation
contracts/    blueprint contracts
examples/     example behavior blueprints
scripts/      smoke verification
dist/         build output
bootstrap.sh  one-shot verification/build path
```

## Scope

This repository reconstructs behavior and interfaces from explicit blueprint data. It does not depend on or require copying an upstream implementation.

## Ownership

Owner: Julius Cameron Hill / Titan Universal AI LLC  
Watermark: `":"`
