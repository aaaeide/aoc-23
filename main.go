package main

import (
	"github.com/aaaeide/aoc-23/day01"
	"github.com/aaaeide/aoc-23/day02"
)

func runDay(
	i string, 
	part1 func() string, 
	part2 func() string,
) {
	println("\n========\nDay 0" + i)
	println("Part 1:", part1())
	println("Part 2:", part2())
} 

func main() {
	runDay("1", day01.Part1, day01.Part2)
	runDay("2", day02.Part1, day02.Part2)
}