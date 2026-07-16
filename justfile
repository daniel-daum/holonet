run:
    go run .
test:
    go test ./...

# format everything, then static analysis
lint:
    golangci-lint fmt
    golangci-lint run

# run before committing
check: lint test
