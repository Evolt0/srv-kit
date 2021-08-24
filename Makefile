PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
NAME = $(shell basename $(PKG))
WORKSPACE ?= $(NAME)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GO = CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go

GORUN = cd ./cmd/$(WORKSPACE) && $(GO) run
GOBUILD=$(GO) build

run: download
	$(GORUN) .

download:
	go mod download -x

build:
	cd ./cmd/$(WORKSPACE) && $(GOBUILD) -o ./$(WORKSPACE)-$(GOOS)-$(GOARCH)