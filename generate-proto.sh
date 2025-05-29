#!/usr/bin/env bash
set -euo pipefail

echo "🔄 buf build"
buf build \
  --as-file-descriptor-set \
  --output gen/protoset-full.bin

echo "🔄 buf generate"

# 1) datablock 서비스만
buf generate . \
  --template buf.gen.yaml \
  --path protos/datablock/ichthys/v1/datablock_service.proto

# 2) syncfolders 서비스만
buf generate . \
  --template buf.gen.yaml \
  --path protos/syncfolders/ichthys/v1/syncfolders_service.proto

echo "✅ 완료!"
