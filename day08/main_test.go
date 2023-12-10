package day08

import (
	"testing"

	"github.com/aaaeide/aoc-23/utils"
)

func TestPart1(t *testing.T) {
	utils.Tester(t, "test1.txt", Part1, "2")()
	utils.Tester(t, "test2.txt", Part1, "6")()
}

func TestPart2(t *testing.T) {
	utils.Tester(t, "test3.txt", Part2, "6")()
}