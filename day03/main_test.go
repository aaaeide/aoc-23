package day03

import (
	"testing"

	"github.com/aaaeide/aoc-23/utils"
)

func TestPart1(t *testing.T) {
	utils.Tester(t, "test1.txt", Part1, "4361")()
}

func TestPart2(t *testing.T) {
	utils.Tester(t, "test1.txt", Part2, "467835")()
}