#!/usr/bin/env bash
set -euo pipefail

# 이 스크립트는 buf.yaml에 정의된 모듈을 BSR에 푸시(push)합니다.
# 사용 전 반드시 환경 변수 BUF_TOKEN을 설정하세요.

: "${BUF_TOKEN:?Environment variable BUF_TOKEN must be set}"

# 기본 라벨(v1)을 첫 번째 인자로 받거나, 인자가 없으면 v1으로 설정합니다.
LABEL="${1:-v1}"

echo "▶ Pushing workspace modules with label '${LABEL}'..."
# 워크스페이스 루트에서 실행 시 buf.yaml 내 모든 모듈을 대상으로 합니다.
buf push \
  --create \
  --create-visibility public \
  --create-default-label "${LABEL}" \
  --label "${LABEL}"

echo "✅ Push completed."
