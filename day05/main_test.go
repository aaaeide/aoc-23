package day05

import (
	"testing"

	"github.com/aaaeide/aoc-23/utils"
)

func TestPart1(t *testing.T) {
	utils.Tester(t, "test.txt", Part1, "35")()
}

// func TestPart2(t *testing.T) {
// 	utils.Tester(t, "test.txt", Part2, "46")()
// }

// func assertRangesAreEqual(t *testing.T, got, want []Range) {
// 	if rangesToString(got) != rangesToString(want) {
// 		t.Errorf("got %s, wanted %s", rangesToString(got), rangesToString(want))
// 	}
// }

// func TestUnion(t *testing.T) {
// 	A := Range{0, 9}
// 	B := Range{10, 15}

// 	got, err := A.Union(B)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	want := Range{0, 15}
// 	if got != want {
// 		t.Errorf("got %s, wanted %s", got.ToString(), want.ToString())
// 	}
// }

// func TestTransform1(t *testing.T) {
// 	r := Range{start: 0, end: 9}
// 	got := transform([]Range{r}, Range{15, 16}, Range{5,6})
// 	want := []Range{{0,4}, {7,9}, {15,16}}
	
// 	assertRangesAreEqual(t, got, want)
// }

// func TestTransform2(t *testing.T) {
// 	r := Range{start: 0, end: 9}
// 	got := transform([]Range{r}, Range{10, 15}, Range{0,5})
// 	want := []Range{{6,15}}

// 	assertRangesAreEqual(t, got, want)
// }

// func TestTransform3(t *testing.T) {
// 	r := Range{start: 0, end: 9}
// 	got := transform([]Range{r}, Range{10, 15}, Range{0,3})
// 	want := []Range{{4,13}}

// 	assertRangesAreEqual(t, got, want)
// }

// func TestTransform4(t *testing.T) {
// 	r := Range{start: 10, end: 19}
// 	got := transform([]Range{r}, Range{0, 3}, Range{10,13})
// 	want := []Range{{0,3}, {14,19}}

// 	assertRangesAreEqual(t, got, want)
// }

// func TestTransform5(t *testing.T) {
// 	r := Range{start: 0, end: 9}
// 	got := transform([]Range{r}, Range{100, 200}, Range{10, 20})
// 	want := []Range{{0, 9}}

// 	assertRangesAreEqual(t, got, want)
// }