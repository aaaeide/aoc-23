package utils

func MinIntSlice(ints []int) int {
	if len(ints) == 0 {
		return 0
	}
	
	min := ints[0]
	for _, e := range ints {
		if e < min {
			min = e
		}
	}
	return min
}