package board

import "github.com/davidwmartines/ticktackgo/internal/grid"

// Square is a space on the board on which a player can place their piece.
type Square struct {
	point *grid.Point
	id    string
	value string
}

// IsEmpty returns true if the square does not yet contain a player piece.
func (square *Square) IsEmpty() bool {
	return square.value != square.id
}
