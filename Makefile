check:
	gofumpt -l -w .
	golangci-lint run

.PHONY: check
