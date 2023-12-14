package utils

func Transpose[T any](grid [][]T) [][]T {
	H := len(grid)
	W := len(grid[0])

	var new [][]T
	for x:=0; x<W; x++ {
		new = append(new, make([]T, H))
	}

	for y:=0; y<H; y++ {
		for x:=0; x<W; x++ {
			new[x][y] = grid[y][x]
		}
	}

	return new
}