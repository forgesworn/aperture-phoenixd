# CLAUDE.md — aperture-phoenixd

Standalone Go module: Phoenixd challenger for Aperture L402 auth.

## Commands

```bash
go build ./...         # Build
go test ./...          # Run all tests
go test -race ./...    # Run with race detector
go vet ./...           # Lint
```

## Structure

```
client.go              # Phoenixd HTTP client (createinvoice, getpayment)
challenger.go          # PhoenixdChallenger (NewChallenge, VerifyInvoiceStatus)
cmd/echo-server/       # Minimal demo API for Aperture to proxy
```

## Conventions

- **British English** — colour, initialise, behaviour, licence
- **Go standard layout** — cmd/ for binaries
- **Git:** commit messages use `type: description` format
- **Git:** Do NOT include `Co-Authored-By` lines
- **testify/require** for all test assertions
- **golangci-lint** with Aperture-compatible linter set
