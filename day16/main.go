package day16

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseGrid(file *os.File) (grid []string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}

type BeamTip struct {
	x, y int;
	dir string;
}

func (bt BeamTip) toKey() string {
	return fmt.Sprintf("%d.%d.%s", bt.x, bt.y, bt.dir)
}

func coordKey(x, y int) string {
	return fmt.Sprintf("%d.%d", x, y)
}

func countEnergizedTiles(grid []string, initialBeamTip BeamTip) int {
	HEIGHT, WIDTH := len(grid), len(grid[0])
	beam := make(map[string]bool)
	energizedTiles := make(map[string]bool)
	tips := []BeamTip{initialBeamTip}
	
	for len(tips) > 0{
		tip := tips[0]
		tips = tips[1:]
		if _, ok := beam[tip.toKey()]; ok ||
			tip.x < 0 || tip.x >= WIDTH || tip.y < 0 || tip.y >= HEIGHT {
			continue
		}
	
		beam[tip.toKey()] = true
		energizedTiles[coordKey(tip.x, tip.y)] = true
	
		switch tip.dir {
		case "RIGHT":
			switch grid[tip.y][tip.x] {
			case '.':
				tips = append(tips, BeamTip{tip.x+1, tip.y, "RIGHT"})
			case '/':
				tips = append(tips, BeamTip{tip.x, tip.y-1, "UP"})
			case '\\':
				tips = append(tips, BeamTip{tip.x, tip.y+1, "DOWN"})
			case '-':
				tips = append(tips, BeamTip{tip.x+1, tip.y, "RIGHT"})
			case '|':
				tips = append(tips, BeamTip{tip.x, tip.y-1, "UP"})
				tips = append(tips, BeamTip{tip.x, tip.y+1, "DOWN"})
			}
		case "LEFT":
			switch grid[tip.y][tip.x] {
			case '.':
				tips = append(tips, BeamTip{tip.x-1, tip.y, "LEFT"})
			case '/':
				tips = append(tips, BeamTip{tip.x, tip.y+1, "DOWN"})
			case '\\':
				tips = append(tips, BeamTip{tip.x, tip.y-1, "UP"})
			case '-':
				tips = append(tips, BeamTip{tip.x-1, tip.y, "LEFT"})
			case '|':
				tips = append(tips, BeamTip{tip.x, tip.y-1, "UP"})
				tips = append(tips, BeamTip{tip.x, tip.y+1, "DOWN"})
			}
		case "UP":
			switch grid[tip.y][tip.x] {
			case '.':
				tips = append(tips, BeamTip{tip.x, tip.y-1, "UP"})
			case '/':
				tips = append(tips, BeamTip{tip.x+1, tip.y, "RIGHT"})
			case '\\':
				tips = append(tips, BeamTip{tip.x-1, tip.y, "LEFT"})
			case '-':
				tips = append(tips, BeamTip{tip.x-1, tip.y, "LEFT"})
				tips = append(tips, BeamTip{tip.x+1, tip.y, "RIGHT"})
			case '|':
				tips = append(tips, BeamTip{tip.x, tip.y-1, "UP"})
			}
		case "DOWN":
			switch grid[tip.y][tip.x] {
			case '.':
				tips = append(tips, BeamTip{tip.x, tip.y+1, "DOWN"})
			case '/':
				tips = append(tips, BeamTip{tip.x-1, tip.y, "LEFT"})
			case '\\':
				tips = append(tips, BeamTip{tip.x+1, tip.y, "RIGHT"})
			case '-':
				tips = append(tips, BeamTip{tip.x-1, tip.y, "LEFT"})
				tips = append(tips, BeamTip{tip.x+1, tip.y, "RIGHT"})
			case '|':
				tips = append(tips, BeamTip{tip.x, tip.y+1, "DOWN"})
			}
		}
	}

	return len(energizedTiles)
}

func Part1(file *os.File) string {
	grid := parseGrid(file)
	energizedTiles := countEnergizedTiles(grid, BeamTip{0, 0, "RIGHT"})

	return strconv.Itoa(energizedTiles)
}

func Part2(file *os.File) string {
	grid := parseGrid(file)
	biggest := 0
	
	for y:=0; y<len(grid); y++ {
		fromLeft := BeamTip{0, 0, "RIGHT"}
		if cnt := countEnergizedTiles(grid, fromLeft); cnt > biggest {
			biggest = cnt
		}

		fromRight := BeamTip{len(grid[0])-1, 0, "LEFT"}
		if cnt := countEnergizedTiles(grid, fromRight); cnt > biggest {
			biggest = cnt
		}
	}

	for x:=0; x<len(grid[0]); x++ {
		fromTop := BeamTip{x, 0, "DOWN"}
		if cnt := countEnergizedTiles(grid, fromTop); cnt > biggest {
			biggest = cnt
		}

		fromBottom := BeamTip{x, len(grid)-1, "UP"}
		if cnt := countEnergizedTiles(grid, fromBottom); cnt > biggest {
			biggest = cnt
		}
	}

	return strconv.Itoa(biggest)
}