lint:
	golangci-lint run
test:
	go test ./... -race
format:
	gofumpt -l -w .
