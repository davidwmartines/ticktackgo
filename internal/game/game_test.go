package game

import (
	"testing"
)

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
