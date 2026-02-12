package model

// Move represents a single disk move in the Towers of Hanoi puzzle.
type Move struct {
	Disk int
	From string
	To   string
}
