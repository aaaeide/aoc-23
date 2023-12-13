package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aaaeide/aoc-23/utils"
)

func generateGrid(file *os.File) ([][]string, []int) {
	scanner := bufio.NewScanner(file)
	
	// generate grid
	var grid [][]string
	y := 0
	var s []int
	for scanner.Scan() {
		grid = append(grid, make([]string, 0))

		for x, tile := range strings.Split(scanner.Text(), "") {
			grid[y] = append(grid[y], tile)
			
			if tile == "S" {
				s = []int{x, y}
			}
		}

		y++
	}

	return grid, s
}

func nextPipeStep(pipe string, xy []int, xy_prev []int) []int {
	switch pipe {
	case "|":
		if xy_prev[1] < xy[1] {
			// going south
			return []int{xy[0], xy[1]+1}
		}
		// going north
		return []int{xy[0], xy[1]-1}
	case "-":
		if xy_prev[0] < xy[0] {
			// going east
			return []int{xy[0]+1, xy[1]}
		}
		// going west
		return []int{xy[0]-1, xy[1]}
	case "L":
		if xy_prev[1] < xy[1] {
			// going east
			return []int{xy[0]+1, xy[1]}
		}
		// going north
		return []int{xy[0], xy[1]-1}
	case "J":
		if xy_prev[1] < xy[1] {
			// going west
			return []int{xy[0]-1, xy[1]}
		}
		// going north
		return []int{xy[0], xy[1]-1}
	case "7":
		if xy_prev[0] < xy[0] {
			// going south
			return []int{xy[0], xy[1]+1}
		}
		// going west
		return []int{xy[0]-1, xy[1]}
	case "F":
		if xy_prev[0] > xy[0] {
			// going south
			return []int{xy[0], xy[1]+1}
		}
		// going east
		return []int{xy[0]+1, xy[1]}
	default:
		panic("didn't understand " + pipe + " !!")
	}
} 

func findPipe(grid [][]string, s []int) ([][]string, map[string]int, string) {
	sx, sy := s[0], s[1]

	HEIGHT := len(grid)
	WIDTH := len(grid[0])

	// find starting pipe ends (two)
	var pipeEnds [][]int
	startPipe := []string{"|", "-", "L", "J", "7", "F"}
	
	// NORTH
	if sy > 0 && 
		(grid[sy-1][sx] == "|" || 
		grid[sy-1][sx] == "7" || 
		grid[sy-1][sx] == "F") {
			pipeEnds = append(pipeEnds, []int{sx, sy-1})
			startPipe = utils.Intersection(startPipe, []string{"|", "L", "J"})
	}
	
	// EAST
	if sx < WIDTH-1 && 
		(grid[sy][sx+1] == "-" || 
		grid[sy][sx+1] == "J" || 
		grid[sy][sx+1] == "7") {
			pipeEnds = append(pipeEnds, []int{sx+1, sy})
			startPipe = utils.Intersection(startPipe, []string{"-", "L", "F"})
	}
	
	// SOUTH
	if sy < HEIGHT-1 && 
		(grid[sy+1][sx] == "|" || 
		grid[sy+1][sx] == "J" || 
		grid[sy+1][sx] == "L") {
			pipeEnds = append(pipeEnds, []int{sx, sy+1})
			startPipe = utils.Intersection(startPipe, []string{"|", "7", "F"})
	}
	
	// WEST
	if sx > 0 && 
		(grid[sy][sx-1] == "-" || 
		grid[sy][sx-1] == "F" || 
		grid[sy][sx-1] == "L") {
			pipeEnds = append(pipeEnds, []int{sx-1, sy})
			startPipe = utils.Intersection(startPipe, []string{"-", "7", "J"})
	}

	if len(startPipe) != 1 {
		fmt.Println(startPipe)
		panic("COULD NOT DETERMINE STARTING PIPE SHAPE")
	}

	grid[sy][sx] = startPipe[0]

	// traverse pipe from each end until meet
	a_prev, b_prev := []int{sx, sy}, []int{sx, sy}
	a, b := pipeEnds[0], pipeEnds[1]
	distance := map[string]int{
		toKey([]int{sx, sy}): 0,
		toKey(a): 1,
		toKey(b): 1,
	}

	d := 2
	for {
		a_next := nextPipeStep(grid[a[1]][a[0]], a, a_prev)
		b_next := nextPipeStep(grid[b[1]][b[0]], b, b_prev)

		// fmt.Println(a_next, grid[a_next[1]][a_next[0]])
		// fmt.Println(b_next, grid[b_next[1]][b_next[0]])
		// println()
		
		distance[toKey(a_next)] = d
		distance[toKey(b_next)] = d

		if (a_next[0] == b_next[0]  && a_next[1] == b_next[1]) ||
			(a_next[0] == b[0] && a_next[1] == b[1]) {
				// the pipes have met
				return grid, distance, strconv.Itoa(d)
		}

		a_prev = a
		a = a_next
		
		b_prev = b
		b = b_next

		d += 1
	}
}

func toKey(xy []int) string {
	return fmt.Sprintf("%d,%d", xy[0], xy[1])
}

func Part1(file *os.File) string {
	grid, s := generateGrid(file)
	_, _, maxDst := findPipe(grid, s)
	return maxDst
}

func Part2(file *os.File) string {
	grid, s := generateGrid(file)
	grid, pipe, _ := findPipe(grid, s)

	// generate new grid with only our main pipe
	HEIGHT := len(grid)
	WIDTH := len(grid[0])

	numInside := 0

	for y:=0; y<HEIGHT; y++ {
		insideLoop := false

		// "up" or "down"
		var onPipeHeading string
		
		for x:=0; x<WIDTH; x++ {
			_, ok := pipe[toKey([]int{x, y})]
			if ok {
				// found pipe
				if grid[y][x] == "|" {
					// hit the pipe from the side
					insideLoop = !insideLoop
					continue
				}

				if grid[y][x] == "L" {
					// hit the pipe going down at a corner
					onPipeHeading = "down"
					continue
				}

				if grid[y][x] == "F" {
					// hit the pipe going up at a corner
					onPipeHeading = "up"
					continue
				}

				if grid[y][x] == "-" {
					// continuing along pipe
					continue
				}

				if grid[y][x] == "7" {
					if onPipeHeading == "down" {
						insideLoop = !insideLoop
					}
					continue
				}

				if grid[y][x] == "J" {
					if onPipeHeading == "up" {
						insideLoop = !insideLoop
					}
				}


				if grid[y][x] == "|" ||
					grid[y][x] == "L" ||
					grid[y][x] == "F" {
					// we just hit a pipe from the side
					insideLoop = !insideLoop
				}
				
				continue
			}
			
			if insideLoop {
				numInside += 1
			}
		}
	}


	return strconv.Itoa(numInside)
}