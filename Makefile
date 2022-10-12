.PHONY: default build clean  prof bench run

export GO111MODULE=on

MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR := $(dir $(MKFILE_PATH))
DEMO_DIR := $(MKFILE_DIR)/demo


GOBUILD := ${DEMO_DIR}/build
GOBIN := ${GOBUILD}/bin
export GOBIN


TARGET=${DEMO_DIR}/build/bin/main
SOURCE=${DEMO_DIR}/main/main.go
ALL_SOURCE=$(shell find ${MKFILE_DIR} -type f -name "*.go")

default: ${TARGET}


${TARGET}: ${ALL_SOURCE}
	@echo "-------------- building ---------------"
	go mod tidy
	mkdir -p ${DEMO_DIR}build/bin/
	cd ${DEMO_DIR} && go build -v -ldflags "-s -w" -o ${TARGET} ${SOURCE}
	mkdir -p ${GOBUILD}/data/ && cp ${DEMO_DIR}/data/* ${GOBUILD}/data/

build: default

clean:
	@rm -rf ${TARGET} &&  rm -rf ${GOBUILD}/data/ && rm -rf ${GOBUILD}

run : ${TARGET}
	${TARGET}

prof: ${TARGET}
	${TARGET} -cpuprofile=downsampling.prof
	go tool pprof ${TARGET} downsampling.prof

bench:
	go test -bench=. ./core/


