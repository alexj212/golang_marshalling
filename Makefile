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



build_info: ## Build the container
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

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

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




example1: build_info ## build example1 binaries in bin dir
	@echo "build example1"
	make BIN_NAME=example1       APP_PATH=$(PROJ_PATH)/example1 build_app
	@echo ''
	@echo ''

example2: build_info ## build example2 binaries in bin dir
	@echo "build example2"
	make BIN_NAME=example2       APP_PATH=$(PROJ_PATH)/example2 build_app
	@echo ''
	@echo ''

example3: build_info ## build example3 binaries in bin dir
	@echo "build example3"
	make BIN_NAME=example3       APP_PATH=$(PROJ_PATH)/example3 build_app
	@echo ''
	@echo ''


example4: build_info ## build example4 binaries in bin dir
	@echo "build example4"
	@protoc -I./example4   --go_out=./example4    ./example4/data.proto
	@protoc -I./example4   --java_out=./example4  ./example4/data.proto
	make BIN_NAME=example4       APP_PATH=$(PROJ_PATH)/example4 build_app
	@echo ''
	@echo ''


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


all: example1 example2 example3## build all
