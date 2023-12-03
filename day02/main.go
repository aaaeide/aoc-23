package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/aaaeide/aoc-23/utils"
)

func parseGame(gameString string) []map[string][]int {
	idPtn := regexp.MustCompile(`Game (\d+): (.+)`)
	drawPtn := regexp.MustCompile(`(\d+) (red|blue|green)`)

	matches := idPtn.FindStringSubmatch(gameString)
	// gameId, _ := strconv.Atoi(matches[1])
	gameStr := matches[2]


	rounds := make([]map[string][]int, 0)

	for idx, round := range strings.Split(gameStr, "; ") {
		rounds = append(rounds, map[string][]int{
			"red": {},
			"blue": {},
			"green": {},
		})

		for _, draw := range strings.Split(round, ", ") {
			matches = drawPtn.FindStringSubmatch(draw)
			amt, _ := strconv.Atoi(matches[1])
			rounds[idx][matches[2]] = append(rounds[idx][matches[2]], amt)
		}
	}

	return rounds
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	gameId := 1
	sumPart1 := 0
	sumPart2 := 0

	for scanner.Scan() {
		rounds := parseGame(scanner.Text())
		validGame := true

		mostThisGame := map[string]int{
			"red": 0,
			"blue": 0,
			"green": 0,
		}

		for _, round := range rounds {
			if (utils.SumIntSlice(round["red"]) > 12 || utils.SumIntSlice(round["green"]) > 13 || utils.SumIntSlice(round["blue"]) > 14) {
				validGame = false
			}

			for _, col := range []string{"red", "blue", "green"} {
				mostThisRound := utils.MaxIntSlice(round[col])
				if mostThisRound > mostThisGame[col] {
					mostThisGame[col] = mostThisRound
				}
			}
		}

		if validGame {
			sumPart1 += gameId
		}
		gameId += 1

		sumPart2 += mostThisGame["red"] * mostThisGame["blue"] * mostThisGame["green"]

	}

	println(sumPart1)
	println(sumPart2)
}