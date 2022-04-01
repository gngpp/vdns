.PHONY: vdns build clean test test-race

VERSION=1.0.0
BIN=vdns
DIR_SRC=.
# cmd
DOCKER_CMD=docker

GO_ENV=CGO_ENABLED=0
GO_FLAGS=-ldflags="-X main.version=$(VERSION) -X 'main.buildTime=`date`' -extldflags -static"
GO=$(GO_ENV) $(shell which go)
GOROOT=$(shell `which go` env GOROOT)
GOPATH=$(shell `which go` env GOPATH)

#build: $(DIR_SRC)/main.go
#	@$(GO) build $(GO_FLAGS) -o $(BIN) $(DIR_SRC)
#
#build_docker_image:
#	@$(DOCKER_CMD) build -f ./Dockerfile -t vdns-server:$(VERSION) .
#
#testing:
#	@$(cd ./testing)
#	@$(GO) testing
#
#testing-race:
#	@$(GO) testing -race ./
#
## clean all build result
#clean:
#	@$(GO) clean ./
#	@rm -f $(BIN)
#	@rm -rf ./dist/*
