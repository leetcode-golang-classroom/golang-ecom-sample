.PHONY=build

build:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/main cmd/main.go

run: build
	@./bin/main

coverage:
	@go test -v -cover ./...

test:
	@go test -v ./...

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down