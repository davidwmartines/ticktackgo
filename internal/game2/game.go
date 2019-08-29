package game2

import (
	"fmt"
	"math/rand"
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

	gameBoard = board.New(3)

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

func playerGo() {
	var input string
	fmt.Println("pick a square: ")
	fmt.Scanln(&input)
	if sq, valid := gameBoard.Square(input); valid {
		if !sq.IsEmpty() {
			fmt.Printf("%v is already taken!\n", input)
			playerGo()
		}
		sq.Value = playerChar
	} else {
		fmt.Printf("%v is not a valid square!\n", input)
		playerGo()
	}
}

func checkWinner() {

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
