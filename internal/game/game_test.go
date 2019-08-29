package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func place(squareID, char string) {
	s, _ := gameBoard.Square(squareID)
	s.Value = char
}
func TestIsWinnerEmpty(t *testing.T) {
	initializeBoard()

	assert.False(t, isWinner(playerChar), "should be no winner")
	assert.False(t, isWinner(compChar), "should be no winner")
}

func TestIsWinnerOneSquare(t *testing.T) {
	initializeBoard()

	place("1", playerChar)
	assert.False(t, isWinner(playerChar), "should be no winner")
	assert.False(t, isWinner(compChar), "should be no winner")
}

func TestIsWinnerTwoSquares(t *testing.T) {
	initializeBoard()

	place("1", playerChar)
	place("2", playerChar)

	assert.False(t, isWinner(playerChar), "should be no winner")
	assert.False(t, isWinner(compChar), "should be no winner")
}

func TestIsWinnerCorner(t *testing.T) {
	initializeBoard()

	place("1", playerChar)
	place("2", playerChar)
	place("4", playerChar)

	assert.False(t, isWinner(playerChar), "should be no winner")
	assert.False(t, isWinner(compChar), "should be no winner")
}

func TestIsWinnerTopRow(t *testing.T) {
	initializeBoard()

	place("1", playerChar)
	place("2", playerChar)
	place("3", playerChar)

	assert.True(t, isWinner(playerChar), "playerChar should be winner")
	assert.False(t, isWinner(compChar), "comp should not be winner")
}

func TestIsWinnerFirstCol(t *testing.T) {
	initializeBoard()

	place("1", playerChar)
	place("4", playerChar)
	place("7", playerChar)

	assert.True(t, isWinner(playerChar), "playerChar should be winner")
	assert.False(t, isWinner(compChar), "comp should not be winner")
}

func TestIsWinnerDiagonal(t *testing.T) {
	initializeBoard()

	place("1", playerChar)
	place("5", playerChar)
	place("9", playerChar)

	assert.True(t, isWinner(playerChar), "playerChar should be winner")
	assert.False(t, isWinner(compChar), "comp should not be winner")
}

func TestIsWinnerDiagonal2(t *testing.T) {
	initializeBoard()

	place("3", playerChar)
	place("5", playerChar)
	place("7", playerChar)

	assert.True(t, isWinner(playerChar), "playerChar should be winner")
	assert.False(t, isWinner(compChar), "comp should not be winner")
}

func TestIsWinnerBend(t *testing.T) {
	initializeBoard()

	place("1", playerChar)
	place("4", playerChar)
	place("8", playerChar)

	assert.False(t, isWinner(playerChar), "should be no winner")
	assert.False(t, isWinner(compChar), "should be no winner")
}

func TestIsWinnerBend2(t *testing.T) {
	initializeBoard()

	place("1", playerChar)
	place("5", playerChar)
	place("8", playerChar)

	assert.False(t, isWinner(playerChar), "should be no winner")
	assert.False(t, isWinner(compChar), "should be no winner")
}

func TestIsWinnerDifferntPlayersStraight(t *testing.T) {
	initializeBoard()

	place("1", playerChar)
	place("2", playerChar)
	place("3", compChar)

	assert.False(t, isWinner(playerChar), "playerChar should not be winner")
	assert.False(t, isWinner(compChar), "comp should not be winner")
}

func TestIsWinnerDifferntPlayersDiag(t *testing.T) {
	initializeBoard()

	place("1", playerChar)
	place("5", playerChar)
	place("9", compChar)

	assert.False(t, isWinner(playerChar), "playerChar should not be winner")
	assert.False(t, isWinner(compChar), "comp should not be winner")
}
