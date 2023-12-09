package day06

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/aaaeide/aoc-23/utils"
)

type Race struct{
	time, record int;
}

func parsePart1(file *os.File) []Race {
	scanner := bufio.NewScanner(file)
	scanner.Scan() // Time line
	times, _ := utils.StringToSliceInt(scanner.Text(), " ")
	times = times[1:]
	
	scanner.Scan() // Distance line
	records, _ := utils.StringToSliceInt(scanner.Text(), " ")
	records = records[1:]

	if len(times) != len(records) {
		panic("times and records not of same length!!")
	}

	var races []Race

	for i := 0; i < len(times); i++ {
		races = append(races, Race{
			time: times[i],
			record: records[i],
		})
	}

	return races
}

func parsePart2(file *os.File) Race {
	scanner := bufio.NewScanner(file)
	scanner.Scan() // Time line
	
	time, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(scanner.Text(), " ", ""), ":")[1])
	
	scanner.Scan() // Distance line
	
	record, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(scanner.Text(), " ", ""), ":")[1])

	return Race{time, record}
}

func quadratic(ai, bi, ci int) (float64, float64) {
	a, b, c := float64(ai), float64(bi), float64(ci)
	x1 := (-b + math.Sqrt(math.Pow(b, 2) - 4*a*c))/2*a
	x2 := (-b - math.Sqrt(math.Pow(b, 2) - 4*a*c))/2*a

	return x1, x2
}

func Part1(file *os.File) string {
	races := parsePart1(file)
	var solution float64 = 1

	// find number of record-breaking hold times for each race
	for _, race := range races {
		t1, t2 := quadratic(-1, race.time, -race.record)
		num := math.Ceil(t2) - math.Floor(t1+1)

		solution *= num
	}


	return strconv.Itoa(int(solution))
}


func Part2(file *os.File) string {
	race := parsePart2(file)

	t1, t2 := quadratic(-1, race.time, -race.record)
	num :=  math.Ceil(t2) - math.Floor(t1+1)

	return strconv.Itoa(int(num))
}