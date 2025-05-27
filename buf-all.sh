#!/usr/bin/env bash
set -euo pipefail

# buf 있어서 메뉴얼 보면서 다시 작업해야 함.

WORKDIR="$(pwd)"
echo "Working directory: $WORKDIR"

# 1) 워크스페이스 설정 buf.yaml 초기화 (이미 있으면 스킵)
if [ ! -f buf.yaml ]; then
  echo "[1/6] Initializing buf.yaml (module)…"
  buf config init
else
  echo "[1/6] buf.yaml already exists, skipping buf config init."
fi

# 2) 의존성 동기화
echo "[3/6] Updating Buf modules (deps)…"
buf dep update

# 3) 풀 디스크립터 세트 생성 (imports + source info)
OUT_DIR="./gen"
mkdir -p "$OUT_DIR"
echo "[4/6] Generating full descriptor set…"
buf build \
  --include-imports \
  --include-source-info \
  --output "$OUT_DIR/protoset-full.bin"

# 5) (선택) Go/C# 코드 생성
if [ -f buf.gen.yaml ]; then
  echo "[5/6] Generating code with buf.gen.yaml…"
  buf generate --template buf.gen.yaml
else
  echo "[5/6] No buf.gen.yaml found, skipping codegen."
fi

echo "[6/6] All done!"
