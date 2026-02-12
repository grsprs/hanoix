package solver

import "github.com/grsprs/HanoiX-/internal/model"

// Solver defines the interface for Towers of Hanoi solvers.
type Solver interface {
	Solve(n int, from, to, aux string) []model.Move
}
