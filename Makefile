build:
	@go build -o bin/transaction-go cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/transaction-go