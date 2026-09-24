# aperture-phoenixd

Standalone Go module: a Phoenixd challenger for Aperture's L402 authentication. It implements Aperture's `mint.Challenger` and `auth.InvoiceChecker` interfaces against a Phoenixd Lightning node instead of LND. Only `strictVerify=false` is supported (Aperture's default), because Phoenixd's WebSocket does not emit invoice cancellation events.

## Build & Test

| Command | Purpose |
|---------|---------|
| `go build ./...` | Build |
| `go test ./...` | Run all tests |
| `go test -race ./...` | Run with race detector |
| `go vet ./...` | Static analysis |
| `golangci-lint run ./...` | Lint (config in `.golangci.yml`) |

## Structure

```
client.go              # Phoenixd HTTP client (createinvoice, getpayment)
challenger.go           # PhoenixdChallenger (NewChallenge, VerifyInvoiceStatus)
doc.go                  # Package documentation
cmd/echo-server/        # Minimal demo API for Aperture to proxy
testdata/               # Aperture integration patch and test fixtures
```

## Conventions

- British English in prose and comments: colour, initialise, behaviour, licence.
- Go standard layout: binaries in `cmd/`.
- `testify/require` for all test assertions.
- Commit messages use `type: description` format; no `Co-Authored-By` lines.
- golangci-lint with the Aperture-compatible linter set.

## Key Files

| File | Purpose |
|------|---------|
| `client.go` | Wraps the Phoenixd HTTP API: `POST /createinvoice`, `GET /payments/incoming/{hash}` |
| `challenger.go` | `PhoenixdChallenger`, using `Client` to create invoices and verify payment status |
| `doc.go` | Package documentation |
| `testdata/aperture-patch.diff` | Diff to wire this challenger into Aperture's `aperture.go` |

## Common Pitfalls

- `strictVerify=true` is not supported; do not add code paths that assume it works.
- All HTTP requests use a 10-second timeout context; keep this when extending `Client`.
- This package has no dependency on Aperture or LND: it uses plain Go types (`[32]byte` for payment hashes) that are assignment-compatible with Aperture's `lntypes.Hash`. Do not add an Aperture or LND import.
