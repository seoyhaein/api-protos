# .devcontainer/Dockerfile
FROM mcr.microsoft.com/devcontainers/base:ubuntu-22.04

# 1) 공통 의존성 설치
RUN apt-get update && apt-get install -y --no-install-recommends \
      wget \
      build-essential \
      unzip \
    && rm -rf /var/lib/apt/lists/*

# 2) Buf CLI 설치
RUN wget -O /usr/local/bin/buf \
      https://github.com/bufbuild/buf/releases/download/v1.54.0/buf-Linux-x86_64 \
    && chmod +x /usr/local/bin/buf

# 3) Go 1.22.4 설치
RUN wget https://go.dev/dl/go1.22.4.linux-amd64.tar.gz \
      -O go.tar.gz \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm go.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# 4) protoc (v26.1) 설치
RUN wget -O protoc.zip \
      https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-linux-x86_64.zip \
    && unzip protoc.zip -d protoc \
    && mv protoc/bin/protoc /usr/local/bin/ \
    && mv protoc/include /usr/local/ \
    && rm -rf protoc protoc.zip

# 5) protoc-gen-go (v1.33.0) 설치
RUN wget -O go-plugin.tar.gz \
      https://github.com/protocolbuffers/protobuf-go/releases/download/v1.33.0/protoc-gen-go.v1.33.0.linux.amd64.tar.gz \
    && tar -xzf go-plugin.tar.gz protoc-gen-go \
    && mv protoc-gen-go /usr/local/bin/ \
    && chmod +x /usr/local/bin/protoc-gen-go \
    && rm go-plugin.tar.gz

# 6) protoc-gen-go-grpc 설치
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    && mv /root/go/bin/protoc-gen-go-grpc /usr/local/bin/ \
    && chmod +x /usr/local/bin/protoc-gen-go-grpc

# 7) 작업 디렉터리 & 볼륨
WORKDIR /app
#VOLUME ["/app/protos"]

# 8) 소스 복사 (gen/와 protos/ 디렉터리는 호스트에서 마운트하거나 COPY)
#COPY gen/ ./gen/
#COPY protos/ ./protos/

# 9) entrypoint: 기본은 make generate
#    → Buf 쓰고 싶을 땐 터미널에서 buf generate, buf lint 명령을 직접 실행
#CMD ["make", "generate"]
