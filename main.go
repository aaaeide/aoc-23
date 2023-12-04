package main

import (
	"os"

	"github.com/aaaeide/aoc-23/day01"
	"github.com/aaaeide/aoc-23/day02"
)

func runDay(
	i string, 
	part1 func(*os.File) string, 
	part2 func(*os.File) string,
) {
	filename := "inputs/day0" + i + ".txt"
	file, err := os.Open(filename)
	
	if err != nil {
		println("\nCOULD NOT OPEN FILE", filename)
		return
	}
	
	defer file.Close()

	println("\n==================\nDay 0" + i + "\n------------------")
	println("Part 1:", part1(file))
	println("Part 2:", part2(file))
	println("==================")
} 

func main() {
	runDay("1", day01.Part1, day01.Part2)
	runDay("2", day02.Part1, day02.Part2)
}