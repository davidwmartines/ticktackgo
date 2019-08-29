package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSquare(t *testing.T) {
	grid := Create(3, 3)
	assert.True(t, len(grid) == 3, "grid should have 3 rows")
	assert.True(t, len(grid[0]) == 3, "grid col 0 should have 3 columns")
}

func TestCreateRect(t *testing.T) {
	grid := Create(5, 7)
	assert.True(t, len(grid) == 5, "grid should have 5 rows")
	assert.True(t, len(grid[0]) == 7, "grid col 0 should have 7 columns")
}
