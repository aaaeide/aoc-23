package day13

import (
	"bufio"
	"os"
	"strconv"
)

func parsePatterns(file *os.File) [][]string {
	scanner := bufio.NewScanner(file)
	var patterns [][]string
	patterns = append(patterns, make([]string, 0))
	index := 0
	
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			patterns = append(patterns, make([]string, 0))
			index++
	
			continue
		}
	
		patterns[index] = append(patterns[index], line)
	}

	return patterns
}

func discrepancy(s1, s2 string) (e int) {
	for i:=0; i<len(s1); i++ {
		if s1[i] != s2[i] {
			e++
		}
	}

	return e
}

func validateMirror(pattern []string, mirror, smudges int) bool {
	discrepancies := 0
	for idx:=1; mirror-1-idx >= 0 && mirror+idx < len(pattern); idx++ {
		discrepancies += discrepancy(pattern[mirror-1-idx], pattern[mirror+idx])
		if  discrepancies > smudges {
			return false
		}
	}

	return discrepancies == smudges
}

func findHorizontalMirror(ptn []string, smudges int) int {
	for mir:=1; mir<len(ptn); mir++ {
		if discrepancy(ptn[mir], ptn[mir-1]) <= smudges {
			if validateMirror(ptn, mir, smudges - discrepancy(ptn[mir], ptn[mir-1])) {
				return mir
			}
		}
	}

	return -1
}

func tilt90(ptn []string) (out []string) {
	out = make([]string, len(ptn[0]))

	for i:=0; i<len(ptn[0]); i++ {
		for j:=len(ptn)-1; j>=0; j-- {
			out[i] += string(ptn[j][i])
		}
	}

	return out
}

func Part1(file *os.File) string {
	ptns := parsePatterns(file)

	var mirs []int
	for _, ptn := range ptns {
		mirs = append(mirs, findHorizontalMirror(ptn, 0))
	}

	sum := 0
	for idx, mir := range mirs {
		if mir == -1 {
			// vertical mirror
			sum += findHorizontalMirror(tilt90(ptns[idx]), 0)
			continue
		}

		sum += 100*mir
	}
	
	return strconv.Itoa(sum)
}

func Part2(file *os.File) string {
	ptns := parsePatterns(file)

	var mirs []int
	for _, ptn := range ptns {
		mirs = append(mirs, findHorizontalMirror(ptn, 1))
	}

	sum := 0
	for idx, mir := range mirs {
		if mir == -1 {
			// vertical mirror
			sum += findHorizontalMirror(tilt90(ptns[idx]), 1)
			continue
		}

		sum += 100*mir
	}
	
	return strconv.Itoa(sum)
}