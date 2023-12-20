package day15

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func HASH(s string) (sum int32) {
	for _, r := range s {
		sum += r
		sum *= 17
		sum %= 256
	}

	return sum
}

func Part1(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var sum int32
	for scanner.Scan() {
		instructions := strings.Split(scanner.Text(), ",")
		for _, instr := range instructions {
			sum += HASH(instr)
		}
	}

	return strconv.Itoa(int(sum))
}

type Box struct {
	lenses map[string]int;
	order map[string]int;
}

func Part2(file *os.File) string {
	// initialize boxes
	boxes := make([]Box, 256)
	for i := range boxes {
		boxes[i] = Box{
			lenses: make(map[string]int),
			order: make(map[string]int),
		}
	}

	labelPtn := regexp.MustCompile(`(\w+)`)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := strings.Split(scanner.Text(), ",")

	for _, instr := range instructions {
		match := labelPtn.FindStringIndex(instr)
		label := instr[:match[1]]
		operation := string(instr[match[1]])
		
		box := HASH(label)
		
		if operation == "=" {
			flength, _ := strconv.Atoi(instr[match[1]+1:])
			boxes[box].lenses[label] = flength
			if _, ok := boxes[box].order[label]; !ok {
				boxes[box].order[label] = len(boxes[box].order)
			}
		} else {
			delete(boxes[box].lenses, label)
			if i, ok := boxes[box].order[label]; ok {
				for l, j := range boxes[box].order {
					if j > i {
						boxes[box].order[l]--
					}
				}
			}
			delete(boxes[box].order, label)
		}
	}

	sum := 0
	for i, box := range boxes {
		for label, flength := range box.lenses {
			power := 1 + i
			power *= 1 + box.order[label]
			power *= flength

			sum += power
		}
	}

	return strconv.Itoa(sum)
}