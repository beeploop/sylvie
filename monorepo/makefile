build:
	@go build -o bin/api/sylvie_api cmd/api/main.go
	@go build -o bin/worker/sylvie_worker cmd/worker/main.go

run-api:
	@go run ./cmd/api/main.go

run-worker:
	@go run ./cmd/worker/main.go

test:
	@go test -v -cover -failfast ./...

clean:
	@rm -rf bin
	@find -type d -name "tmp" -exec rm -rf {} +
	@find -type d -name "transcoder_test_input" -exec rm -rf {} +

.PHONY: test clean run-api run-worker
