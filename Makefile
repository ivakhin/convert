.PHONY: test
test:
	go test -v -race -cover ./...

.PHONY: benchmark
benchmark:
	go test -bench ./...

.PHONY: linter-install
linter-install:
	 go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: lint
lint:
	golangci-lint run

.PHONY: badge
badge:
	go test ./... -covermode=count -coverprofile=coverage.out fmt
	go tool cover -func=coverage.out -o=coverage.out
	gobadge -filename=coverage.out
	rm coverage.out