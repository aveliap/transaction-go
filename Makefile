build:
	@go build -o bin/transaction-go cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/transaction-go

migration:
# $(HOME)/go/bin/migrate means migrate keyword
	@$(HOME)/go/bin/migrate create -ext sql -dir cmd/migrate/migration $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down