package game

import (
	"github.com/buger/goterm"
	"github.com/davidwmartines/ticktackgo/internal/board"
	"github.com/davidwmartines/ticktackgo/internal/constants"
)

func render(gameBoard *board.Board) {
	goterm.Clear()
	goterm.MoveCursor(1, 1)
	goterm.Println("Tick-Tack-Go!")
	goterm.Printf("You are %v, playing against me (the computer), %v.\n", constants.PlayerChar, constants.CompChar)

	goterm.Println()
	gameBoard.Draw()
	goterm.Println()
	goterm.Flush()
}
