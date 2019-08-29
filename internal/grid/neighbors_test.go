package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNeighbors00(t *testing.T) {
	grid := Create(3, 3)
	neighbors := grid.Neighbors(&Point{0, 0})
	assert.True(t, len(neighbors) == 3, "0, 0 should return 3 neighbors")
	assert.True(t, neighbors[0].Row == 0, "neighbor 0 should be in row 0")
	assert.True(t, neighbors[0].Col == 1, "neighbor 0 should be in col 1")
	assert.True(t, neighbors[1].Row == 1, "neighbor 1 should be in row 1")
	assert.True(t, neighbors[1].Col == 0, "neighbor 1 should be in col 0")
	assert.True(t, neighbors[2].Row == 1, "neighbor 2 should be in row 1")
	assert.True(t, neighbors[2].Col == 1, "neighbor 2 should be in col 1")
}

func TestNeighbors22(t *testing.T) {
	grid := Create(3, 3)
	neighbors := grid.Neighbors(&Point{2, 2})
	assert.True(t, len(neighbors) == 3, "2, 2 should return 3 neighbors")
	assert.True(t, neighbors[0].Row == 1, "neighbor 0 should be in row 1")
	assert.True(t, neighbors[0].Col == 1, "neighbor 0 should be in col 1")
	assert.True(t, neighbors[1].Row == 1, "neighbor 1 should be in row 1")
	assert.True(t, neighbors[1].Col == 2, "neighbor 1 should be in col 2")
	assert.True(t, neighbors[2].Row == 2, "neighbor 2 should be in row 2")
	assert.True(t, neighbors[2].Col == 1, "neighbor 2 should be in col 1")
}

func TestNeighbors01(t *testing.T) {
	grid := Create(3, 3)
	neighbors := grid.Neighbors(&Point{0, 1})
	assert.True(t, len(neighbors) == 5, "0, 1 should return 5 neighbors")
}

func TestNeighbors11(t *testing.T) {
	grid := Create(3, 3)
	neighbors := grid.Neighbors(&Point{1, 1})
	assert.True(t, len(neighbors) == 8, "1, 1 should return 8 neighbors")
}

func TestNeighbors12(t *testing.T) {
	grid := Create(3, 3)
	neighbors := grid.Neighbors(&Point{1, 2})
	assert.True(t, len(neighbors) == 5, "1, 2 should return 5 neighbors")
}
