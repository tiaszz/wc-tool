build:
	@go build -o bin/wc

run: build
	@./bin/wc

test:
	@go test ./... --count=1
