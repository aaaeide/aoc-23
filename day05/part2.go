package day05

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	u "github.com/aaaeide/aoc-23/utils"
)

// start to end (inclusive)
type Range struct {
	start, end int;
}

func (r Range) ToString() string {
	return fmt.Sprintf("(%d, %d)", r.start, r.end)
}

func (r Range) Length() int {
	return r.end - r.start + 1
}

func rangesToString(rs []Range) string {
	out := "[ "
	for _, r := range rs {
		out += r.ToString() + " "
	}
	return out + "]"
}

func sortRanges(rs []Range) []Range {
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].start < rs[j].start
	})

	return rs
}

// C = A \cap	B
func (A Range) Intersect(B Range) Range {
	C := Range{}

	if B.end < A.start || B.start > A.end {
		return C
	}

	C.start = u.Tif(B.start < A.start, A.start, B.start)
	C.end = u.Tif(A.end < B.end, A.end, B.end)

	return C
}

// C = A \setminus B
func (A Range) Minus(B Range) []Range {
	null := Range{}
	cap := A.Intersect(B)
	if cap != null {
		var before, after Range

		if A.start < cap.start {
			before.start = A.start
			before.end = cap.start - 1
		}

		if A.end > cap.end {
			after.start = cap.end + 1
			after.end = A.end
		}

		var C []Range
		C = u.Tif(before != null, append(C, before), C)
		C = u.Tif(after != null, append(C, after), C)

		return C
	}

	return []Range{A}
}

func (A Range) Union(B Range) (Range, error) {
	if A.start > B.start {
		return Range{}, errors.New("error: A must be smaller than B")
	}

	if A.end == B.start - 1 {
		return Range{A.start, B.end}, nil
	}

	cap := A.Intersect(B)
	null := Range{}

	if cap == null {
		return Range{}, errors.New("no overlap")
	}

	if cap == A {
		return B, nil
	}

	if cap == B {
		return A, nil
	}

	left, right := A.Minus(cap)[0], B.Minus(cap)[0]
	return Range{left.start, right.end}, nil
}

func doCoalesce(in, out []Range) []Range {
	// println("\ndocoal")
	// println(rangesToString(in))
	// println(rangesToString(out))
	if len(in) == 0 {
		return out
	}

	last := out[len(out)-1]

	union, err := last.Union(in[0])
	
	if err != nil {
		out = append(out, in[0])
		return doCoalesce(in[1:], out)
	}

	return doCoalesce(in[1:], append(out[:len(out)-1], union))
}

// [{0,9},{10,19}] => [{0,19}]
// [{0,9},{5, 19}] => [{0,19}]
func coalesce(rs []Range) []Range {
	rs = sortRanges(rs)
	return doCoalesce(rs[1:], []Range{rs[0]})
}

func transform(rs []Range, d, s Range) []Range {
	println("\n\n\ntransform")
	println(rangesToString(rs))
	println(d.ToString(), s.ToString())
	out := []Range{}
	null := Range{}

	for _, r := range rs {
		println("checking r=" + r.ToString()+" against s="+s.ToString())
		if r.Intersect(s) == null {
			println("no intersection")
			out = append(out, r)
			continue
		}

		println("intersects:")

		res := r.Minus(s)
		println("\twithout r\\s="+rangesToString(res))

		diff := d.start - s.start
		println("\tdiff=",diff)
		start := u.Tif(r.start < s.start, s.start, r.start)
		end := u.Tif(r.end < s.end, r.end, s.end)
		println("\tend=", end)

		res = append(res, Range{start + diff, end + diff})
		println("\tresult of transformation r~(s,d)="+rangesToString(res))

		out = append(out, res...)
	}

	return coalesce(out)
}


func Part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	seeds, _ := u.StringToSliceInt(strings.Split(scanner.Text(), ": ")[1], " ")

	rs := []Range{}
	for i := 0; i < len(seeds); i+=2 {
		rs = append(rs, Range{seeds[i], seeds[i] + seeds[i+1] - 1})
	}

	for scanner.Scan() {
		line := scanner.Text()

		nums, err := u.StringToSliceInt(line, " ")
		if line == "" || err != nil {
			continue
		}

		fmt.Println(nums)
		rs = transform(rs, Range{nums[0], nums[0] + nums[2] - 1}, Range{nums[1], nums[1] + nums[2] - 1})
	}

	return strconv.Itoa(rs[0].start)
}