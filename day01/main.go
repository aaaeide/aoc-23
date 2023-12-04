package day01

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func Part1() string {
	file, _ := os.Open("inputs/day01.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum  := 0

	for scanner.Scan() {
		digits := make([]int, 0)
		for _, c := range scanner.Text() {
			if unicode.IsDigit(c) {
				digits = append(digits, int(c-'0'))
			}
		}

		sum += digits[0] * 10 + digits[len(digits) - 1]
	}

	return strconv.Itoa(sum)
}

func Part2() string {
	file, _ := os.Open("inputs/day01.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum  := 0

	numbers := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	for scanner.Scan() {
		line := scanner.Text()
		digits := [2]int{0, 0}

		for i := range line {
			for str, val := range numbers {
				if i+len(str) > len(line) {
					continue
				}

				if line[i:i+len(str)] == str {
					if digits[0] == 0 {
						digits[0] = val
					}

					digits[1] = val
				}
			}
		}

		sum += digits[0] * 10 + digits[1]
	}

	return strconv.Itoa(sum)
}
