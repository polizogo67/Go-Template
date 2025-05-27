VERSION := 1.0.0
BUILD_TIME := $(shell date)
# BUILD_TIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

hello:
	@echo "Hello"
	@echo $(BUILD_TIME) Version: $(VERSION)

build:
	go build -ldflags='-s -w -X main.version=$(VERSION) -X "main.buildTime=$(BUILD_TIME)"' -o bin/main ./cmd/main

run:
	go run ./cmd/main