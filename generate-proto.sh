#!/usr/bin/env bash
set -euo pipefail

echo "🔄 buf build"
buf build \
  --as-file-descriptor-set \
  --output gen/protoset-full.bin

echo "🔄 buf generate"

# 서비스가 생길때마다 여기에 추가 해야함
# 1) datablock 서비스만
buf generate . \
  --template buf.gen.yaml \
  --path protos/datablock/ichthys/v1/datablock_service.proto

# 2) syncfolders 서비스만
buf generate . \
  --template buf.gen.yaml \
  --path protos/syncfolders/ichthys/v1/syncfolders_service.proto

echo "✅ 완료!"
