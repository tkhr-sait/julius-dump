#!/bin/bash

mkdir -p target
ARCH=amd64
for OS in darwin linux windows ; do
  SUFFIX=""
	if [ "${OS}" = "windows" ]; then
    SUFFIX=".exe"
	fi
  GOOS=${OS} go build -o target/julius-dump.${OS}.${ARCH}${SUFFIX}
done

OS=linux
ARCH=arm64
GOOS=${OS} GOARCH=${ARCH} go build -o target/julius-dump.${OS}.${ARCH}
