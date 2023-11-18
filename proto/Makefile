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
BUF_BASE = prototpl
buf_base=$(BUF_BASE)
BUF_PATHS = --path $(buf_base)
buf_paths=$(BUF_PATHS)

#yapi
YAPI_TOKEN_DEFAULT = 526b3b7f2926a35c71552f23d64651f5cf6f57bcb208c8a8be753f95173760a4
YAPI_ADDR_DEFAULT = http://127.0.0.1:40001
yapi_addr=$(YAPI_ADDR_DEFAULT)
yapi_token=$(YAPI_TOKEN_DEFAULT)

#commit
RM_COMMIT_LEFT = proto
rm_commit_left = $(RM_COMMIT_LEFT)
ADD_COMMIT_LEFT = .
add_commit_left = $(ADD_COMMIT_LEFT)

.PHONY: help

help:
	@echo "Usage: make [target] [buf_base=] [buf_paths=] [rm_commit_left=] [add_commit_left] ..."
	@echo "Options:"
	@echo "  \033[0;31mbuf_base\033[0m: buf base path, eg: make generate buf_base=gen"
	@echo "  \033[0;31mbuf_paths\033[0m: buf paths, eg: make generate buf_paths=--path gen"
	@echo "  \033[0;31mrm_commit_left\033[0m: rm commit left, eg: make commit-swagger rm_commit_left=proto"
	@echo "  \033[0;31madd_commit_left\033[0m: add commit left, eg: make commit-swagger add_commit_left=."
	@echo "Targets:"
	@echo "  \033[0;31mgenerate\033[0m: generate proto buf and push code, eg: make generate [buf_base=] [buf_paths=]"
	@echo "  \033[0;31mswagger\033[0m: generate swagger file by proto file"
	@echo "  \033[0;31mdel-swagger\033[0m: delete generate swagger file"
	@echo "  \033[0;31myapi\033[0m: push swagger file to yapi, eg: make yapi [yapi_addr=] [yapi_token=]"
	@echo "  \033[0;31mlocal-swagger\033[0m: generate swagger file by proto file, with local proto file change"
	@echo "  \033[0;31mlocal-yapi\033[0m: push swagger file to yapi, eg: make local-yapi [yapi_addr=] [yapi_token=]"
	@echo "  \033[0;31mcommit-swagger\033[0m: generate swagger file by proto file, with latest commit proto file"
	@echo "  \033[0;31mcommit-yapi\033[0m: push swagger file to yapi, eg: make commit-yapi [yapi_addr=] [yapi_token=]"
	@echo "  \033[0;31mtree-proto\033[0m: tree proto file"
	@echo "  \033[0;31mtree-dir\033[0m: tree dir"
	@echo "  \033[0;31mgit-quick-merge\033[0m: quick merge current branch to target branch"
	@echo "  \033[0;31mgit-drop-change\033[0m: drop local change for git"
	@echo "  \033[0;31mgit-sync-master\033[0m: sync current branch to master branch"


# code api
.PHONY: generate swagger del-swagger local-swagger yapi local-yapi commit-swagger commit-yapi tree-proto tree-dir

# generate proto buf and push code
#buf generate --template buf/buf.gen.yaml $(buf_paths);
generate:
		buf generate --template buf/buf.gen.go.yaml $(buf_paths);
#		buf generate --template buf/buf.gen.java.yaml $(buf_paths);
#		buf generate --template buf/buf.gen.cpp.yaml $(buf_paths);
#		buf generate --template buf/buf.gen.rust.yaml $(buf_paths);
#		buf generate --template buf/buf.gen.python.yaml $(buf_paths);
		git add .;
		git commit -m "Chore: generate code by auto"
		#git push origin main;

# generate swagger file by proto file
swagger:
	@ for file in $(FILES); \
            do \
            protoc -I . --openapiv2_out . $$file; \
            done

# delete generate swagger file
del-swagger:
	@ for file in $(foreach dir, $(DIRS), $(wildcard $(dir)/*.swagger.json)); \
			do \
			$(RM) $$file; \
			done

# push swagger file to yapi
yapi:
	$(shell make swagger)
	@ for file in $(foreach dir, $(DIRS), $(wildcard $(dir)/*.swagger.json)); \
            do \
            CONTENT=$$(cat $$file); \
            curl -d "type=swagger&merge=merge&token=$(yapi_token)&json=$$CONTENT" -H "Content-Type: application/x-www-form-urlencoded" -X POST $(yapi_addr)/api/open/import_data; \
            $(RM) $$file; \
            done

# generate swagger file by proto file, with local proto file change
local-swagger:
	@ for file in $(COMMIT_LOCAL_PROTO_FILES); \
			do \
			protoc -I . --openapiv2_out . $$file; \
			done

# push swagger file to yapi
local-yapi:
	$(shell make local-swagger)
	@ for file in $(foreach dir, $(DIRS), $(wildcard $(dir)/*.swagger.json)); \
            do \
            CONTENT=$$(cat $$file); \
            curl -d "type=swagger&merge=merge&token=$(yapi_token)&json=$$CONTENT" -H "Content-Type: application/x-www-form-urlencoded" -X POST $(yapi_addr)/api/open/import_data; \
            $(RM) $$file; \
            done

# generate swagger file by proto file, with latest commit proto file
commit-swagger:
	@ for file in $(COMMIT_PROTO_FILES); \
			do \
			protoc -I . --openapiv2_out . $(add_commit_left)$${file#$(rm_commit_left)}; \
			done

# push swagger file to yapi
commit-yapi:
	$(shell make commit-swagger)
	@ for file in $(foreach dir, $(DIRS), $(wildcard $(dir)/*.swagger.json)); \
            do \
            CONTENT=$$(cat $$file); \
            curl -d "type=swagger&merge=merge&token=$(yapi_token)&json=$$CONTENT" -H "Content-Type: application/x-www-form-urlencoded" -X POST $(yapi_addr)/api/open/import_data; \
            $(RM) $$file; \
            done

# demo
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

# git
.PHONY: git-quick-merge git-drop-change git-sync-master

MAIN_BRANCH = main
TARGET_BRANCH = $(target-branch)

# quick merge current branch to target branch
# eg. make git-quick-merge target-branch=dev
git-quick-merge:
ifneq ($(TARGET_BRANCH), )
	@echo $(TARGET_BRANCH) PASSED
	- git branch -D $(TARGET_BRANCH);
	git fetch;
	export branch=`git branch | grep \* | grep -Eo ' .+'` && \
		echo "current branch: $$branch" && \
		git checkout $(TARGET_BRANCH) && \
		git pull --rebase && \
		git merge origin/$(MAIN_BRANCH) && \
		echo "merge: \033[0;31morigin/$(MAIN_BRANCH)\033[0m  to \033[0;31m$(TARGET_BRANCH)\033[0m" && \
		git merge $$branch && \
		echo "merge: \033[0;31m$$branch\033[0m to \033[0;31m$(TARGET_BRANCH)\033[0m" && \
		git push && \
		git checkout $$branch;
else
	@echo "\033[0;31mFAILED\033[0m, eg. make quick-merge target-branch=dev"
	$(eval $(exit -1))
endif

# drop local change for git
git-drop-change:
	git add .; git stash; git stash drop

# sync current branch to master branch
# err(failed to push some refs to): git pull --rebase origin $$branch
git-sync-master:
	export branch=`git branch | grep \* | grep -Eo ' .+'` && \
		git checkout $(MAIN_BRANCH) && \
		git pull --rebase && \
		git checkout $$branch && \
		git rebase $(MAIN_BRANCH);

print:
#ifneq ($(yapi_addr), )
#	export addr=$(yapi_addr) && \
#	echo "yapi_addr: $$addr"
#endif
#ifneq ($(yapi_token), )
#	export token=$(yapi_token) && \
#	echo "yapi_token: $$token"
#endif
	echo "$(yapi_addr):$(yapi_token)"
