package board

import (
	"fmt"
	"strconv"

	"github.com/davidwmartines/ticktackgo/internal/grid"
)

//Board is a field for playing a game of tic tac go.
type Board struct {
	grid           grid.Grid
	squaresByID    map[string]*Square
	squaresByPoint map[*grid.Point]*Square
	allSquares     []*Square
}

// New creates a Board.
func New(size int) (board Board) {
	board.grid = grid.Create(size, size)
	board.squaresByID = make(map[string]*Square)
	board.squaresByPoint = make(map[*grid.Point]*Square)
	id := 0
	for _, row := range board.grid {
		for _, point := range row {
			id++
			val := strconv.Itoa(id)
			square := Square{point, val, val}
			board.squaresByID[val] = &square
			board.squaresByPoint[point] = &square
			board.allSquares = append(board.allSquares, &square)
		}
	}
	return
}

//DrawStdout renders the board to stdout.
func (board *Board) DrawStdout() {
	for _, row := range board.grid {
		rowText := ""
		for _, point := range row {
			rowText += " " + board.squaresByPoint[point].Value
		}
		fmt.Println(rowText)
	}
}

//Square gets a square by id.
func (board *Board) Square(id string) (*Square, bool) {
	square, ok := board.squaresByID[id]
	return square, ok
}

//Squares gets a list of the squares
func (board *Board) Squares() []*Square {
	return board.allSquares
}

// Neighbors gets a list of the square's neighbors
func (board *Board) Neighbors(source *Square) (neighbors []*Square) {
	for _, point := range board.grid.Neighbors(source.Point) {
		neighbors = append(neighbors, board.squaresByPoint[point])
	}
	return
}
