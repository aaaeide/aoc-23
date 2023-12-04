package main

import (
	"github.com/aaaeide/aoc-23/day01"
	"github.com/aaaeide/aoc-23/day02"
)

func runDay(
	i string, 
	part1 func(string) string, 
	part2 func(string) string,
) {
	filename := "inputs/day0" + i + ".txt"

	println("\n==================\nDay 0" + i + "\n------------------")
	println("Part 1:", part1(filename))
	println("Part 2:", part2(filename))
	println("==================")
} 

func main() {
	runDay("1", day01.Part1, day01.Part2)
	runDay("2", day02.Part1, day02.Part2)
}