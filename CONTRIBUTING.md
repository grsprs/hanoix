# Contributing to HanoiX

## Prerequisites

- Go 1.21 or later
- Git

## Building

```bash
go build -o hanoix ./cmd
```

## Running

```bash
go run ./cmd -n=5 -mode=recursive
go run ./cmd -n=5 -mode=iterative
```

## Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

## Benchmarking

```bash
go test -bench=. ./...
```

## Code Quality

Format code:
```bash
go fmt ./...
```

Lint code:
```bash
go vet ./...
```

## Commit Messages

Follow Conventional Commits specification:
- `feat:` new feature
- `fix:` bug fix
- `docs:` documentation changes
- `test:` test additions or modifications
- `refactor:` code refactoring

## Pull Requests

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and ensure they pass
5. Submit a pull request

All PRs must pass CI checks before merging.
