package utils

// Uses the Euclidean algorithm to find the greatest common divisor of a and b.
func GCD(x, y int) int {
	for y != 0 {
		z := y
		y = x % y
		x = z
	}

	return x
}

// Find the Least Common Multiple using GCD
func LCM(x, y int, zs ...int) int {
	res := x * y / GCD(x, y)
	for i := 0; i < len(zs); i++ {
		res = LCM(res, zs[i])
	}

	return res
}