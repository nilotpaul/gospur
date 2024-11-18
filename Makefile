BUILD_DIR = ./bin
BINARY = $(BUILD_DIR)/gospur

# CLI Commands
init: build
	@$(BINARY) init

# Development Commands
run: build
	@$(BINARY)

build: 
	@go build -o $(BINARY) .

test:
	@go test -v ./...

test-race:
	@go test -v ./... --race