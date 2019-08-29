package board

import "github.com/davidwmartines/ticktackgo/internal/grid"

// Square is a space on the board on which a player can place their piece.
type Square struct {
	Point *grid.Point
	ID    string
	Value string
}

// IsEmpty returns true if the square does not yet contain a player piece.
func (square *Square) IsEmpty() bool {
	return square.Value == square.ID
}

//RowMatch returns true if the the squares on on the same row.
func (square *Square) RowMatch(other *Square) bool {
	return square.Point.Row == other.Point.Row
}

//ColMatch returns true if the the squares on on the same column.
func (square *Square) ColMatch(other *Square) bool {
	return square.Point.Col == other.Point.Col
}
