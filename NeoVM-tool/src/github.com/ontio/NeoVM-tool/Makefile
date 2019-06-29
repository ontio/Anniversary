GC=go build
VERSION := $(shell git describe --always --tags --long)

ARCH=$(shell uname -m)
SRC_FILES = $(shell git ls-files | grep -e .go$ | grep -v _test.go)

nvm-tool: $(SRC_FILES)
	$(GC)  $(BUILD_NODE_PAR) -o nvm-tool main.go
 
all: nvm-tool-windows nvm-tool-linux nvm-tool-darwin

nvm-tool-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o nvm-tool-windows-amd64.exe main.go

nvm-tool-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o nvm-tool-linux-amd64 main.go

nvm-tool-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o nvm-tool-darwin-amd64 main.go

