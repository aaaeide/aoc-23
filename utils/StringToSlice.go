package utils

import (
	"strconv"
	"strings"
)

func StringToSliceInt(lst string, sep string) []int {
	var ints []int

	for _, numStr := range strings.Split(lst, sep) {
		if numStr == "" {
			continue
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic("Error converting string to int: " + err.Error())
		}

		ints = append(ints, num)
	}

	return ints
}