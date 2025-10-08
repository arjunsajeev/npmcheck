# Release Process Guide

## Prerequisites

1. **Create Homebrew Tap Repository**
   - Create a GitHub repo: `homebrew-tap`
   - Keep it public (Homebrew requirement)

2. **Set up GitHub Token (Recommended)**
   - Create a Personal Access Token with `repo` scope
   - Add as `HOMEBREW_TAP_GITHUB_TOKEN` secret in npmcheck repo settings
   - This allows GoReleaser to automatically update the Homebrew formula

## Releasing a New Version

### 1. Prepare the Release

```bash
# Make sure everything is committed
git status

# Run checks
just ci

# Test the build locally
just snapshot
```

### 2. Create a Git Tag

```bash
# Create an annotated tag (semantic versioning)
git tag -a v1.0.0 -m "Release v1.0.0"

# Push the tag
git push origin v1.0.0
```

### 3. Monitor GitHub Actions

- Go to: <https://github.com/arjunsajeev/npmcheck/actions>
- Watch the "Release" workflow run
- It will:
  - Build binaries for Linux, macOS, Windows (amd64 and arm64)
  - Create a GitHub Release with all artifacts
  - Push the Homebrew formula to your tap

### 4. Verify the Release

Check that:

- GitHub Release is created: <https://github.com/arjunsajeev/npmcheck/releases>
- Homebrew formula is in tap: <https://github.com/arjunsajeev/homebrew-tap>
- You can install it:

  ```bash
  brew install arjunsajeev/tap/npmcheck
  ```

## Installation Instructions for Users

### Homebrew (macOS and Linux)

```bash
brew install arjunsajeev/tap/npmcheck
```

### Direct Download

Download binaries from: <https://github.com/arjunsajeev/npmcheck/releases>

### Go Install

```bash
go install github.com/arjunsajeev/npmcheck@latest
```

## Troubleshooting

### Release Failed?

1. Check GitHub Actions logs
2. Validate goreleaser config: `just check`
3. Ensure GitHub token has correct permissions

### Formula Not Updated?

1. Check if `HOMEBREW_TAP_GITHUB_TOKEN` is set correctly
2. Verify the homebrew-tap repository exists and is public
3. Check the workflow has `contents: write` permission

### Can't Install via Homebrew?

1. Update brew: `brew update`
2. Check the tap: `brew tap arjunsajeev/tap`
3. Verify formula exists in homebrew-tap repo

## Version Numbering

Follow semantic versioning:

- `v1.0.0` - Major release (breaking changes)
- `v1.1.0` - Minor release (new features, backwards compatible)
- `v1.1.1` - Patch release (bug fixes)

## Quick Commands

```bash
# Test locally before release
just snapshot

# Check goreleaser config
just check

# Create and push tag
git tag -a v1.0.0 -m "Release v1.0.0" && git push origin v1.0.0
```
