package utils

import (
	"os"
	"testing"
)

func Tester(
	t *testing.T, 
	filename string, 
	fn func(*os.File) string,
	want string) func() {
	return func() {
		file, err := os.Open(filename)
		
		if err != nil {
			t.Errorf("\nCOULD NOT OPEN FILE " + filename)
			return
		}

		got := fn(file)

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}