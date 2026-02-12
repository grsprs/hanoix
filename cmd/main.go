package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/grsprs/HanoiX-/internal/solver"
)

func main() {
	n := flag.Int("n", 0, "number of disks")
	mode := flag.String("mode", "", "solver mode: recursive or iterative")
	flag.Parse()

	if err := validate(*n, *mode); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		flag.Usage()
		os.Exit(1)
	}

	var s solver.Solver
	switch *mode {
	case "recursive":
		s = &solver.Recursive{}
	case "iterative":
		s = &solver.Iterative{}
	}

	moves := s.Solve(*n, "A", "C", "B")

	for _, m := range moves {
		fmt.Printf("Move disk %d from %s to %s\n", m.Disk, m.From, m.To)
	}
	fmt.Printf("\nTotal moves: %d\n", len(moves))
}

func validate(n int, mode string) error {
	if n < 0 {
		return fmt.Errorf("number of disks cannot be negative")
	}
	if n == 0 {
		return fmt.Errorf("number of disks must be greater than 0")
	}
	if mode != "recursive" && mode != "iterative" {
		return fmt.Errorf("mode must be 'recursive' or 'iterative'")
	}
	return nil
}
