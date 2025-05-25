PROTOC_PATH:=/usr/local/bin/protoc
PLUGIN_GO:=/usr/local/bin/protoc-gen-go
PLUGIN_GRPC:=/usr/local/bin/protoc-gen-go-grpc

# 컨테이너 내 실제 위치
PROTO_PATH:=./protos
OUT_DIR:=./gen/go

#PROTO_FILES:=$(wildcard $(PROTO_PATH)/*.proto)
PROTO_FILES := $(shell find $(PROTO_PATH) -name '*.proto')

.PHONY: all generate clean
all: generate

generate: $(PROTO_FILES)
	@mkdir -p $(OUT_DIR)
	@for file in $^; do \
	  $(PROTOC_PATH) --proto_path=$(PROTO_PATH) \
	    --plugin=$(PLUGIN_GO) \
	    --go_out=$(OUT_DIR) \
	    --go_opt=paths=source_relative $$file; \
	  $(PROTOC_PATH) --proto_path=$(PROTO_PATH) \
	    --plugin=$(PLUGIN_GRPC) \
	    --go-grpc_out=$(OUT_DIR) \
	    --go-grpc_opt=paths=source_relative $$file; \
	done

clean:
	rm -f $(OUT_DIR)/*.pb.go