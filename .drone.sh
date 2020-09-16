#!/bin/sh

set -eux

# only execute this script as part of the pipeline.
[ -z "$CI" ] && echo "missing ci environment variable" && exit 2

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-X github.com/drone/drone/version.VersionDev=build.'${DRONE_BUILD_NUMBER} -o release/drone-server github.com/drone/drone/cmd/drone-server
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-X github.com/drone/drone/version.VersionDev=build.'${DRONE_BUILD_NUMBER} -o release/drone-agent github.com/drone/drone/cmd/drone-agent