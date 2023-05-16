.PHONY: jsonrpc

GOBIN = ./build/bin
GO ?= latest
GORUN = go

all: jsonrpc alert-server


jsonrpc:
	@mkdir -p ${GOBIN}
	$(GORUN) build ./cmd/jsonrpc
	@mv jsonrpc ${GOBIN}

alert-server:
	@mkdir -p ${GOBIN}
	$(GORUN) build ./cmd/alert-server
	@mv alert-server ${GOBIN}
