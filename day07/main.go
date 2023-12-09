package day07

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/aaaeide/aoc-23/utils"
)

type Card string

func (c1 Card) geq(c2 Card, withJokers bool) bool {
	values := map[Card]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": utils.Tif(withJokers, 1, 11),
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	return values[c1] >= values[c2]
}

type Hand struct {
	kind string;
	cards []Card;
	bid int;
	cardcounts map[Card]int;
}

func (h1 Hand) geq (h2 Hand, withJokers bool) bool {
	values := map[string]int{
		"five-of-a-kind": 7,
		"four-of-a-kind": 6,
		"full-house": 5,
		"three-of-a-kind": 4,
		"two-pair": 3,
		"one-pair": 2,
		"high-card": 1,
	}

	if values[h1.kind] == values[h2.kind] {
		for i := 0; i < len(h1.cards); i++ {
			if h1.cards[i] != h2.cards[i] {
				return h1.cards[i].geq(h2.cards[i], withJokers)
			}
		}
	}
	
	return values[h1.kind] >= values[h2.kind]
}

func findKind(cardcounts map[Card]int) string {
	var counts []int
	for _, count := range cardcounts {
		if count != 0 {
			counts = append(counts, count)
		}
	}
	sort.Ints(counts)

	switch len(counts) {
	case 5:
		return "high-card"
	case 4:
		return "one-pair"
	case 3:
		if counts[len(counts) - 1] == 2 {
			return "two-pair"
		} else {
			return "three-of-a-kind"
		}
	case 2:
		if counts[len(counts) - 1] == 3 {
			return "full-house"
		} else {
			return "four-of-a-kind"
		}
	case 1:
		return "five-of-a-kind"
	}

	panic("could not find kind!!")
}

func parseHand(str string, bid int) Hand {
	var cardcounts = make(map[Card]int, 0)
	var cards []Card

	for _, c := range strings.Split(str, "") {
		cards = append(cards, Card(c))

		count, ok := cardcounts[Card(c)]
		if !ok {
			cardcounts[Card(c)] = 1
			continue
		}
		
		cardcounts[Card(c)] = count + 1
	}

	kind := findKind(cardcounts)

	return Hand{kind, cards, bid, cardcounts}
}

func parseHands(file *os.File, withJokers bool) []Hand {
	scanner := bufio.NewScanner(file)
	
	var hands []Hand
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		h, b := tokens[0], tokens[1]
		bid, _ := strconv.Atoi(b)
	
		hands = append(hands, parseHand(h, bid))
	}

	return hands
}

func Part1(file *os.File) string {
	hands := parseHands(file, false)
	
	sort.Slice(hands, func(i, j int) bool {
		return !hands[i].geq(hands[j], false)
	})

	sum := 0
	for i := 0; i < len(hands); i++ {
		sum += hands[i].bid * (i+1)
	}
	
	return strconv.Itoa(sum)
}

func Part2(file *os.File) string {
	hands := parseHands(file, true)

	var jokeredHands []Hand
	for _, hand := range hands {
		jokers, ok := hand.cardcounts["J"]
		if !ok || jokers == 5 {
			jokeredHands = append(jokeredHands, hand)
			continue
		}

		var commonest Card
		for card, count := range hand.cardcounts {
			if card != "J" && count >= hand.cardcounts[commonest] {
				commonest = card
			}
		}

		hand.cardcounts[commonest] += jokers
		hand.cardcounts["J"] = 0
		hand.kind = findKind(hand.cardcounts)

		jokeredHands = append(jokeredHands, hand)
	}

	sort.Slice(jokeredHands, func(i, j int) bool {
		return !jokeredHands[i].geq(jokeredHands[j], true)
	})

	sum := 0
	for i := 0; i < len(jokeredHands); i++ {
		sum += jokeredHands[i].bid * (i+1)
	}

	return strconv.Itoa(sum)
}