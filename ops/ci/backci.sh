#!/bin/sh

set -eu -o pipefail

cd backend

echo ""
echo "[BACKEND CI] Installing dependencies"
go get -v ./...
wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.2

echo ""
echo "[BACKEND CI] Running linter"
GOFLAGS=-buildvcs=false golangci-lint run --verbose --show-stats

echo ""
echo "[BACKEND CI] Running tests"
go test -v ./...
