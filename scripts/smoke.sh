#!/usr/bin/env bash
set -Eeuo pipefail
rm -rf dist/generated
./dist/clone validate -in examples/board-lite.json
./dist/clone generate -in examples/board-lite.json -out dist/generated
test -s dist/generated/main.go
( cd dist/generated && go test ./... && CGO_ENABLED=0 go build -trimpath -o ../board-lite . )
test -x dist/board-lite
echo "SMOKE PASS generated product compiled"
