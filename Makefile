
.PHONY: build run fmt lint test test-race test-cover test-cover-show check clean publish-check release

EXAMPLE_FILE := ./example/main.go
EXAMPLE_BIN := ./bin/example
V ?= v1.4.0

build:
	@echo "Building example..."
	@mkdir -p bin
	@go build -o $(EXAMPLE_BIN) $(EXAMPLE_FILE)

run: build
	@echo "Running example..."
	@$(EXAMPLE_BIN)

fmt:
	@echo "Formatting..."
	@go fmt ./...

lint:
	@echo "Running go vet..."
	@go vet ./...

test:
	@echo "Running tests..."
	@go test -v ./...

test-race:
	@echo "Running tests with race detector..."
	@go test -race ./...

test-cover:
	@echo "Running tests with coverage..."
	@go test -cover ./...

test-cover-show:
	@echo "Generating coverage report..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

check:
	@echo "Running release checks..."
	@go fmt ./...
	@go vet ./...
	@go test ./...
	@go test -race ./...

clean:
	@echo "Cleaning..."
	@rm -rf bin coverage.out

publish-check:
	@echo "Checking $(V) on Go proxy..."
	@GOPROXY=https://proxy.golang.org go list -m github.com/cushydigit/go-freelancer-sdk/freelancer@$(V)

release: check
	@echo "Releasing $(V)..."
	@git diff --exit-code
	@git tag $(V)
	@git push origin $(V)
	@echo "Release $(V) pushed to GitHub!"
	@echo "Checking Go proxy..."
	@$(MAKE) publish-check V=$(V)
