package warmup2

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("CodingBat Warmup 2 in Go")
}

// Given a string and a non-negative int n,
// return a larger string that is n copies of the original string.
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

// Given a string, return a new string made of every other char starting with the first,
// so "Hello" yields "Hlo".
func StringBits(word string) string {
	var final string

	for i, char := range word {
		if i%2 == 0 {
			final += string(char)
		}
	}
	return final
}

// Given an array of ints, return the number of 9's in the array.
func arrayCount(array []int) int {
	var nines int

	for _, n := range array {
		if n == 9 {
			nines += 1
		}
	}

	return nines
}

// Given an array of ints, return True if one of the first 4 elements in the array is a 9.
// The array length may be less than 4.
func arrayFront(array []int) bool {
	return false
}
