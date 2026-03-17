# Contributing

## Setup

```bash
git clone https://github.com/TheCryptoDonkey/aperture-phoenixd.git
cd aperture-phoenixd
go build ./...
```

Requires Go 1.24+.

## Testing

```bash
go test ./...          # Run all tests
go test -race ./...    # Run with race detector
go vet ./...           # Static analysis
```

## Linting

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5
golangci-lint run ./...
```

The linter configuration is in `.golangci.yml`.

## Code style

- **British English** — colour, initialise, behaviour, licence
- **testify/require** for all test assertions
- **Go standard layout** — binaries in `cmd/`
- Commit messages: `type: description` (e.g. `feat:`, `fix:`, `docs:`)

## Pull requests

1. Fork and create a branch from `main`
2. Make your changes
3. Ensure `go test -race ./...` and `golangci-lint run ./...` pass
4. Open a PR against `main`

## Licence

By contributing you agree that your contributions will be licensed under the [MIT Licence](LICENSE).
