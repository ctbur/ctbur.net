PKG_DIR := ./cmd/server
BINARY_NAME := server

.PHONY: build dev test clean

build:
	CGO_ENABLED=0 go build -o bin/$(BINARY_NAME) $(PKG_DIR)

dev:
	CGO_ENABLED=0 air

test:
	go test -v ./...

clean:
	@rm -f $(BINARY_NAME)
