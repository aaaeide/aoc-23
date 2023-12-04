package day03

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
)

type PartNumber struct {
	number, part string;
	x1, x2, y int;
}

type Symbol struct {
	part string;
	adjPartNums []int;
}

func solve(file *os.File) (string, string) {
	scanner := bufio.NewScanner(file)
	
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	numberPtn := regexp.MustCompile(`(\d+)`)
	
	partNumbers := make([]PartNumber, 0)
	
	for idx, line := range lines {
		numberPositions := numberPtn.FindAllStringIndex(line, -1)
		for _, numberPosition := range numberPositions {
			start, end := numberPosition[0], numberPosition[1]
			number := line[start:end]
			
			partNumbers = append(partNumbers, PartNumber{
				number: number,
				part: "not set",
				x1: start,
				x2: end-1,
				y: idx,
			})
		}
	}

	symbolPtn := regexp.MustCompile(`([^\d.])`)

	symbols := make([]Symbol, 0)

	for idx, line := range lines {
		symbolPositions := symbolPtn.FindAllStringIndex(line, -1)

		for _, symbolPosition := range symbolPositions {
			start, end := symbolPosition[0], symbolPosition[1]
			part := line[start:end]
			var adjPartNums []int

			for pni, partNumberStruct := range partNumbers {

				if math.Abs(float64(idx - partNumberStruct.y)) <= 1.0 && partNumberStruct.x1 - 1 <= start && start <= partNumberStruct.x2 + 1 {
					partNumberStruct.part = part
					partNumberInt, _ := strconv.Atoi(partNumberStruct.number)
					adjPartNums = append(adjPartNums, partNumberInt)
				}

				partNumbers[pni] = partNumberStruct
			}

			symbols = append(symbols, Symbol{
				part: part,
				adjPartNums: adjPartNums,
			})
		}
	}

	sum1 := 0

	for _, partNumberStruct := range partNumbers {
		if partNumberStruct.part != "not set" {
			num, _ := strconv.Atoi(partNumberStruct.number)
			sum1 += num
		}
	}

	sum2 := 0

	for _, symbolStruct := range symbols {
		if symbolStruct.part == "*" && len(symbolStruct.adjPartNums) == 2 {
			sum2 += symbolStruct.adjPartNums[0] * symbolStruct.adjPartNums[1]
		}
	}

	return strconv.Itoa(sum1), strconv.Itoa(sum2)
}

func Part1(file *os.File) string {
	sol, _ := solve(file)
	return sol
}

func Part2(file *os.File) string {
	_, sol := solve(file)
	return sol
}