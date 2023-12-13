package day10

import (
	"testing"

	"github.com/aaaeide/aoc-23/utils"
)

func TestPart1(t *testing.T) {
	utils.Tester(t, "test1.txt", Part1, "4")()
	utils.Tester(t, "test2.txt", Part1, "8")()
}

func TestPart2(t *testing.T) {
	utils.Tester(t, "test1.txt", Part2, "1")()
	utils.Tester(t, "test2.txt", Part2, "1")()
	utils.Tester(t, "test3.txt", Part2, "4")()
	utils.Tester(t, "test4.txt", Part2, "4")()
	utils.Tester(t, "test5.txt", Part2, "8")()
}