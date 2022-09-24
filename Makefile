PACKAGE=tool
PREFIX=$(shell pwd)
OUTPUT_DIR=${PREFIX}/bin

generate:
	@echo "+ $@"
	go generate ./...

build: generate
	@echo "+ build"
	go build -o ${OUTPUT_DIR}/${PACKAGE}