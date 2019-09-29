#!/usr/bin/env bash
set -e
set -u
set -o pipefail
set -x

go test -v -race ./...
