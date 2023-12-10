package day09

import (
	"bufio"
	"os"
	"strconv"

	"github.com/aaaeide/aoc-23/utils"
)

func extrapolate(line string) (int, int) {
	nums, _ := utils.StringToSliceInt(line, " ")
	var diffs [][]int = [][]int{nums}

	for {
		prev := diffs[len(diffs)-1]
		var cur []int

		allZeros := true
		
		for i := 1; i < len(prev); i++ {
			diff := prev[i] - prev[i-1]
			cur = append(cur, diff)
			if diff != 0 {
				allZeros = false
			}
		}
		
		diffs = append(diffs, cur)

		if allZeros {
			break
		}
	}

	next, prev := 0, 0
	for i := len(diffs)-2; i >= 0; i-- {
		next = diffs[i][len(diffs[i])-1] + next
		prev = diffs[i][0] - prev
	}

	return next, prev
}

func solve(file *os.File) (int, int) {
	scanner := bufio.NewScanner(file)
	
	sum1, sum2 := 0, 0
	for scanner.Scan() {
		next, prev := extrapolate(scanner.Text())
		sum1 += next
		sum2 += prev
	}

	return sum1, sum2
}

func Part1(file *os.File) string {
	sol, _ := solve(file)
	return strconv.Itoa(sol)
}

func Part2(file *os.File) string {
	_, sol := solve(file)
	return strconv.Itoa(sol)
}