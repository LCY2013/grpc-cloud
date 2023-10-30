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
BUF_BASE = gen
BUF_PATHS = --path $(BUF_BASE)

#yapi
YAPI_TOKEN = 526b3b7f2926a35c71552f23d64651f5cf6f57bcb208c8a8be753f95173760a4
YAPI_ADDR = http://127.0.0.1:40001
yapi_addr=$(YAPI_ADDR)
yapi_token=$(YAPI_TOKEN)

#commit
RM_COMMIT_LEFT = proto
ADD_COMMIT_LEFT = .

.PHONY: help

help:
	@echo "Usage: make [target] ..."
	@echo "Targets:"
	@echo "  generate: generate proto buf and push code"
	@echo "  swagger: generate swagger file by proto file"
	@echo "  del-swagger: delete generate swagger file"
	@echo "  yapi: push swagger file to yapi"
	@echo "  local-swagger: generate swagger file by proto file, with local proto file change"
	@echo "  local-yapi: push swagger file to yapi"
	@echo "  commit-swagger: generate swagger file by proto file, with latest commit proto file"
	@echo "  commit-yapi: push swagger file to yapi"
	@echo "  tree-proto: demo"
	@echo "  tree-dir: demo"
	@echo "  git-quick-merge: quick merge current branch to target branch"
	@echo "  git-drop-change: drop local change for git"
	@echo "  git-sync-master: sync current branch to master branch"


# code api
.PHONY: generate swagger del-swagger local-swagger yapi local-yapi commit-swagger commit-yapi tree-proto tree-dir

# generate proto buf and push code
generate:
		buf generate --template buf/buf.gen.yaml $(BUF_PATHS);
		cd $(BUF_BASE);
		git add .;
		git commit -m "Chore: generate code by auto"
		git push origin main;

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
			protoc -I . --openapiv2_out . $(ADD_COMMIT_LEFT)$${file#$(RM_COMMIT_LEFT)}; \
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
