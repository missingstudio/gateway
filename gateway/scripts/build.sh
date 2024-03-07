#!/bin/sh

BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT=$(git rev-parse HEAD)
VERSION=$(git describe --tags --abbrev=0 | tr -d '\n')

if [ ! -z "${GOOS}" ]; then
  CGO_ENABLED=0 GOOS="${GOOS}" go build -ldflags="-s -w -X 'github.com/missingstudio/ai/gateway/internal/version.gitVersion=$VERSION' -X 'github.com/missingstudio/ai/gateway/internal/version.buildDate=$BUILD_DATE' -X 'github.com/missingstudio/ai/gateway/internal/version.gitCommit=$GIT_COMMIT'" -o bin/gateway main.go
else
  CGO_ENABLED=0 go build -ldflags="-s -w -X 'github.com/missingstudio/ai/gateway/internal/version.gitVersion=$VERSION' -X 'github.com/missingstudio/ai/gateway/internal/version.buildDate=$BUILD_DATE' -X 'github.com/missingstudio/ai/gateway/internal/version.gitCommit=$GIT_COMMIT'" -o bin/gateway main.go
fi
