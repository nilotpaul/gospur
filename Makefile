run: build
	@./bin/gospur

build: 
	@go build -o bin/gospur main.go