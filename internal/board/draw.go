package board

import "github.com/buger/goterm"

//Draw renders the board using goterm.
func (board *Board) Draw() {
	for _, row := range board.grid {
		for _, point := range row {
			goterm.Print(" ")
			square := board.squaresByPoint[point]
			val := square.Value
			if square.IsEmpty() {
				val = goterm.Color(val, goterm.WHITE)
			} else if val == "X" {
				val = goterm.Color(val, goterm.GREEN)
			} else {
				val = goterm.Color(val, goterm.RED)
			}
			goterm.Print(val)
		}
		goterm.Println()
	}
}
