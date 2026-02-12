package solver

import (
	"testing"
)

// TestIterativeMoveCount verifies total moves equals 2^n - 1
func TestIterativeMoveCount(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 3},
		{3, 7},
		{5, 31},
		{10, 1023},
	}

	it := &Iterative{}
	for _, tt := range tests {
		moves := it.Solve(tt.n, "A", "C", "B")
		if len(moves) != tt.expected {
			t.Errorf("n=%d: got %d moves, want %d", tt.n, len(moves), tt.expected)
		}
	}
}

// TestIterativeCorrectness validates move correctness
func TestIterativeCorrectness(t *testing.T) {
	it := &Iterative{}
	moves := it.Solve(3, "A", "C", "B")

	pegs := map[string][]int{
		"A": {3, 2, 1},
		"B": {},
		"C": {},
	}

	for _, m := range moves {
		if len(pegs[m.From]) == 0 {
			t.Fatalf("invalid move: %s is empty", m.From)
		}
		disk := pegs[m.From][len(pegs[m.From])-1]
		pegs[m.From] = pegs[m.From][:len(pegs[m.From])-1]

		if len(pegs[m.To]) > 0 && pegs[m.To][len(pegs[m.To])-1] < disk {
			t.Fatalf("invalid move: disk %d on smaller disk", disk)
		}
		pegs[m.To] = append(pegs[m.To], disk)
	}

	if len(pegs["C"]) != 3 || pegs["C"][0] != 3 || pegs["C"][2] != 1 {
		t.Errorf("final state incorrect: %v", pegs)
	}
}

// TestIterativeEdgeCases covers edge cases
func TestIterativeEdgeCases(t *testing.T) {
	it := &Iterative{}

	moves := it.Solve(0, "A", "C", "B")
	if len(moves) != 0 {
		t.Errorf("n=0: got %d moves, want 0", len(moves))
	}

	moves = it.Solve(1, "A", "C", "B")
	if len(moves) != 1 {
		t.Errorf("n=1: got %d moves, want 1", len(moves))
	}
	if moves[0].Disk != 1 || moves[0].From != "A" || moves[0].To != "C" {
		t.Errorf("n=1: incorrect move %v", moves[0])
	}
}
