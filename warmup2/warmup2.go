package warmup2

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("CodingBat Warmup 2 in Go")
}

// Given a string and a non-negative int n, return a larger string that is n copies of the original string.
func StringTimes(word string, n uint) string {
	return strings.Repeat(word, int(n))
}
