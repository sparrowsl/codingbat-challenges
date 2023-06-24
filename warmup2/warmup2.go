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

// Given a string and a non-negative int n,
// we'll say that the front of the string is the first 3 chars,
// or whatever is there if the string is less than length 3.
// Return n copies of the front;
func FrontTimes(word string, n uint) string {
	return strings.Repeat(word[:3], int(n))
}
