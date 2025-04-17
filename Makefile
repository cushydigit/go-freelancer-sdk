
.PHONY: example test fmt lint clean
EXAMPLE_FILE=./example/main.go
EXAMPLE_BIN=./bin/exmaple
build:
	@echo "Building..."
	@go build -o $(EXAMPLE_BIN) $(EXAMPLE_FILE)
run: build
	@echo "Running..."
	@$(EXAMPLE_BIN)
test:
	@go test -v ./...
fmt:
	@go fmt ./...
lint:
	@go vet ./...
clean:
	@rm -rf bin/
