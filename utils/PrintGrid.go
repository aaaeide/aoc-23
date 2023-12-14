package utils

import "fmt"

func SPrintGrid(grid [][]string) string {
	s := ""
	for y:=0; y<len(grid); y++ {
		s += fmt.Sprintf("%v\n", grid[y])
	}
	return s
}