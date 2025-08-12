PKG_DIR := ./cmd/server
BINARY_NAME := server

.PHONY: build test clean

build:
	CGO_ENABLED=0 go build -o bin/$(BINARY_NAME) $(PKG_DIR)

test:
	go test -v ./...

clean:
	@rm -f $(BINARY_NAME)
