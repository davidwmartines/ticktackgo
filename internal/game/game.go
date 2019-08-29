package game

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/davidwmartines/ticktackgo/internal/board"
)

var gameBoard board.Board

const size int = 3

const playerChar string = "X"
const compChar string = "O"

//Start begins a new game.
func Start() {

	initializeBoard()

	gameBoard.Draw()

	for {
		playerGo()
		gameBoard.Draw()
		checkWinner()

		computerGo()
		gameBoard.Draw()
		checkWinner()
	}

}

func initializeBoard() {
	gameBoard = board.New(3)
}

func playerGo() {
	var input string
	fmt.Println("pick a square: ")
	fmt.Scanln(&input)
	if sq, valid := gameBoard.Square(input); valid {
		if !sq.IsEmpty() {
			fmt.Printf("%v is already taken!\n", input)
			playerGo()
		} else {
			sq.Value = playerChar
		}
	} else {
		fmt.Printf("%v is not a valid square!\n", input)
		playerGo()
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

func computerGo() {
	fmt.Println("computer turn...")
	choice := getNextSquareForComputer()
	choice.Value = compChar
}

func getNextSquareForComputer() *board.Square {

	for _, sq := range gameBoard.Squares() {
		// if square is owned, try getting a neighbor
		if sq.Value == compChar {
			for _, n := range gameBoard.Neighbors(sq) {
				if n.IsEmpty() {
					return n
				}
			}
		}
	}
	return getRandomEmptySquare()
}

func getRandomEmptySquare() *board.Square {
	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)
	for {
		randomNumber := gen.Intn((size*size)-1) + 1
		possibleSquare, _ := gameBoard.Square(strconv.Itoa(randomNumber))
		if possibleSquare.IsEmpty() {
			return possibleSquare
		}
	}
}

func isWinner(char string) bool {
	for _, sq := range gameBoard.Squares() {
		if sq.Value == char {
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
	for _, sq := range gameBoard.Squares() {
		if sq.IsEmpty() {
			return false
		}
	}
	return true
}

func getAdjacentMatch(sq *board.Square) *board.Square {
	for _, n := range gameBoard.Neighbors(sq) {
		if n.Value == sq.Value {
			return n
		}
	}
	return nil
}

func getInlineMatch(prev *board.Square, sq *board.Square) *board.Square {
	for _, n := range gameBoard.Neighbors(sq) {
		if n.ID != prev.ID {
			if (n.RowMatch(sq) && n.RowMatch(prev)) || (n.ColMatch(sq) && n.ColMatch(prev)) {
				if n.Value == sq.Value {
					return n
				}
			}
		}
	}
	return nil
}
