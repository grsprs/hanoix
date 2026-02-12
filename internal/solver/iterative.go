package solver

import "github.com/grsprs/hanoix/internal/model"

// Iterative implements the iterative solution for Towers of Hanoi.
type Iterative struct{}

// Solve returns all moves using iterative algorithm.
func (it *Iterative) Solve(n int, from, to, aux string) []model.Move {
	if n == 0 {
		return []model.Move{}
	}

	pegs := map[string][]int{
		from: make([]int, n),
		aux:  {},
		to:   {},
	}
	for i := 0; i < n; i++ {
		pegs[from][i] = n - i
	}

	var moves []model.Move
	totalMoves := (1 << n) - 1

	for i := 1; i <= totalMoves; i++ {
		var fromPeg, toPeg string

		if n%2 == 1 {
			switch i % 3 {
			case 1:
				fromPeg, toPeg = from, to
			case 2:
				fromPeg, toPeg = from, aux
			case 0:
				fromPeg, toPeg = aux, to
			}
		} else {
			switch i % 3 {
			case 1:
				fromPeg, toPeg = from, aux
			case 2:
				fromPeg, toPeg = from, to
			case 0:
				fromPeg, toPeg = aux, to
			}
		}

		if len(pegs[fromPeg]) == 0 || (len(pegs[toPeg]) > 0 && pegs[toPeg][len(pegs[toPeg])-1] < pegs[fromPeg][len(pegs[fromPeg])-1]) {
			fromPeg, toPeg = toPeg, fromPeg
		}

		disk := pegs[fromPeg][len(pegs[fromPeg])-1]
		pegs[fromPeg] = pegs[fromPeg][:len(pegs[fromPeg])-1]
		pegs[toPeg] = append(pegs[toPeg], disk)

		moves = append(moves, model.Move{Disk: disk, From: fromPeg, To: toPeg})
	}

	return moves
}
