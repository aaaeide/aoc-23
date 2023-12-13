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
	"github.com/aaaeide/aoc-23/day09"
	"github.com/aaaeide/aoc-23/day10"
)

func runPart(i string, part func(*os.File) string) string {
	filename := "inputs/day" + i + ".txt"
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
	println("Day " + i)
	println("------------------")
	println("Part 1:", runPart(i, part1))
	println("Part 2:", runPart(i, part2))
	println("==================")
} 

func main() {
	runDay("01", day01.Part1, day01.Part2)
	runDay("02", day02.Part1, day02.Part2)
	runDay("03", day03.Part1, day03.Part2)
	runDay("04", day04.Part1, day04.Part2)
	runDay("05", day05.Part1, day05.Part2)
	runDay("06", day06.Part1, day06.Part2)
	runDay("07", day07.Part1, day07.Part2)
	runDay("08", day08.Part1, day08.Part2)
	runDay("09", day09.Part1, day09.Part2)
	runDay("10", day10.Part1, day10.Part2)
}