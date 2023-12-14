package day11

import (
	"os"
	"testing"

	"github.com/aaaeide/aoc-23/utils"
)

func TestTranspose(t *testing.T) {
	grid := [][]string{
		{"1",  "2",  "3"},
		{"4",  "5",  "6"},
		{"7",  "8",  "9"},
		{"10", "11", "12"},
	}

	got := utils.Transpose(grid)
	want := [][]string{
		{"1", "4", "7", "10"},
		{"2", "5", "8", "11"},
		{"3", "6", "9", "12"},
	}

	for y:=0; y<4; y++ {
		for x:=0; x<3; x++ {
			if got[x][y] != want[x][y] {
				t.Errorf("got\n%s wanted\n%s", utils.SPrintGrid(got), utils.SPrintGrid(want))
			}
		}
	}
}

func TestPart1(t *testing.T) {
	utils.Tester(t, "test1.txt", Part1, "374")()
}

func TestPart2_1(t *testing.T) {
	utils.Tester(t, "test1.txt", func(f *os.File) string {
		return Part2Testable(f, 10)	
	}, "1030")()
}

// func TestPart2_2(t *testing.T) {
// 	utils.Tester(t, "test1.txt", func(f *os.File) string {
// 		return Part2Testable(f, 100)	
// 	}, "8410")()
// }