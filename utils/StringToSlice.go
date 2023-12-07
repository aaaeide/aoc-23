package utils

import (
	"strconv"
	"strings"
)

func StringToSliceInt(lst string, sep string) ([]int, error) {
	var ints []int

	for _, numStr := range strings.Split(lst, sep) {
		if numStr == "" {
			continue
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return []int{}, err
		}

		ints = append(ints, num)
	}

	return ints, nil
}