package day07

import (
	"testing"

	"github.com/aaaeide/aoc-23/utils"
)

func TestPart1(t *testing.T) {
	utils.Tester(t, "test.txt", Part1, "6440")()
}

func TestPart2(t *testing.T) {
	utils.Tester(t, "test.txt", Part2, "5905")()
}