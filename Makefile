BUILD_DIR = ./bin
BINARY = $(BUILD_DIR)/gospur

# CLI Commands
init: build
	@$(BINARY) init

version: build	
	@$(BINARY) version

# Development Commands
run: build
	@$(BINARY)

build: 
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o $(BINARY) .
	@GOOS=windows GOARCH=amd64 go build -o $(BINARY).exe . 

test:
	@go test -v ./...

test-race:
	@go test -v ./... --race

# Only for local testing
release:
	@goreleaser release --snapshot --clean