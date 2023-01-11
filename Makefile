GO ?= go
PROTO_FILES := $(shell find . -name "*.proto")
PROTO_GO_FILES := $(patsubst %.proto,%.proto.go,$(PROTOFILES))

.PHONY: protoc
protoc: $(PROTO_GO_FILES)

%.proto.go: %.proto
	protoc --go_out=. $<

.PHONY: golangci-lint
golangci-lint:
	golangci-lint run

.PHONY: test
test:
	$(GO) test -race -v ./...
