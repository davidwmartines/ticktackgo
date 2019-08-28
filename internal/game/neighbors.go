package game

import "math"

func neighbors(row, col float64) (output []*square) {
	limit := float64(size - 1)
	for x := math.Max(0, row-1); x <= math.Min(row+1, limit); x++ {
		for y := math.Max(0, col-1); y <= math.Min(col+1, limit); y++ {
			if x != row || y != col {
				sq := board[int(x)][int(y)]
				output = append(output, sq)
			}
		}
	}
	return
}
