BUILD_DIR = ./bin
BINARY = $(BUILD_DIR)/gospur

# CLI Commands
init: build
	@$(BINARY) init

run: build
	@$(BINARY)

build: 
	@go build -o $(bin/gospur) .