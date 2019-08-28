package game

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const size int = 3

const playerChar string = "X"
const compChar string = "O"

type square struct {
	row int
	col int
	val string
	id  int
}

var board [size][size]*square
var boardMap map[string]*square

//Run runs the game.
func Run() {

	initializeBoard()
	drawBoard()

	for {
		playerGo()
		drawBoard()
		checkWinner()

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
		if sq.val == playerChar || sq.val == compChar {
			fmt.Printf("%v is already taken!\n", input)
			playerGo()
		}
		sq.val = playerChar
	} else {
		fmt.Printf("%v is not a valid square!\n", input)
		playerGo()
	}
}

func computerGo() {
	fmt.Println("computer turn...")
	choice := getNextSquareForComputer()
	choice.val = compChar
}

func getNextSquareForComputer() *square {
	for _, sq := range boardMap {

		// if square is owned, try getting a neighbor
		if sq.val == compChar {
			neighbors := neighbors(float64(sq.row), float64(sq.col))
			for _, n := range neighbors {
				if isEmpty(n) {
					return n
				}
			}
		}
	}
	return getRandomEmptySquare()
}

func isEmpty(sq *square) bool {
	return sq.val != compChar && sq.val != playerChar
}

func getRandomEmptySquare() *square {
	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)
	for {
		randomNumber := gen.Intn((size*size)-1) + 1
		possibleSquare := boardMap[strconv.Itoa(randomNumber)]
		if isEmpty(possibleSquare) {
			return possibleSquare
		}
	}
}

func checkWinner() {
	if isWinner(playerChar) {
		fmt.Println("You Win!")
		os.Exit(0)
	} else if isWinner(compChar) {
		fmt.Println("Computer Wins!")
		os.Exit(0)
	} else if isTie() {
		fmt.Println("no winner")
		os.Exit(0)
	}
}

func isWinner(char string) bool {
	for _, sq := range boardMap {
		if sq.val == char {
			nex := getAdjacentMatch(sq)
			if nex != nil {
				last := getInlineMatch(sq, nex)
				if last != nil {
					return true
				}
			}
		}
	}
	return false
}

func isTie() bool {
	for _, sq := range boardMap {
		if isEmpty(sq) {
			return false
		}
	}
	return true
}

func getAdjacentMatch(sq *square) *square {
	neighbors := neighbors(float64(sq.row), float64(sq.col))
	for _, n := range neighbors {
		if n.val == sq.val {
			return n
		}
	}
	return nil
}

func getInlineMatch(prev *square, sq *square) *square {
	neighbors := neighbors(float64(sq.row), float64(sq.col))
	for _, n := range neighbors {
		if n.id != prev.id {
			if (n.row == sq.row && n.row == prev.row) || (n.col == sq.col && n.col == prev.col) {
				if n.val == sq.val {
					return n
				}
			}
		}
	}
	return nil
}

func initializeBoard() {
	boardMap = make(map[string]*square)
	number := 0
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			number++
			numberVal := strconv.Itoa(number)
			sq := square{row, col, numberVal, number}
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
