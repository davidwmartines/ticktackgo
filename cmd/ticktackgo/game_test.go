package main

import (
	"testing"
)

func TestGetNeighborsCount(t *testing.T) {

	res := getNeighbors(0, 0)
	if len(res) != 3 {
		t.Errorf("0, 0 should return 3 neighbors")
	}

	res1 := getNeighbors(0, 1)
	if len(res1) != 5 {
		t.Errorf("0, 1 should return 5 neighbors")
	}

	res2 := getNeighbors(1, 1)
	if len(res2) != 8 {
		t.Errorf("1, 1 should return 8 neighbors")
	}

	res3 := getNeighbors(1, 2)
	if len(res3) != 5 {
		t.Errorf("1, 2 should return 5 neighbors")
	}

	res4 := getNeighbors(2, 2)
	if len(res4) != 3 {
		t.Errorf("2, 2 should return 3 neighbors")
	}

}

func TestIsWinnerEmpty(t *testing.T) {
	initializeBoard()
	if isWinner(playerChar) {
		t.Errorf("should be no winner")
	}
	if isWinner(compChar) {
		t.Errorf("should be no winner")
	}
}

func TestIsWinnerOneSquare(t *testing.T) {
	initializeBoard()

	boardMap["1"].val = playerChar

	if isWinner(playerChar) {
		t.Errorf("should be no winner")
	}
	if isWinner(compChar) {
		t.Errorf("should be no winner")
	}
}

func TestIsWinnerTwoSquares(t *testing.T) {
	initializeBoard()

	boardMap["1"].val = playerChar
	boardMap["2"].val = playerChar

	if isWinner(playerChar) {
		t.Errorf("should be no winner")
	}
	if isWinner(compChar) {
		t.Errorf("should be no winner")
	}
}

func TestIsWinnerThreeNotInRow(t *testing.T) {
	initializeBoard()

	boardMap["1"].val = playerChar
	boardMap["2"].val = playerChar
	boardMap["4"].val = playerChar

	if isWinner(playerChar) {
		t.Errorf("should be no winner")
	}
	if isWinner(compChar) {
		t.Errorf("should be no winner")
	}
}

func TestIsWinnerTopRow(t *testing.T) {
	initializeBoard()

	boardMap["1"].val = playerChar
	boardMap["2"].val = playerChar
	boardMap["3"].val = playerChar

	if !isWinner(playerChar) {
		t.Errorf("player should be winner")
	}
	if isWinner(compChar) {
		t.Errorf("comp should not be winner")
	}
}

func TestIsWinnerFirstCol(t *testing.T) {
	initializeBoard()

	boardMap["1"].val = playerChar
	boardMap["4"].val = playerChar
	boardMap["7"].val = playerChar

	if !isWinner(playerChar) {
		t.Errorf("player should be winner")
	}
	if isWinner(compChar) {
		t.Errorf("comp should not be winner")
	}
}

func TestIsWinnerDiagonal(t *testing.T) {
	initializeBoard()

	boardMap["1"].val = playerChar
	boardMap["5"].val = playerChar
	boardMap["9"].val = playerChar

	if !isWinner(playerChar) {
		t.Errorf("player should be winner")
	}
	if isWinner(compChar) {
		t.Errorf("comp should not be winner")
	}
}
