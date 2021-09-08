#!/bin/bash

# you will need to install the thrift compiler from: `brew install thrift`

# thrift --version
# Thrift version 0.12.0

set -e
set -x

rm -rf _gen
mkdir -p _gen/if

D=./if

THRIFT_FILES="
  ${D}/service_v1.thrift
"

for f in $THRIFT_FILES;
do
  echo "=> THRIFT: generate: $f";
  cp -rfv $f _gen/if/
  thrift -r -gen py -o _gen $f;
  thrift -r -gen go:package_prefix=blackbox/gen/ -o _gen $f;
done

echo "=> THRIFT: write to: ./go/src/"
rsync -rc -delete _gen/gen-go/* ./go/src/blackbox/gen/
echo "=> THRIFT: write to: ./py/gen-py/"
rsync -rc -delete _gen/gen-py ./py/

# rm -rf _gen/gen-go
# rm -rf _gen/gen-py
