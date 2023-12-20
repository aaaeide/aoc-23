package day12

import (
	"testing"

	"github.com/aaaeide/aoc-23/utils"
)

func TestCheckLine(t *testing.T) {
	line := Line{
		"#.#.###",
		[]int{1,1,3},
	}

	if !checkLine(line) {
		t.Errorf("%s should be valid", line.str)
	}
}

func TestCountValidArrangements1(t *testing.T) {
	line := Line{
		"???.###",
		[]int{1,1,3},
	}

	got := countValidArrangements(line, 0, false)
	want := 1
	if got != want {
		t.Errorf("%s %v should have %d arrangements, got %d", line.str, line.groupSizes, want, got)
		return
	}
}

func TestCountValidArrangements2(t *testing.T) {
	line := Line{
		".??..??...?##.",
		[]int{1,1,3},
	}

	got := countValidArrangements(line, 0, false)
	want := 4
	if got != want {
		t.Errorf("%s %v should have %d arrangements, got %d", line.str, line.groupSizes, want, got)
		return
	}
}

func TestCountValidArrangements3(t *testing.T) {
	line := Line{
		"?#?#?#?#?#?#?#?",
		[]int{1,3,1,6},
	}

	got := countValidArrangements(line, 0, false)
	want := 1
	if got != want {
		t.Errorf("%s %v should have %d arrangements, got %d", line.str, line.groupSizes, want, got)
		return
	}
}

func TestCountValidArrangements4(t *testing.T) {
	line := Line{
		"????.#...#...",
		[]int{4,1,1},
	}

	got := countValidArrangements(line, 0, false)
	want := 1
	if got != want {
		t.Errorf("%s %v should have %d arrangements, got %d", line.str, line.groupSizes, want, got)
		return
	}
}

func TestCountValidArrangements5(t *testing.T) {
	line := Line{
		"????.######..#####.",
		[]int{1,6,5},
	}

	got := countValidArrangements(line, 0, false)
	want := 4
	if got != want {
		t.Errorf("%s %v should have %d arrangements, got %d", line.str, line.groupSizes, want, got)
		return
	}
}

func TestCountValidArrangements6(t *testing.T) {
	line := Line{
		"?###????????",
		[]int{3,2,1},
	}

	got := countValidArrangements(line, 0, false)
	want := 10
	if got != want {
		t.Errorf("%s %v should have %d arrangements, got %d", line.str, line.groupSizes, want, got)
		return
	}
}

func TestPart1(t *testing.T) {
	utils.Tester(t, "test.txt", Part1, "21")()
}

// func TestPart2(t *testing.T) {
// 	utils.Tester(t, "test.txt", Part2, "525152")()
// }