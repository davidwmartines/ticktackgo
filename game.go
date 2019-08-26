package main

import (
	"fmt"
	"strconv"
)

const size int = 3

const playerChar string = "X"
const compChar string = "O"

type square struct {
	row int
	col int
	val string
}

var board [size][size]*square
var boardMap map[string]*square

var winner = false

func main() {

	initializeBoard()
	drawBoard()

	for !winner {
		playerGo()
		drawBoard()
		computerGo()
		drawBoard()
		checkWinner()
	}
}

func playerGo() {
	var input string
	fmt.Println("pick a square: ")
	fmt.Scanln(&input)
	if sq, valid := boardMap[input]; valid {
		sq.val = playerChar
	} else {
		fmt.Println("?")
		playerGo()
	}
}

func computerGo() {
	fmt.Println("computer turn...")
	went := false
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			sq := board[row][col]
			if sq.val != playerChar && sq.val != compChar {
				sq.val = compChar
				went = true
				break
			}
		}
		if went {
			break
		}
	}
}

func checkWinner() {

}

func initializeBoard() {
	boardMap = make(map[string]*square)
	number := 0
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			number++
			numberVal := strconv.Itoa(number)
			sq := square{row, col, numberVal}
			board[row][col] = &sq
			boardMap[numberVal] = &sq
		}
	}
}

func drawBoard() {
	for row := 0; row < size; row++ {
		rowText := ""
		for col := 0; col < size; col++ {
			rowText += " " + board[row][col].val
		}
		fmt.Println(rowText)
	}
}
