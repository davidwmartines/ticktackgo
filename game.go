package main

import (
	"fmt"
	"math"
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

var winner = false

func main() {

	initializeBoard()
	drawBoard()

	for !winner {
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
	went := false
	for _, sq := range boardMap {
		if sq.val != playerChar && sq.val != compChar {

			//TODO: see if sq could be a winning sequence

			// sq.val = compChar
			// went = true
			// break
		}
		if went {
			break
		}
	}
	if !went {
		computerGoFirst()
	}
}

func computerGoFirst() {
	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)
	for {
		randomNumber := gen.Intn((size*size)-1) + 1
		possibleSquare := boardMap[strconv.Itoa(randomNumber)]
		if possibleSquare.val != compChar && possibleSquare.val != playerChar {
			possibleSquare.val = compChar
			break
		}
	}
}

func checkWinner() {
	for _, sq := range boardMap {
		if sq.val == playerChar || sq.val == compChar {
			nex := getAdjacentMatch(sq)
			if nex != nil {
				last := getInlineMatch(sq, nex)
				if last != nil {
					fmt.Printf("%v WINS!\n", sq.val)
					os.Exit(0)
				}
			}
		}
	}

	checkTie()
}

func checkTie() {
	empty := false
	for _, sq := range boardMap {
		if sq.val != compChar && sq.val != playerChar {
			empty = true
			break
		}
	}
	if !empty {
		fmt.Println("no winner")
		os.Exit(0)
	}
}

func getAdjacentMatch(sq *square) *square {
	neighbors := getNeighbors(float64(sq.row), float64(sq.col))
	for _, n := range neighbors {
		if n.val == sq.val {
			return n
		}
	}
	return nil
}

func getInlineMatch(prev *square, sq *square) *square {
	neighbors := getNeighbors(float64(sq.row), float64(sq.col))
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

func getNeighbors(row float64, col float64) []*square {
	var output []*square
	limit := float64(size - 1)
	for x := math.Max(0, row-1); x <= math.Min(row+1, limit); x++ {
		for y := math.Max(0, col-1); y <= math.Min(col+1, limit); y++ {
			if x != row || y != col {
				sq := board[int(x)][int(y)]
				output = append(output, sq)
			}
		}
	}
	return output
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
