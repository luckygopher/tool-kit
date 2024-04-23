PACKAGE=tool
PREFIX=$(shell pwd)
OUTPUT_DIR=${PREFIX}/bin

generate:
	@echo "+ $@"
	go generate ./...

build_linux:
	@echo "+ build"
	go build -o ${OUTPUT_DIR}/${PACKAGE}

build_windows:
	@echo "+ build"
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${OUTPUT_DIR}/${PACKAGE}.exe

build_mac: generate
	@echo "+ build"
	go build -o ${OUTPUT_DIR}/${PACKAGE}