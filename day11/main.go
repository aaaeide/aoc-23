package day11

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Galaxy struct {
	x, y int;
}

func parseUniverse(file *os.File) []Galaxy {
	scanner := bufio.NewScanner(file)

	y := 0
	var galaxies []Galaxy
	for scanner.Scan() {
		for x, tile := range strings.Split(scanner.Text(), "") {
			if tile == "#" {
				galaxies = append(galaxies, Galaxy{x,y})
			}
		}

		y++
	}

	return galaxies
}

// c is the constant to _add_ to each row/column expanded
func expandUniverse(galaxies []Galaxy, c int) []Galaxy {
	xs, ys := make(map[int]int, 0), make(map[int]int, 0)
	
	// sort by x-coord
	sort.Slice(galaxies, func(i, j int) bool {
		return galaxies[i].x < galaxies[j].x
	})
	
	var xExpandedGalaxies []Galaxy
	ix := 0
	for _, g := range galaxies {
		if _, ok := xs[g.x]; !ok {
			xs[g.x] = ix
			ix++
		}

		xExpandedGalaxies = append(xExpandedGalaxies, Galaxy{
			x: g.x + (g.x-xs[g.x])*c,
			y: g.y,
		})
	}
	
	// sort by y-coord
	sort.Slice(xExpandedGalaxies, func(i, j int) bool {
		return xExpandedGalaxies[i].y < xExpandedGalaxies[j].y
	})
	
	var xyExpandedGalaxies []Galaxy
	iy := 0
	for _, g := range xExpandedGalaxies {
		if _, ok := ys[g.y]; !ok {
			ys[g.y] = iy
			iy++
		}

		xyExpandedGalaxies = append(xyExpandedGalaxies, Galaxy{
			x: g.x,
			y: g.y + (g.y-ys[g.y])*c,
		})
	}

	return xyExpandedGalaxies
}

func solve(galaxies []Galaxy) int {
	sum := 0
	for len(galaxies) > 0 {
		galaxyA := galaxies[0]

		for b:=1; b<len(galaxies); b++ {
			galaxyB := galaxies[b]

			sum += int(math.Abs(float64(galaxyA.x) - float64(galaxyB.x)) + 
				math.Abs(float64(galaxyA.y) - float64(galaxyB.y)))
		}

		galaxies = galaxies[1:]
	}

	return sum
}

func Part1(file *os.File) string {
	println("part 1")
	galaxies := parseUniverse(file)
	// c = 1, adding one to each empty row doubles it.
	galaxies = expandUniverse(galaxies, 1)
	
	return strconv.Itoa(solve(galaxies))
}

func Part2Testable(file *os.File, expansionFactor int) string {
	println("\npart 2")
	galaxies := parseUniverse(file)
	// adding ef-1 to each emnpty row is the same as multiplying it by ef.
	galaxies = expandUniverse(galaxies, expansionFactor-1)
	
	return strconv.Itoa(solve(galaxies))
}

func Part2(file *os.File) string {
	galaxies := parseUniverse(file)	
	galaxies = expandUniverse(galaxies, 1000000-1)
	
	return strconv.Itoa(solve(galaxies))
}