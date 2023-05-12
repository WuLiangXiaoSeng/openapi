#!/bin/bash
set -e
set -x

export CONFIGURED_ARCH=arm64
export ARCH_LG=aarch64-linux-gnu
export BLDENV=buster
export CROSS_COMPILE="/usr/bin/aarch64-linux-gnu-"

mkdir -p ./target/${CONFIGURED_ARCH}

#### GO settings ####
if [ ${CONFIGURED_ARCH} == 'arm64' ]; then
  export GOARCH=${CONFIGURED_ARCH}
  export CC=${CROSS_COMPILE}gcc
fi
export GOPROXY="https://goproxy.io"

##### make httpserver
GODEBUG=madvdontneed=1 go build -o httpserver equantum.com/httpserver

${CROSS_COMPILE}strip httpserver || true
cp httpserver ./target/${CONFIGURED_ARCH}
