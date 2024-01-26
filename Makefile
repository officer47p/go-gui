build:
	@go build -o ./bin/main .

dev:
	@go run .

test:
	@go test ./...