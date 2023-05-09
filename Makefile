.PHONY: jsonrpc

GOBIN = ./build/bin
GO ?= latest
GORUN = go

all:
	@mkdir -p ${GOBIN}
	$(GORUN) build ./cmd/jsonrpc
	@mv jsonrpc ${GOBIN}
