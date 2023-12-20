package day15

import (
	"testing"

	"github.com/aaaeide/aoc-23/utils"
)

func TestHASH(t *testing.T) {
	if HASH("HASH") != 52 {
		t.Error("HASH of HASH should be 52!!")
	}
}

func TestPart1(t *testing.T) {
	utils.Tester(t, "test.txt", Part1, "1320")()
}

func TestPart2(t *testing.T) {
	utils.Tester(t, "test.txt", Part2, "145")()
}