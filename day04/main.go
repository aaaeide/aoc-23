package day04

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/aaaeide/aoc-23/utils"
)

func parseCards(file *os.File) map[int]int {
	scanner := bufio.NewScanner(file)
	cardPtn := regexp.MustCompile(`Card\s+(\d+): ([\d ]+) \| ([\d ]+)`)

	winsPerCard := make(map[int]int, 0)

	for scanner.Scan() {
		match := cardPtn.FindAllStringSubmatch(scanner.Text(), -1)[0]
		
		cardId, _ := strconv.Atoi(match[1])
		winning := utils.StringToSliceInt(match[2], " ")
		actual := utils.StringToSliceInt(match[3], " ")
		numWinning := len(utils.Intersection[int](winning, actual))

		winsPerCard[cardId] = numWinning
	}

	return winsPerCard
}

func Part1(file *os.File) string {
	winsPerCard := parseCards(file)

	sum := 0
	for _, wins := range winsPerCard {
		sum += int(math.Pow(2, float64(wins)-1))
	}

	return strconv.Itoa(sum)
}

func Part2(file *os.File) string {
	winsPerCard := parseCards(file)

	// Make initial set of cards
	var cards []int
	for i := 1; i <= len(winsPerCard); i++ {
		cards = append(cards, i)
	}

	i := 0
	for {
		if i == len(cards) {
			break
		}

		card := cards[i]
		for j := card+1; j <= card+winsPerCard[card]; j++ {
			cards = append(cards, j)
		}

		i++
	}

	return strconv.Itoa(len(cards))
}