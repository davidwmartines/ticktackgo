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
		}
	}
	return
}

//Draw renders the board to stdout.
func (board *Board) Draw() {
	for _, row := range board.grid {
		rowText := ""
		for _, point := range row {
			rowText += " " + board.squaresByPoint[point].value
		}
		fmt.Println(rowText)
	}
}
