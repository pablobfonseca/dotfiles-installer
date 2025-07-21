APP_NAME := dotfiles
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOINSTALL := $(GOCMD) install

BINARY_NAME := $(GOBIN)/$(APP_NAME)
DIST_DIR := $(GOBASE)/dist

# Build flags
LDFLAGS := -ldflags "-X main.version=$(VERSION) -s -w"

.PHONY: all build clean install run release release-all test

all: build

build:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf $(DIST_DIR)

install:
	$(GOINSTALL)

run: build
	$(BINARY_NAME)

test:
	$(GOCMD) test -v ./...

# Cross-compilation targets
release-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(APP_NAME)-darwin-amd64 -v

release-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(APP_NAME)-darwin-arm64 -v

release-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(APP_NAME)-linux-amd64 -v

release-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(APP_NAME)-linux-arm64 -v

# Build all release binaries
release-all: clean
	mkdir -p $(DIST_DIR)
	$(MAKE) release-darwin-amd64
	$(MAKE) release-darwin-arm64
	$(MAKE) release-linux-amd64
	$(MAKE) release-linux-arm64
	@echo "Release binaries built in $(DIST_DIR)/"
	@ls -la $(DIST_DIR)/

# Quick release (current platform)
release: clean
	mkdir -p $(DIST_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(APP_NAME) -v
