package solver

import "github.com/grsprs/HanoiX-/internal/model"

// Recursive implements the recursive solution for Towers of Hanoi.
type Recursive struct{}

// Solve returns all moves using recursive algorithm.
func (r *Recursive) Solve(n int, from, to, aux string) []model.Move {
	var moves []model.Move
	r.hanoi(n, from, to, aux, &moves)
	return moves
}

func (r *Recursive) hanoi(n int, from, to, aux string, moves *[]model.Move) {
	if n == 1 {
		*moves = append(*moves, model.Move{Disk: 1, From: from, To: to})
		return
	}
	r.hanoi(n-1, from, aux, to, moves)
	*moves = append(*moves, model.Move{Disk: n, From: from, To: to})
	r.hanoi(n-1, aux, to, from, moves)
}
