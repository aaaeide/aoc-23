package day08

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/aaaeide/aoc-23/utils"
)

func parse(file *os.File) ([]string, map[string]map[string]string) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	
	directions := strings.Split(scanner.Text(), "")
	nodes := make(map[string]map[string]string, 0)
	
	scanner.Scan() // empty line
	
	ptn := regexp.MustCompile(`(\w\w\w) = \((\w\w\w), (\w\w\w)\)`)
	
	for scanner.Scan() {
		matches := ptn.FindAllStringSubmatch(scanner.Text(), -1)[0]
		from, l, r := matches[1], matches[2], matches[3]
	
		nodes[from] = map[string]string{"L": l, "R": r}
	}

	return directions, nodes
}

func Part1(file *os.File) string {
	directions, nodes := parse(file)

	cnt := 0
	cur := "AAA"
	for {
		dir := directions[cnt % len(directions)]
		cur = nodes[cur][dir]
		if cur == "ZZZ" {
			return strconv.Itoa(cnt+1)
		}

		cnt += 1
	}
}

func Part2(file *os.File) string {
	directions, nodes := parse(file)

	var current []string
	for nd := range nodes {
		if nd[len(nd)-1] == 'A' {
			current = append(current, nd)
		}
	}

	// assume that each start node has a path to exactly one end node
	var endNodesAtStep []int

	for _, cur := range current {
		at, cnt := cur, 0
		visited := map[string]bool{
			// we never go back to the starting node.
			fmt.Sprintf("%s %d", at, cnt): false,
		}

		for {
			offset := cnt % len(directions)
			dir := directions[offset]
			at = nodes[at][dir]

			if at[2] == 'Z' {
				endNodesAtStep = append(endNodesAtStep, cnt + 1)
				break
			}

			key := fmt.Sprintf("%s %d", at, offset)

			visited[key] = true
			cnt += 1
		}
	}

	lcm := utils.LCM(endNodesAtStep[0], endNodesAtStep[1], endNodesAtStep...)
	return strconv.Itoa(lcm)
}