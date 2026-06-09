.PHONY: check test build fmt vet lint clean test-race cover

check: fmt vet lint build test

fmt:
	gofumpt -l -w .

vet:
	go vet ./...

lint:
	golangci-lint run

build:
	go build ./...

test:
	go test ./...

test-race:
	go test -race ./...

cover:
	go test -coverprofile=cover.out ./...
	go tool cover -func=cover.out
	rm cover.out

clean:
	go clean ./...
	rm -f magpie cover.out
