# Go CI/CD Workflow Template

A production-ready GitHub Actions workflow template for Go projects with comprehensive CI/CD, security scanning, and automation.

## Features

- **CI Pipeline**: Linting, testing, building, and commit signature verification
- **Security Scanning**: Secret detection, vulnerability scanning, SBOM generation
- **Release Automation**: Tag-based releases with GoReleaser
- **Performance Tracking**: Benchmarks with regression detection
- **Integration Testing**: Template for service-dependent tests
- **Local Automation**: 30+ Taskfile commands for development

## Quick Start

1. Copy this template to your Go project root
2. Replace placeholders:
   - `your-app` → your binary name
   - `your-org/your-repo` → your GitHub repository
   - `./cmd/your-app` → your main package path
3. Commit and push

```bash
# Copy template files
cp -r go-workflow-template/.github your-project/
cp go-workflow-template/.golangci.yml your-project/
cp go-workflow-template/.goreleaser.yml your-project/
cp go-workflow-template/.gitattributes your-project/
cp go-workflow-template/.gitignore your-project/
cp go-workflow-template/Taskfile.yml your-project/
```

## 3 Must-Have Improvements

Every workflow includes these best practices:

### 1. Caching

Dramatically reduces build times by caching Go modules and tool outputs.

```yaml
- name: Set up Go
  uses: actions/setup-go@v5
  with:
    go-version: stable
    cache: true  # Built-in module caching
```

### 2. Concurrency Control

Prevents resource waste by canceling redundant workflow runs.

```yaml
concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true  # Cancel older runs on new push
```

> **Note**: Release workflows use `cancel-in-progress: false` to never interrupt a release.

### 3. Job Timeouts

Prevents hung jobs from consuming resources indefinitely.

```yaml
jobs:
  build:
    timeout-minutes: 15  # Fails if job exceeds this
```

## Workflows

### CI (`ci.yml`)

Runs on: push to main, pull requests, manual trigger

| Job | Timeout | Description |
|-----|---------|-------------|
| `lint` | 10 min | golangci-lint, gofmt, go vet, go mod tidy |
| `test` | 15 min | Tests with race detection and coverage |
| `build` | 10 min | Build and verify binary |
| `verify-signatures` | 5 min | Check commit signatures (PRs only) |

### Security (`security.yml`)

Runs on: push to main, pull requests, weekly schedule, manual trigger

| Job | Timeout | Description |
|-----|---------|-------------|
| `gitleaks` | 10 min | Scan for hardcoded secrets |
| `govulncheck` | 15 min | Check for known vulnerabilities |
| `gosec` | 15 min | Static security analysis |
| `dependency-review` | 10 min | Review dependency changes (PRs only) |
| `sbom` | 10 min | Generate SBOM (main branch only) |

### Release (`release.yml`)

Runs on: version tags (`v*`)

| Job | Timeout | Description |
|-----|---------|-------------|
| `verify` | 10 min | Validate tag format, generate changelog |
| `build` | 30 min | GoReleaser build, sign, publish |

**Creating a release:**
```bash
git tag -s v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

### Benchmarks (`benchmarks.yml`)

Runs on: push to main, pull requests, manual trigger

| Job | Timeout | Description |
|-----|---------|-------------|
| `benchmark` | 20 min | Run benchmarks, compare with baseline |

- Alerts on 150% regression (main branch)
- Fails on 200% regression (PRs)

### Integration Tests (`integration-test.yml`)

Runs on: push to main, pull requests, manual trigger

| Job | Timeout | Description |
|-----|---------|-------------|
| `integration` | 20 min | Run tests with `-tags=integration` |

## Configuration Files

### `.golangci.yml`

35+ linters including:
- **Error handling**: errcheck, errorlint, nilerr
- **Security**: gosec
- **Performance**: prealloc, goconst
- **Style**: gofmt, goimports, revive, stylecheck
- **Complexity**: gocyclo, gocognit, nestif

### `.goreleaser.yml`

Cross-platform release automation:
- **Platforms**: Linux, macOS, Windows (amd64, arm64)
- **Packaging**: tar.gz, zip
- **Distribution**: GitHub Releases, Docker, Homebrew, Scoop
- **Security**: GPG signing, checksums, SBOM

### `.gitattributes`

- LF line ending normalization
- Go-specific diff highlighting
- Binary file handling
- Export ignores for clean archives

### `.gitignore`

Comprehensive ignores for:
- Go build artifacts
- IDE/editor files
- OS-generated files
- Secrets and credentials
- Test outputs and coverage

## Taskfile Commands

Install Task: `go install github.com/go-task/task/v3/cmd/task@latest`

### Development

```bash
task dev          # Run with live reload (requires air)
task run          # Run the application
```

### Building

```bash
task build        # Build binary
task build:all    # Build for all platforms
task build:linux  # Build for Linux (amd64, arm64)
task build:darwin # Build for macOS (amd64, arm64)
task build:windows # Build for Windows
```

### Testing

```bash
task test              # Run all tests
task test:short        # Run short tests only
task test:coverage     # Generate coverage report
task test:coverage:view # Open coverage in browser
task test:integration  # Run integration tests
```

### Benchmarks

```bash
task bench         # Run benchmarks
task bench:compare # Compare with previous run
```

### Linting & Formatting

```bash
task lint      # Run all linters
task lint:fix  # Fix linter issues
task fmt       # Format code
task fmt:check # Check formatting
task vet       # Run go vet
```

### Dependencies

```bash
task deps        # Download dependencies
task deps:tidy   # Tidy go.mod
task deps:verify # Verify dependencies
task deps:update # Update all dependencies
task deps:vuln   # Check for vulnerabilities
```

### Security

```bash
task security         # Run all security checks
task security:vuln    # Vulnerability scan
task security:gosec   # Static analysis
task security:secrets # Secret detection
```

### Release

```bash
task release:snapshot # Test release locally
task release:check    # Validate config
```

### Docker

```bash
task docker:build # Build Docker image
task docker:run   # Run container
```

### Utilities

```bash
task tools      # Install all required tools
task clean      # Clean build artifacts
task clean:all  # Clean everything
task ci         # Run full CI locally
task pre-commit # Pre-commit checks
```

## Required Secrets

Configure these in your repository settings:

| Secret | Required | Description |
|--------|----------|-------------|
| `GITHUB_TOKEN` | Auto | Provided by GitHub Actions |
| `CODECOV_TOKEN` | Optional | For coverage uploads |
| `GPG_PRIVATE_KEY` | Optional | For signing releases |
| `GPG_PASSPHRASE` | Optional | GPG key passphrase |
| `GPG_FINGERPRINT` | Optional | GPG key fingerprint |

## Customization

### Adding a New Linter

Edit `.golangci.yml`:

```yaml
linters:
  enable:
    - your-new-linter
```

### Changing Build Targets

Edit `.goreleaser.yml`:

```yaml
builds:
  - goos:
      - linux
      - darwin
      - windows
      - freebsd  # Add new OS
    goarch:
      - amd64
      - arm64
      - arm  # Add new arch
```

### Adding Integration Test Services

Edit `.github/workflows/integration-test.yml`:

```yaml
jobs:
  integration:
    services:
      postgres:
        image: postgres:16
        env:
          POSTGRES_PASSWORD: postgres
        ports:
          - 5432:5432
```

## File Structure

```
your-project/
├── .github/
│   └── workflows/
│       ├── ci.yml              # Main CI pipeline
│       ├── security.yml        # Security scanning
│       ├── release.yml         # Release automation
│       ├── benchmarks.yml      # Performance tracking
│       └── integration-test.yml # Integration tests
├── .golangci.yml               # Linter configuration
├── .goreleaser.yml             # Release configuration
├── .gitattributes              # Git attributes
├── .gitignore                  # Git ignores
└── Taskfile.yml                # Development tasks
```

## License

This template is provided as-is. Customize and use freely in your projects.
