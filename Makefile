export PROJ_PATH=github.com/alexj212/golang_marshalling

export DATE := $(shell date +%Y.%m.%d-%H%M)
export LATEST_COMMIT := $(shell git log --pretty=format:'%h' -n 1)
export BRANCH := $(shell git branch |grep -v "no branch"| grep \*|cut -d ' ' -f2)
export BUILT_ON_IP := $(shell [ $$(uname) = Linux ] && hostname -i || hostname )
export BIN_DIR=./bin


export BUILT_ON_OS=$(shell uname -a)
ifeq ($(BRANCH),)
BRANCH := master
endif

export COMMIT_CNT := $(shell git rev-list HEAD | wc -l | sed 's/ //g' )
export BUILD_NUMBER := ${BRANCH}-${COMMIT_CNT}



build_info: ## Display environment variables
	@echo ''
	@echo '---------------------------------------------------------'
	@echo 'BUILT_ON_IP       $(BUILT_ON_IP)'
	@echo 'BUILT_ON_OS       $(BUILT_ON_OS)'
	@echo 'DATE              $(DATE)'
	@echo 'LATEST_COMMIT     $(LATEST_COMMIT)'
	@echo 'BRANCH            $(BRANCH)'
	@echo 'COMMIT_CNT        $(COMMIT_CNT)'
	@echo 'BUILD_NUMBER      $(BUILD_NUMBER)'
	@echo 'PATH              $(PATH)'
	@echo 'PACKR2_EXECUTABLE $(PACKR2_EXECUTABLE)'
	@echo '---------------------------------------------------------'
	@echo ''


####################################################################################################################
##
## help for each task - https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
##
####################################################################################################################
.PHONY: help

##help: ## This help.
##	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
help: ## This help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help



####################################################################################################################
##
## Build of binaries
##
####################################################################################################################

create_dir:
	@mkdir -p $(BIN_DIR)

build_app: create_dir
	go build -o $(BIN_DIR)/$(BIN_NAME) -a $(APP_PATH)


####################################################################################################################
##
## Cleanup of binaries
##
####################################################################################################################

clean:  ## clean all binaries in bin dir


clean_binary: ## clean binary in bin dir
	rm -f $(BIN_DIR)/$(BIN_NAME)


test: ## run tests
	go test -count=1 $(PROJ_PATH)/...

fmt: ## run fmt
	go fmt $(PROJ_PATH)/...


all: gob_example xml_example json_example protobuf_example ## build all examples into ./bin dir.


gob_example: build_info ## Build gob_example binary in bin dir.
	@echo "build gob_example"
	make BIN_NAME=gob_example       APP_PATH=$(PROJ_PATH)/example1 build_app
	@echo ''
	@echo ''

xml_example: build_info ## Build xml_example binary in bin dir
	@echo "build xml_example"
	make BIN_NAME=xml_example       APP_PATH=$(PROJ_PATH)/example2 build_app
	@echo ''
	@echo ''

json_example: build_info ## Build json_example binary in bin dir
	@echo "build json_example"
	make BIN_NAME=json_example       APP_PATH=$(PROJ_PATH)/example3 build_app
	@echo ''
	@echo ''


protobuf_example: build_info## Build protobuf_example binary in bin dir
	@echo "build protobuf_example"
	@protoc -I./example4   --go_out=./example4    ./example4/data.proto
	@protoc -I./example4   --java_out=./example4  ./example4/data.proto
	make BIN_NAME=protobuf_example       APP_PATH=$(PROJ_PATH)/example4 build_app
	@echo ''
	@echo ''

