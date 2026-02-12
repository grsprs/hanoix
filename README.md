# HanoiX: Towers of Hanoi in Go

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![CI Status](https://github.com/grsprs/hanoix/workflows/CI/badge.svg)](https://github.com/grsprs/hanoix/actions)

Professional implementation of the classic Towers of Hanoi problem in Go, demonstrating clean architecture, idiomatic Go code, comprehensive testing, and performance benchmarking.

## Author

**Spiros Nikoloudakis** ([@grsprs](https://github.com/grsprs))  
Contact: sp.nikoloudakis@gmail.com  
Copyright © 2026

## Features

- **Dual Implementations**: Recursive and iterative solvers
- **CLI Interface**: Easy-to-use command-line flags
- **Comprehensive Testing**: 100% test coverage with edge case validation
- **Performance Benchmarks**: Detailed comparison of both algorithms
- **Clean Architecture**: Modular design following Go best practices
- **CI/CD Pipeline**: Automated testing and quality checks

## Quick Start

### Installation

```bash
go install github.com/grsprs/hanoix/cmd@latest
```

Or clone and build:

```bash
git clone https://github.com/grsprs/hanoix.git
cd hanoix
go build -o hanoix ./cmd
```

### Usage

```bash
# Recursive solution with 5 disks
go run ./cmd -n=5 -mode=recursive

# Iterative solution with 5 disks
go run ./cmd -n=5 -mode=iterative
```

**Flags:**
- `-n` (int): Number of disks (must be > 0)
- `-mode` (string): Solver mode (`recursive` or `iterative`)

### Example Output

```
Move disk 1 from A to C
Move disk 2 from A to B
Move disk 1 from C to B
Move disk 3 from A to C
Move disk 1 from B to A
Move disk 2 from B to C
Move disk 1 from A to C

Total moves: 7
```

## Project Structure

```
hanoix/
├── cmd/
│   └── main.go              # CLI entry point
├── internal/
│   ├── model/
│   │   └── move.go          # Move data structure
│   └── solver/
│       ├── solver.go        # Solver interface
│       ├── recursive.go     # Recursive implementation
│       ├── iterative.go     # Iterative implementation
│       ├── *_test.go        # Unit tests
│       └── benchmark_test.go # Performance benchmarks
├── docs/
│   └── adr/
│       └── 001-iterative-algorithm.md # Architecture decisions
├── .github/
│   ├── workflows/
│   │   └── ci.yml           # CI/CD pipeline
│   └── ISSUE_TEMPLATE/      # Issue templates
├── go.mod
├── LICENSE
├── README.md
├── CONTRIBUTING.md
└── CHANGELOG.md
```

## Testing

### Run All Tests

```bash
go test ./...
```

### Run Tests with Coverage

```bash
go test -cover ./...
```

**Test Coverage:** 100% for solver package

### Test Cases

- ✅ Move count validation (2^n - 1)
- ✅ Move correctness (legal moves only)
- ✅ Edge cases (n=0, n=1)
- ✅ Both recursive and iterative implementations

## Benchmarks

### Run Benchmarks

```bash
go test -bench=. ./internal/solver
go test -bench=. -benchmem ./internal/solver
```

### Performance Results

**Execution Time:**

| Algorithm | n=10 | n=15 | n=20 |
|-----------|------|------|------|
| Recursive | 52 μs | 4.1 ms | 96.7 ms |
| Iterative | 239 μs | 12.0 ms | 251.9 ms |

**Memory Allocations:**

| Algorithm | n=10 | n=15 | n=20 |
|-----------|------|------|------|
| Recursive | 89 KB (11 allocs) | 6.76 MB (24 allocs) | 215 MB (39 allocs) |
| Iterative | 90 KB (22 allocs) | 6.76 MB (35 allocs) | 215 MB (52 allocs) |

### Performance Analysis

**Recursive Solution:**
- ✅ **3-4x faster** execution time
- ✅ Fewer memory allocations
- ✅ Simpler, more readable code
- ⚠️ Stack depth limited by recursion

**Iterative Solution:**
- ✅ No recursion depth limit
- ✅ Predictable stack usage
- ⚠️ Slower due to state tracking overhead
- ⚠️ More complex implementation

**Recommendation:** Use recursive for most cases (n < 25). Use iterative for very large n or stack-constrained environments.

## Complexity Analysis

### Time Complexity

**Both algorithms:** O(2^n)
- Must perform exactly 2^n - 1 moves
- Each move is O(1)
- Total: O(2^n)

### Space Complexity

**Recursive:** O(n)
- Call stack depth: n
- Move storage: O(2^n)

**Iterative:** O(n)
- Explicit state tracking: O(n)
- Move storage: O(2^n)

### Mathematical Properties

- **Minimum moves:** 2^n - 1 (proven optimal)
- **Move pattern:** Follows binary counting pattern
- **Disk k moves:** 2^(k-1) times

## Development

### Prerequisites

- Go 1.21 or later
- Git

### Build

```bash
go build -o hanoix ./cmd
```

### Code Quality

```bash
# Format code
go fmt ./...

# Static analysis
go vet ./...

# Run linter (if installed)
golangci-lint run
```

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Commit Convention

Follow [Conventional Commits](https://www.conventionalcommits.org/):
- `feat:` new feature
- `fix:` bug fix
- `docs:` documentation
- `test:` tests
- `refactor:` code refactoring

## License

MIT License - see [LICENSE](LICENSE) file for details.

Copyright © 2026 Spiros Nikoloudakis

## Acknowledgments

This project demonstrates professional Go development practices including:
- Clean architecture and separation of concerns
- Interface-based design
- Comprehensive testing (unit tests + benchmarks)
- CI/CD automation
- Proper documentation and project structure

## References

- [Towers of Hanoi - Wikipedia](https://en.wikipedia.org/wiki/Tower_of_Hanoi)
- [Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
