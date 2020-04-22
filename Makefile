.PHONY: default build clean depend vget vclean glide

MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR := $(dir $(MKFILE_PATH))

GOPATH := ${MKFILE_DIR}
export GOPATH
GOBUILD := ${GOPATH}build/
GOBIN := ${GOBUILD}bin
export GOBIN

GLIDE_CONF=${MKFILE_DIR}glide.yaml
GLIDE=${GOBIN}/glide
GLIDE_SYS := $(shell command -v glide)
TARGET=${MKFILE_DIR}build/bin/main
SOURCE=${MKFILE_DIR}src/main/main.go
ALL_SOURCE=$(shell find ${MKFILE_DIR}src -type f -name "*.go")

default: ${TARGET}


${TARGET}: ${ALL_SOURCE}
	@echo "-------------- building ---------------"
	mkdir -p ${MKFILE_DIR}build/bin/
	cd ${MKFILE_DIR} && go build -i -v -ldflags "-s -w" -o ${TARGET} ${SOURCE}
	mkdir -p ${GOBUILD}data/ && cp ${MKFILE_DIR}data/* ${GOBUILD}data/

build: default

clean: 
	@rm -rf ${TARGET} &&  rm -rf ${GOBUILD}data/ && rm -rf ${GOBUILD}

${GLIDE}:
	if [ -f ${GLIDE_SYS} ]; then \
		mkdir -p ${GOBIN}; \
		ln -s ${GLIDE_SYS} ${GLIDE}; \
	fi
	if [ ! -f ${GLIDE} ]; then  \
		GOPATH=${GOBUILD}; \
		go get -v github.com/Masterminds/glide; \
		rm -rf ${GOBUILD}pkg ${GOBUILD}src; \
	fi


vget: ${GLIDE}
	if [ ! -f ${GLIDE_CONF} ]; then ${GLIDE} init; fi
	${GLIDE} install
	ln -s ${MKFILE_DIR}vendor ${MKFILE_DIR}src/vendor

vupd: ${GLIDE}
	rm -rf $(HOME)/.glide/cache
	${GLIDE} update && ${GLIDE} install
	#rm -rf ${MKFILE_DIR}src/vendor ${GLIDE} ${MKFILE_DIR}build/pkg ${MKFILE_DIR}build/src
	ln -s ${MKFILE_DIR}vendor ${MKFILE_DIR}src/vendor

vclean:
	rm -f ${MKFILE_DIR}glide.lock
	rm -rf ${MKFILE_DIR}src/vendor ${GLIDE} ${MKFILE_DIR}build/pkg ${MKFILE_DIR}build/src
	rm -dRf ${MKFILE_DIR}src/vendor ${MKFILE_DIR}vendor


