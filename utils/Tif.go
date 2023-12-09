package utils

// Ternary if. If true, returns a, else b.
func Tif[T any](cond bool, a, b T) T {
	if cond {
		return a
	}

	return b
}