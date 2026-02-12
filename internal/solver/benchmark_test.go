package solver

import (
	"testing"
)

func BenchmarkRecursive10(b *testing.B) {
	r := &Recursive{}
	for i := 0; i < b.N; i++ {
		r.Solve(10, "A", "C", "B")
	}
}

func BenchmarkRecursive15(b *testing.B) {
	r := &Recursive{}
	for i := 0; i < b.N; i++ {
		r.Solve(15, "A", "C", "B")
	}
}

func BenchmarkRecursive20(b *testing.B) {
	r := &Recursive{}
	for i := 0; i < b.N; i++ {
		r.Solve(20, "A", "C", "B")
	}
}

func BenchmarkIterative10(b *testing.B) {
	it := &Iterative{}
	for i := 0; i < b.N; i++ {
		it.Solve(10, "A", "C", "B")
	}
}

func BenchmarkIterative15(b *testing.B) {
	it := &Iterative{}
	for i := 0; i < b.N; i++ {
		it.Solve(15, "A", "C", "B")
	}
}

func BenchmarkIterative20(b *testing.B) {
	it := &Iterative{}
	for i := 0; i < b.N; i++ {
		it.Solve(20, "A", "C", "B")
	}
}
