package utils

import "testing"

func Tester(
	t *testing.T, 
	filename string, 
	fn func(string) string,
	want string) func() {
	return func() {
		got := fn(filename)

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}