# npmcheck

Check npm package name availability from the command line.

## Installation

### Homebrew (macOS and Linux)

```bash
brew install arjunsajeev/tap/npmcheck
```

### Go Install

```bash
go install github.com/arjunsajeev/npmcheck@latest
```

### Download Binary

Download the latest binary from the [releases page](https://github.com/arjunsajeev/npmcheck/releases).

## Usage

Check if one or more package names are available on npm:

```bash
npmcheck my-awesome-package another-package
```

Output:

- ✅ indicates the package name is available
- ❌ indicates the package name is already taken
- ⚠️ indicates an error occurred while checking

## Examples

```bash
# Check a single package
npmcheck my-new-package

# Check multiple packages at once
npmcheck pkg1 pkg2 pkg3
```

## Development

### Prerequisites

- Go 1.25 or later

### Building

Using `just` (recommended):

```bash
# Install just: brew install just (or see https://github.com/casey/just)

# List all available commands
just

# Build the binary
just build

# Run tests
just test

# Run all checks
just ci
```

Or using `make`:

```bash
# Build the binary
make build

# Run tests
make test
```

Or directly with Go:

```bash
# Build the binary
go build -o npmcheck

# Run tests
go test ./...

# Install locally
go install
```

### Testing GoReleaser

To test the GoReleaser configuration locally without publishing:

```bash
goreleaser release --snapshot --clean
```

## Releasing

This project uses [GoReleaser](https://goreleaser.com/) for automated releases.

To create a new release:

1. Tag the commit with a semantic version:

   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

2. GitHub Actions will automatically:
   - Build binaries for multiple platforms (Linux, macOS, Windows)
   - Create a GitHub release
   - Update the Homebrew formula

## License

MIT
