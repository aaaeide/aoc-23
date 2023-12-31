package day05

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/aaaeide/aoc-23/utils"
)

// 50 98 2
// 52 50 48
// The first line has a destination range start of 50, a source range start of 98, and a range
// length of 2. This line means that the source range starts at 98 and contains two values:
// 98 and 99. The destination range is the same length, but it starts at 50, so its two values
// are 50 and 51. With this information, you know that seed number 98 corresponds to soil number
// 50 and that seed number 99 corresponds to soil number 51.

type Mapper struct{
	destRangeStart int;
	sourceRangeStart int;
	rangeLength int;
}

type Almanac struct {
	seeds []int;
	fromTo map[string]string;
	mappers map[string][]Mapper;
}

func parseAlmanac(file *os.File) Almanac {
	scanner := bufio.NewScanner(file)

	scanner.Scan()

	// "seeds: 79 14 55 13"
	seeds, _ := utils.StringToSliceInt(strings.Split(scanner.Text(), ": ")[1], " ")

	mapPtn := regexp.MustCompile(`(\w+)-to-(\w+) map:`)

	almanac := Almanac{
		seeds: seeds,
		fromTo: make(map[string]string, 0),
		mappers: make(map[string][]Mapper, 0),
	}

	var currentType string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		nums, err := utils.StringToSliceInt(line, " ")
		if err != nil {
			// "x-to-y map:"
			headerMatch := mapPtn.FindAllStringSubmatch(line, -1)[0]
			from, to := headerMatch[1], headerMatch[2]

			almanac.fromTo[from] = to
			currentType = from
			
			continue
		}

		almanac.mappers[currentType] = append(almanac.mappers[currentType], Mapper{
			destRangeStart: nums[0],
			sourceRangeStart: nums[1],
			rangeLength: nums[2],
		})
	}

	return almanac
}

func mapDown(val int, mappers []Mapper) int {
	for _, mapper := range mappers {
		if val >= mapper.sourceRangeStart && val <= mapper.sourceRangeStart + mapper.rangeLength {
			// found special mapping
			offset := val - mapper.sourceRangeStart
			return mapper.destRangeStart + offset
		}
	}

	// return default mapping
	return val
}

func findBottom(from string, val int, almanac Almanac) int {
	if from == "location" {
		return val
	}

	to := almanac.fromTo[from]
	
	return findBottom(to, mapDown(val, almanac.mappers[from]), almanac)
}

func Part1(file *os.File) string {
	almanac := parseAlmanac(file)

	least := int(math.Inf(1))
	for _, seed := range almanac.seeds {
		loc := findBottom("seed", seed, almanac)
		if loc < least {
			least = loc
		}
	}

	return strconv.Itoa(least)
}
