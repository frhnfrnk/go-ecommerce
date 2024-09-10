build:
	@go build -o bin/go-ecommerce cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-ecommerce