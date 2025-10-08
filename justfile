# npmcheck justfile
# Install just: https://github.com/casey/just

# List available recipes
default:
    "@just --list"

# Build the binary
build:
    go build -o npmcheck

# Run tests
test:
    go test -v ./...

# Run tests with coverage
test-coverage:
    go test -v -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out

# Clean build artifacts
clean:
    rm -f npmcheck
    rm -rf dist/
    rm -f coverage.out

# Install locally
install:
    go install

# Run the tool with arguments
run *args:
    go run main.go {{args}}

# Format code
fmt:
    go fmt ./...

# Tidy dependencies
tidy:
    go mod tidy

# Test goreleaser locally without publishing
snapshot:
    goreleaser release --snapshot --clean

# Validate goreleaser configuration
check:
    goreleaser check

# Run goreleaser in release mode (requires proper git tag)
release:
    goreleaser release --clean

# Run all checks (fmt, test)
ci: fmt test
    @echo "✅ All checks passed!"

# Create and push a new version tag (e.g., just tag v1.2.3)
tag version:
    git tag -a {{version}} -m "Release {{version}}"
    git push origin {{version}}
    @echo "✅ Tag {{version}} created and pushed!"
