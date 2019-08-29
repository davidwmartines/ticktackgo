package grid

// Grid represents a 2D grid of points
type Grid [][]*Point

//Create creates a grid.
func Create(rows, cols int) (grid Grid) {
	grid = make([][]*Point, rows)
	for i := range grid {
		grid[i] = make([]*Point, cols)
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			point := Point{row, col}
			grid[row][col] = &point
		}
	}
	return
}
