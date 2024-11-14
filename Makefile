check:
	gofumpt -l -w .
	golangci-lint run

test:
	go test ./...

.PHONY: check test
