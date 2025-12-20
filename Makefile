
.PHONY: example test fmt lint clean release
EXAMPLE_FILE=./example/main.go
EXAMPLE_BIN=./bin/exmaple
V=v1.0.0

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

release:
	@echo "Releasing $(V)..."
	git tag $(V)
	git push origin $(V)
	@echo "Release $(V) pushed to GitHub!"
	@echo "Updating Go proxy..."
	GOPROXY=proxy.golang.org go list -m github.com/cushydigit/go-freelancer-sdk/freelancer/v1@$(V)
