package utils

import (
	"strconv"
	"strings"
)

func StringToSliceInt(lst string, sep string) ([]int, error) {
	var ints []int
	var err error = nil

	for _, numStr := range strings.Split(lst, sep) {
		if numStr == "" {
			continue
		}

		var num int
		num, err = strconv.Atoi(numStr)

		ints = append(ints, num)
	}

	return ints, err
}