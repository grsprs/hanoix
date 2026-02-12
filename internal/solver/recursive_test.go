package solver

import (
	"testing"
)

// TestRecursiveMoveCount verifies total moves equals 2^n - 1
func TestRecursiveMoveCount(t *testing.T) {
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

	r := &Recursive{}
	for _, tt := range tests {
		moves := r.Solve(tt.n, "A", "C", "B")
		if len(moves) != tt.expected {
			t.Errorf("n=%d: got %d moves, want %d", tt.n, len(moves), tt.expected)
		}
	}
}

// TestRecursiveCorrectness validates move correctness
func TestRecursiveCorrectness(t *testing.T) {
	r := &Recursive{}
	moves := r.Solve(3, "A", "C", "B")

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

// TestRecursiveEdgeCases covers edge cases
func TestRecursiveEdgeCases(t *testing.T) {
	r := &Recursive{}

	moves := r.Solve(1, "A", "C", "B")
	if len(moves) != 1 {
		t.Errorf("n=1: got %d moves, want 1", len(moves))
	}
	if moves[0].Disk != 1 || moves[0].From != "A" || moves[0].To != "C" {
		t.Errorf("n=1: incorrect move %v", moves[0])
	}
}
