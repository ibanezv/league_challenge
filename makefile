linter: ### check by golangci linter
	golangci-lint run
.PHONY: linter

test: ### run test
	go test -v -cover -race ./internal/...
.PHONY: test

run: ### run app
	go run ./cmd/main.go
.PHONY: test

build: ### build app
	go build -o league_challenge ./cmd/main.go 
.PHONY: build
