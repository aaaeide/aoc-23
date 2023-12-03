package utils

func MaxIntSlice(ints []int) int {
	if len(ints) == 0 {
		return 0
	}
	
	max := ints[0]
	for _, e := range ints {
		if e > max {
			max = e
		}
	}
	return max
}