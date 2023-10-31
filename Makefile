bin := raft-server
.DEFAULT_GOAL := run
os := windows
arch := amd64

.PHONY: all build clean

release:
	@echo "building..."
	@GOARCH=$(arch) GOOS=$(os) go build -o build/$(bin)-$(os) cmd/app/main.go

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
	@./build/$(bin)-$(os)

dev:
	@echo "runnung..."
	@GOARCH=$(arch) GOOS=$(os) go run cmd/app/main.go

print:
	@echo $(os)