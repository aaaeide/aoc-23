package main

import (
	"os"

	"github.com/aaaeide/aoc-23/day01"
	"github.com/aaaeide/aoc-23/day02"
	"github.com/aaaeide/aoc-23/day03"
	"github.com/aaaeide/aoc-23/day04"
	"github.com/aaaeide/aoc-23/day05"
	"github.com/aaaeide/aoc-23/day06"
	"github.com/aaaeide/aoc-23/day07"
	"github.com/aaaeide/aoc-23/day08"
)

func runPart(i string, part func(*os.File) string) string {
	filename := "inputs/day0" + i + ".txt"
	file, err := os.Open(filename)
	
	if err != nil {
		return "\nCOULD NOT OPEN FILE " + filename
	}
	
	defer file.Close()

	return part(file)
}

func runDay(
	i string, 
	part1 func(*os.File) string, 
	part2 func(*os.File) string,
) {

	println("\n==================")
	println("Day 0" + i)
	println("------------------")
	println("Part 1:", runPart(i, part1))
	println("Part 2:", runPart(i, part2))
	println("==================")
} 

func main() {
	runDay("1", day01.Part1, day01.Part2)
	runDay("2", day02.Part1, day02.Part2)
	runDay("3", day03.Part1, day03.Part2)
	runDay("4", day04.Part1, day04.Part2)
	runDay("5", day05.Part1, day05.Part2)
	runDay("6", day06.Part1, day06.Part2)
	runDay("7", day07.Part1, day07.Part2)
	runDay("8", day08.Part1, day08.Part2)
}