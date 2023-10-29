# Makefile ends here

#TOP_PATH = $(shell pwd)
TOP_PATH = .
IGNORE_DIR = ! -path "$(TOP_PATH)/gapic/*" ! -path "$(TOP_PATH)/google/*" ! -path "$(TOP_PATH)/grafeas/*" ! -path "$(TOP_PATH)/grpc-gateway/*"
DIRS = $(shell find $(TOP_PATH) $(IGNORE_DIR) -maxdepth 50 -type d)
FILE_TYPE = *.proto
FILES = $(foreach dir, $(DIRS), $(wildcard $(dir)/$(FILE_TYPE)))
COMMIT_LOCAL_PROTO_FILES = $(shell find $(TOP_PATH) $(IGNORE_DIR) -maxdepth 50 -mtime -1 -type f -name *.proto)
COMMIT_PROTO_FILES = $(shell git diff HEAD^ HEAD --name-only | grep -G .proto)

#order
MV = mv -f
RM = rm -rf
LN = ln -sf

# buf
BUF_PATHS = --path gen

#yapi
YAPI_TOKEN = 526b3b7f2926a35c71552f23d64651f5cf6f57bcb208c8a8be753f95173760a4
YAPI_ADDR = http://127.0.0.1:40001

#commit
RM_COMMIT_LEFT = proto
ADD_COMMIT_LEFT = .

.PHONY: generate swagger del-swagger local-swagger yapi local-yapi commit-swagger commit-yapi

generate:
		buf generate --template buf/buf.gen.yaml $(BUF_PATHS);
		git add $(BUF_PATHS)/*;
		git commit -m "Chore: generate code by auto"

swagger:
	@ for file in $(FILES); \
            do \
            protoc -I . --openapiv2_out . $$file; \
            done

del-swagger:
	@ for file in $(foreach dir, $(DIRS), $(wildcard $(dir)/*.swagger.json)); \
			do \
			$(RM) $$file; \
			done

yapi:
	$(shell make swagger)
	@ for file in $(foreach dir, $(DIRS), $(wildcard $(dir)/*.swagger.json)); \
            do \
            CONTENT=$$(cat $$file); \
            curl -d "type=swagger&merge=merge&token=$(YAPI_TOKEN)&json=$$CONTENT" -H "Content-Type: application/x-www-form-urlencoded" -X POST $(YAPI_ADDR)/api/open/import_data; \
            $(RM) $$file; \
            done

local-swagger:
	@ for file in $(COMMIT_LOCAL_PROTO_FILES); \
			do \
			protoc -I . --openapiv2_out . $$file; \
			done

local-yapi:
	$(shell make local-swagger)
	@ for file in $(foreach dir, $(DIRS), $(wildcard $(dir)/*.swagger.json)); \
            do \
            CONTENT=$$(cat $$file); \
            curl -d "type=swagger&merge=merge&token=$(YAPI_TOKEN)&json=$$CONTENT" -H "Content-Type: application/x-www-form-urlencoded" -X POST $(YAPI_ADDR)/api/open/import_data; \
            $(RM) $$file; \
            done

commit-swagger:
	@ for file in $(COMMIT_PROTO_FILES); \
			do \
			protoc -I . --openapiv2_out . $(ADD_COMMIT_LEFT)$${file#$(RM_COMMIT_LEFT)}; \
			done

commit-yapi:
	$(shell make commit-swagger)
	@ for file in $(foreach dir, $(DIRS), $(wildcard $(dir)/*.swagger.json)); \
            do \
            CONTENT=$$(cat $$file); \
            curl -d "type=swagger&merge=merge&token=$(YAPI_TOKEN)&json=$$CONTENT" -H "Content-Type: application/x-www-form-urlencoded" -X POST $(YAPI_ADDR)/api/open/import_data; \
            $(RM) $$file; \
            done

tree-proto:
	@ for file in $(FILES); \
        do \
        echo $$file; \
        done

tree-dir:
	@ for file in $(FILES); \
        do \
        echo $$file; \
        echo $${file%/*}; \
        done
