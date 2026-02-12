# ADR 001: Iterative Algorithm Implementation

## Status
Accepted

## Context
The Towers of Hanoi problem can be solved using either recursive or iterative approaches. While recursion is more intuitive, we needed to provide both implementations for comparison and to handle stack-constrained environments.

## Decision
Implement an iterative solution using explicit state tracking with peg state management, rather than stack-based recursion simulation.

## Consequences

### Positive
- No recursion depth limit
- Predictable memory usage
- Works in stack-constrained environments
- Demonstrates alternative algorithmic thinking

### Negative
- More complex implementation (state tracking)
- Slower execution (3-4x) due to overhead
- More allocations per operation

### Trade-offs
- Chose correctness and clarity over raw performance
- State tracking ensures legal moves validation
- Acceptable performance penalty for flexibility gained

## Alternatives Considered
1. Stack-based recursion simulation - rejected due to complexity
2. Binary counting pattern - considered but state tracking chosen for clarity

## Date
2026-01-01
