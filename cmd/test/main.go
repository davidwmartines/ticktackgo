package main

import "github.com/davidwmartines/ticktackgo/internal/board"

func main() {
	board := board.New(3)
	board.Draw()
}
