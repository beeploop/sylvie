build:
	@go build -o bin/api/sylvie_api cmd/api/main.go
	@go build -o bin/worker/sylvie_worker cmd/worker/main.go

run-api:
	@go run ./cmd/api/main.go

run-worker:
	@go run ./cmd/worker/main.go

test:
	@go test -v -failfast ./...

clean:
	@rm -rf bin

.PHONY: test clean run-api run-worker
