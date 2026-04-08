GRC = $(shell command -v grc 2>/dev/null)

build:
	@go build -o bin/api/sylvie_api cmd/api/main.go
	@go build -o bin/processor/sylvie_worker cmd/processor/main.go

run-api:
	@go run ./cmd/api/main.go

run-processor:
	@go run ./cmd/processor/main.go

views:
	@go tool templ generate

test:
	@$(GRC) go test -v -cover -failfast ./...

clean:
	@rm -rf bin
	@find -type d -name "tmp" -exec rm -rf {} +
	@find -type d -name "transcoder_test_input" -exec rm -rf {} +

.PHONY: test clean run-api run-worker
