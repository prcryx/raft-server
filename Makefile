include .env
export

BINARY_NAME := raft-server
.DEFAULT_GOAL := run
OS := $(OS_NAME)
ARCH := $(ARCH_NAME)

.PHONY: all build clean

release:
	@echo "building..."
	@GOARCH=$(ARCH) GOOS=$(OS) go build -o build/$(BINARY_NAME)-$(OS) cmd/app/main.go

gen:
	@echo "generating wire_gen.go..."
	@wire ./...
dry:
	@echo "building..."
	@go run cmd/app/main.go

clean:
	@go clean
	@rm -rf ./build/*

run: build
	@echo "\nrunning..."
	@./build/$(BINARY_NAME)-$(OS)

dev:
	@echo "runnung..."
	@GOARCH=$(ARCH) GOOS=$(OS) go run cmd/app/main.go

print:
	@echo $(OS)