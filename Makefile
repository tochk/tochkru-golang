build:
	go fmt ./...
	GOOS=linux go build -o bin/tochknet cmd/tochknet/main.go

generate:
	go generate ./...

run:
	go run cmd/tochknet/main.go -log=debug

migrate:
	go run cmd/migrations/main.go up ./migrations

migrate_down:
	go run cmd/migrations/main.go down ./migrations

test:
	go test ./...

all: generate build