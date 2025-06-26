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
  --template protos/datablock/ichthys/v1/buf.gen.datablock.yaml \
  --path protos/datablock/ichthys/v1/datablock_service.proto

# 2) syncfolders 서비스만
buf generate . \
  --template protos/syncfolders/ichthys/v1/buf.gen.syncfolders.yaml \
  --path protos/syncfolders/ichthys/v1/syncfolders_service.proto

# 3) volres 서비스만
buf generate . \
  --template protos/volres/ichthys/v1/buf.gen.volres.yaml \
  --path protos/volres/ichthys/v1/volres_service.proto

# 4) tool 서비스만
  buf generate . \
    --template protos/tool/ichthys/v1/buf.gen.tool.yaml \
    --path protos/tool/ichthys/v1/tool_service.proto

echo "✅ 완료!"
