.PHONY: default build clean depend vget vclean glide

MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR := $(dir $(MKFILE_PATH))


GOBUILD := ${MKFILE_DIR}build/
GOBIN := ${GOBUILD}bin
export GOBIN


TARGET=${MKFILE_DIR}build/bin/main
SOURCE=${MKFILE_DIR}/main/main.go
ALL_SOURCE=$(shell find ${MKFILE_DIR} -type f -name "*.go")

default: ${TARGET}


${TARGET}: ${ALL_SOURCE}
	@echo "-------------- building ---------------"
	mkdir -p ${MKFILE_DIR}build/bin/
	cd ${MKFILE_DIR} && go build -v -ldflags "-s -w" -o ${TARGET} ${SOURCE}
	mkdir -p ${GOBUILD}data/ && cp ${MKFILE_DIR}data/* ${GOBUILD}data/

build: default

clean: 
	@rm -rf ${TARGET} &&  rm -rf ${GOBUILD}data/ && rm -rf ${GOBUILD} 

run : ${TARGET}
	${TARGET}

prof: ${TARGET}
	${TARGET} -cpuprofile=downsampling.prof
	go tool pprof ${TARGET} downsampling.prof

bench:
	go test -bench=. ./core/

vget: 
	go get ./...




