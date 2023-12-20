package day12

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/aaaeide/aoc-23/utils"
)

type Line struct {
	str string;
	groupSizes []int;
}

func parseLines(file *os.File, part int) (lines []Line) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		str, groupsStr := split[0], split[1]

		if part == 2 {
			str = strings.Repeat(str + "?", 5)
			str = str[:len(str)-1]
			groupsStr = strings.Repeat(groupsStr + ",", 5)
		}

		str = regexp.MustCompile(`(\.+)`).ReplaceAllString(str, ".")

		groups, _ := utils.StringToSliceInt(groupsStr, ",")
		lines = append(lines, Line{str, groups})
	}

	return lines
}

func checkLine(line Line) bool {
	// no more question marks -- are we happy?
	regex := regexp.MustCompile(`(#+)`)
	index := 0

	for _, groupSize := range line.groupSizes {
		loc := regex.FindStringIndex(line.str[index:])
		if loc == nil {
			// out of groups
			return false
		}

		group := line.str[index+loc[0]:index+loc[1]]
		if len(group) != groupSize {
			return false
		}

		index += loc[1]
	}

	return regex.FindStringIndex(line.str[index:]) == nil 
}

func countValidArrangements(line Line, start int, inGroup bool) int {
	// fmt.Println("\n"+line.str, line.groupSizes)
	for i:=start; i<len(line.str); i++ {
		c := line.str[i]

		switch c {
		case '#':
			if len(line.groupSizes) == 0 {
				// println(strings.Repeat(" ", i) + "^")
				// fmt.Println("not valid: encountered # with no groups left!")
				// fmt.Println(line.groupSizes)
				return 0
			}
			if line.groupSizes[0] == 0 {
				// println(strings.Repeat(" ", i) + "^")
				// fmt.Println("not valid: encountered # for empty group!")
				// fmt.Println(line.groupSizes)
				return 0
			}
			
			inGroup = true
			line.groupSizes[0] -= 1
		case '.':
			if inGroup {
				if len(line.groupSizes) >= 1 && line.groupSizes[0] > 0 {
					// println(strings.Repeat(" ", i) + "^")
					// fmt.Println("not valid: encountered . with unfinished group!")
					// fmt.Println(line.groupSizes)
					return 0
				}

				if len(line.groupSizes) > 1 {
					line.groupSizes = line.groupSizes[1:]
				} else if len(line.groupSizes) == 1 {
					line.groupSizes = []int{}
				}
				inGroup = false
			}
		case '?':
			gs1, gs2 := make([]int, len(line.groupSizes)), make([]int, len(line.groupSizes))
			copy(gs1, line.groupSizes)
			copy(gs2, line.groupSizes)

			return countValidArrangements(Line{
				line.str[:i] + "." + line.str[i+1:],
				gs1,
			}, i, inGroup) + countValidArrangements(Line{
				line.str[:i] + "#" + line.str[i+1:],
				gs2,
			}, i, inGroup)
		}
	}

	if len(line.groupSizes) == 0 || 
		(len(line.groupSizes) == 1 && line.groupSizes[0] == 0) {
		// fmt.Println("valid!")
		// fmt.Println(line.groupSizes)
		return 1
	}

	// fmt.Println("not valid: string ended with unfinished groups")
	// fmt.Println(line.groupSizes)
	return 0
}

func Part1(file *os.File) string {
	lines := parseLines(file, 1)
	
	sum := 0
	for _, line := range lines {
		sum += countValidArrangements(line, 0, false)
	}

	return strconv.Itoa(sum)
}

func Part2(file *os.File) string {
	return "too slow"

	// lines := parseLines(file, 2)
	// println("STUDYING", len(lines), "LINES")
	
	// sum := 0
	// for _, line := range lines {
	// 	sum += countValidArrangements(line, 0, false)
	// 	println("SUM =", sum)
	// }

	// return strconv.Itoa(sum)
}