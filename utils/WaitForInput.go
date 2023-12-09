package utils

import (
	"bufio"
	"os"
)

func WaitForInput() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}