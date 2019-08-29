package grid

import "math"

// Neighbors gets the adjacent points to a point.
func (source Grid) Neighbors(point *Point) (output []*Point) {
	row := float64(point.Row)
	col := float64(point.Col)
	xLimit := float64(len(source) - 1)
	for x := math.Max(0, row-1); x <= math.Min(row+1, xLimit); x++ {
		yLimit := float64(len(source[int(x)]) - 1)
		for y := math.Max(0, col-1); y <= math.Min(col+1, yLimit); y++ {
			if x != row || y != col {
				output = append(output, source[int(x)][int(y)])
			}
		}
	}
	return
}
