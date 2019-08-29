package game

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/davidwmartines/ticktackgo/internal/board"
	"github.com/davidwmartines/ticktackgo/internal/constants"
)

var gameBoard board.Board

//Start begins a new game.
func Start() {

	initializeBoard()

	render(&gameBoard)

	for {
		playerGo()
		render(&gameBoard)
		checkWinner()

		computerGo()
		render(&gameBoard)
		checkWinner()
	}

}

func initializeBoard() {
	gameBoard = board.New(constants.Size)
}

func playerGo() {
	var input string
	fmt.Print("Your turn. Pick a square by entering the number: ")
	fmt.Scanln(&input)
	if sq, valid := gameBoard.Square(input); valid {
		if !sq.IsEmpty() {
			fmt.Printf("%v is already taken!\n", input)
			playerGo()
		} else {
			sq.Value = constants.PlayerChar
		}
	} else {
		fmt.Printf("%v is not a valid square!\n", input)
		playerGo()
	}
}

func checkWinner() {
	if isWinner(constants.PlayerChar) {
		fmt.Println("You Win!")
		os.Exit(0)
	} else if isWinner(constants.CompChar) {
		fmt.Println("Computer Wins!")
		os.Exit(0)
	} else if isTie() {
		fmt.Println("no winner")
		os.Exit(0)
	}
}

func computerGo() {
	fmt.Println("My turn...")
	time.Sleep(time.Second * constants.ThinkTime)
	choice := getNextSquareForComputer()
	choice.Value = constants.CompChar
}

func getNextSquareForComputer() *board.Square {

	var next *board.Square

	//first, block opponent from winning
	//TODO

	//next see if I can win in this move
	//TODO

	//next just try to start a run from an occupied square
	next = getNextFromOccupied()
	if next != nil {
		return next
	}

	//lastly, if this is the first turn, pick at random
	return getRandomEmptySquare()
}

func getNextFromOccupied() *board.Square {
	// if square is occupied, try getting a neighbor
	for _, sq := range gameBoard.Squares() {
		if sq.Value == constants.CompChar {
			for _, n := range gameBoard.Neighbors(sq) {
				if n.IsEmpty() {
					return n
				}
			}
		}
	}
	return nil
}

func getRandomEmptySquare() *board.Square {
	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)
	for {
		randomNumber := gen.Intn((constants.Size*constants.Size)-1) + 1
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
				lineMatch := getInlineMatch(sq, nex)
				if lineMatch != nil {
					return true
				}
				if !sq.RowMatch(nex) && !sq.ColMatch(nex) {
					diagMatch := getDiagonalMatch(sq, nex)
					if diagMatch != nil {
						return true
					}
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

func getDiagonalMatch(prev *board.Square, sq *board.Square) *board.Square {
	for _, n := range gameBoard.Neighbors(sq) {
		if n.ID != prev.ID {
			if !n.RowMatch(prev) && !n.RowMatch(sq) && !n.ColMatch(prev) && !n.ColMatch(sq) {
				if n.Value == sq.Value {
					return n
				}
			}
		}
	}
	return nil
}
