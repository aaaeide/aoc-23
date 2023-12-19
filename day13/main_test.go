package day13

import (
	"os"
	"testing"

	"github.com/aaaeide/aoc-23/utils"
)

func TestValidateMirror(t *testing.T) {
	file, _ := os.Open("test_horiz.txt")
	ptn := parsePatterns(file)[0]

	if validateMirror(ptn, 2, 0) {
		t.Error("2 is not a valid mirror location!")
	}

	if !validateMirror(ptn, 4, 0) {
		t.Error("5 should be a valid mirror location!")
	}
}

func TestValidateSmudgedMirror(t *testing.T) {
	file, _ := os.Open("test.txt")

	ptns := parsePatterns(file)
	ptn1, ptn2 := ptns[0], ptns[1]

	if !validateMirror(ptn1, 3, 1) {
		t.Error("ptn1: 3 should be valid mirror location (1 smudge)!")
	}

	if validateMirror(ptn1, 5, 1) {
		t.Error("ptn1: 5 should not be a valid mirror location (1 smudge)!")
	}

	if validateMirror(ptn2, 1, 1) {
		t.Error("ptn2: 1 should be a valid mirror location (1 smudge)!")
	}

	if validateMirror(ptn2, 4, 1) {
		t.Error("ptn2: 1 should not be a valid mirror location (1 smudge)!")
	}
}

func TestFindHorizontalMirror(t *testing.T) {
	file, _ := os.Open("test.txt")

	ptns := parsePatterns(file)
	vptn, hptn := ptns[0], ptns[1]
	hmir := findHorizontalMirror(hptn, 0)
	vmir := findHorizontalMirror(vptn, 0)

	if hmir != 4 {
		t.Error("hmir should be 4, got", hmir)
	}

	if vmir != -1 {
		t.Error("vmir should be 4, got", vmir)
	}
}

func TestFindHorizontalSmudgedMirror(t *testing.T) {
	file, _ := os.Open("test.txt")
	
	ptns := parsePatterns(file)
	ptn1, ptn2 := ptns[0], ptns[1]
	
	mir1 := findHorizontalMirror(ptn1, 1)
	mir2 := findHorizontalMirror(ptn2, 1)

	if mir1 != 3 {
		t.Error("mir1 should be 3, got", mir1)
	}
	
	if mir2 != 1 {
		t.Error("mir2 should be 1, got", mir2)
	}
}

func TestPart1(t *testing.T) {
	utils.Tester(t, "test.txt", Part1, "405")()
}

func TestPart2(t *testing.T) {
	utils.Tester(t, "test.txt", Part2, "400")()
}