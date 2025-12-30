
.PHONY: example test fmt lint clean release test_cover test_cover_show
EXAMPLE_FILE=./example/main.go
EXAMPLE_BIN=./bin/exmaple
V=v1.2.2

build:
	@echo "Building..."
	@go build -o $(EXAMPLE_BIN) $(EXAMPLE_FILE)
run: build
	@echo "Running..."
	@$(EXAMPLE_BIN)
fmt:
	@go fmt ./...
lint:
	@go vet ./...
clean:
	@rm -rf bin/
test:
	@go test -v ./...
test_cover:
	@go test -cover ./...

test_cover_show:
	@go test -coverprofile=coverage.out ./freelancer/... && go tool cover -html=coverage.out



release:
	@echo "Releasing $(V)..."
	git tag $(V)
	git push origin $(V)
	@echo "Release $(V) pushed to GitHub!"
	@echo "Updating Go proxy..."
	GOPROXY=proxy.golang.org go list -m github.com/cushydigit/go-freelancer-sdk/freelancer@$(V)
