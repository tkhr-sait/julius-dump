#!/bin/bash

mkdir -p target
for OS in darwin linux windows ; do
  GOOS=${OS} go build -o target/julius-dump.${OS}
done
GOOS=linux GOARCH=arm go build -o target/julius-dump.linux.arm
