APP_NAME := dotfiles-cli

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOINSTALL := $(GOCMD) install

BINARY_NAME := $(GOBIN)/$(APP_NAME)

.PHONY: all build clean install

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

install:
	$(GOINSTALL)

run: build
	$(BINARY_NAME)
