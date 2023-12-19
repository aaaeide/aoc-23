package day14

import (
	"bufio"
	"os"
	"strconv"
)

func Part1(file *os.File) string {
	scanner := bufio.NewScanner(file)
	
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	HEIGHT, WIDTH := len(lines), len(lines[0])
	// for each column, the topmost index at which an O will stop
	blockers := make(map[int]int, 0)
	// for each column, the number of Os found below lowest blocker
	sliders := make(map[int]int, 0)
	for i:=0; i<WIDTH; i++ {
		blockers[i] = 0
		sliders[i] = 0
	}

	sum := 0
	for i:=0; i<HEIGHT; i++ {
		for j:=0; j<WIDTH; j++ {
			switch lines[i][j] {
			case 'O':
				sliders[j]++
			case '#':
				for k:=0; k<sliders[j]; k++ {
					sum += HEIGHT - blockers[j] - k
				}
				sliders[j] = 0
				blockers[j] = i + 1
			case '.':
				continue
			}
		}
	}

	for j:=0; j<WIDTH; j++ {
		for k:=0; k<sliders[j]; k++ {
			sum += HEIGHT - blockers[j] - k
		}
	}

	return strconv.Itoa(sum)
}

func Part2(file *os.File) string {
	return "not implemented"
}