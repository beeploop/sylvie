build:
	@go build -o bin/sylvie cmd/main.go

run:
	@go run cmd/main.go

clean:
	@rm -rf bin

test:
	@go test -v -failfast ./...

.PHONY: build run clean test
