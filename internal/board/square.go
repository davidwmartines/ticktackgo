package board

import "github.com/davidwmartines/ticktackgo/internal/grid"

// Square is a space on the board on which a player can place their piece.
type Square struct {
	point *grid.Point
	id    string
	value string
}
