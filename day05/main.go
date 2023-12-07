package day05

import (
	"bufio"
	"fmt"
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
	fromTo, toFrom map[string]string;
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
		toFrom: make(map[string]string, 0),
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
			almanac.toFrom[to] = from
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

func mapUp(val int, mappers []Mapper) int {
	for _, mapper := range mappers {
		if val >= mapper.destRangeStart && val <= mapper.destRangeStart + mapper.rangeLength {
			// found special mapping
			offset := val - mapper.destRangeStart
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

func findTop(to string, val int, almanac Almanac) int {
	if to == "seed" {
		return val
	}

	print(to, " ", val)

	from := almanac.toFrom[to]
	up := mapUp(val, almanac.mappers[to])

	println(" corresponds to", from, up)
	
	return findTop(from, up, almanac)
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

type SeedRange struct {
	start, length int;
}

func Part2(file *os.File) string {
	almanac := parseAlmanac(file)

	println(mapDown(50, almanac.mappers["seed"]))
	println(mapUp(52, almanac.mappers["soil"]))

	var seedRanges []SeedRange
	for i := 0; i < len(almanac.seeds); i += 2 {
		seedRanges = append(seedRanges, SeedRange{
			start: almanac.seeds[i],
			length: almanac.seeds[i+1],
		})
	}
	
	location := 46
	for {
		seed := findTop("location", location, almanac)

		for _, seedRange := range seedRanges {
			if seed >= seedRange.start && seed <= seedRange.start + seedRange.length {
				println("found seed", seed)
				return strconv.Itoa(location)
			}
		}

		location++
	}
}

type Cache map[string]int

func memoizedFindBottom(from string, val int, almanac Almanac, cache Cache) (int, Cache) {
	if from == "location" {
		return val, cache
	}

	key := fmt.Sprintf(from + "%d", val)
	hit, ok := cache[key]
	if ok {
		println("cache hit!")
		return hit, cache
	}

	to := almanac.fromTo[from]
	res, cache := memoizedFindBottom(to, mapDown(val, almanac.mappers[from]), almanac, cache)
	cache[key] = res
	return res, cache
}

// still super slow
func MemoizedPart2(file *os.File) string {
	almanac := parseAlmanac(file)
	cache := make(Cache, 0)

	var loc int
	least := int(math.Inf(1))

	for i := 0; i < len(almanac.seeds); i += 2 {
		start, end := almanac.seeds[i], almanac.seeds[i] + almanac.seeds[i+1]
		for seed := start; seed < end; seed++ {
			loc, cache = memoizedFindBottom("seed", seed, almanac, cache)
			if loc < least {
				least = loc
			}
		}
	}

	return strconv.Itoa(least)
}