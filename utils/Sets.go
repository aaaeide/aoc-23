package utils

import "golang.org/x/exp/slices"

// The intersection of two sets is the set of elements common to both.
func Intersection[T comparable](slice1, slice2 []T) []T {
	var out []T

	for _, el := range slice1 {
		if slices.Contains(slice2, el) {
			out = append(out, el)
		}
	}

	return out
}