GO_FILES = $(shell find . -type f -name '*.go')
GOPATH = $(shell pwd):$(shell pwd)/vendor
VENDORS_PATH = $(shell pwd)/vendor
all: setup build

build: $(GO_FILES)
	@GOPATH=$(GOPATH) go build -o bin/for

test:
	go test config -v

clean:
	rm -rf bin/* pkg/*

setup:
	@GOPATH=$(VENDORS_PATH) go get github.com/golang/lint/golint
	@GOPATH=$(VENDORS_PATH) go get github.com/BurntSushi/toml