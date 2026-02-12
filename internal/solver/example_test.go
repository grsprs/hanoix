package solver_test

import (
	"fmt"

	"github.com/grsprs/hanoix/internal/solver"
)

// ExampleRecursive demonstrates basic usage of the recursive solver.
func ExampleRecursive() {
	r := &solver.Recursive{}
	moves := r.Solve(3, "A", "C", "B")

	fmt.Printf("Total moves: %d\n", len(moves))
	fmt.Printf("First move: disk %d from %s to %s\n", moves[0].Disk, moves[0].From, moves[0].To)
	// Output:
	// Total moves: 7
	// First move: disk 1 from A to C
}

// ExampleIterative demonstrates basic usage of the iterative solver.
func ExampleIterative() {
	it := &solver.Iterative{}
	moves := it.Solve(3, "A", "C", "B")

	fmt.Printf("Total moves: %d\n", len(moves))
	// Output:
	// Total moves: 7
}
