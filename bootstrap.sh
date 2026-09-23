#!/usr/bin/env bash
set -Eeuo pipefail
go test ./...
go vet ./...
mkdir -p dist
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o dist/clone ./cmd/clone
./scripts/smoke.sh
