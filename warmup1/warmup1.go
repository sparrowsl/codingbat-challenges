package warmup1

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("CodingBat Warmup 1 in Go")
}

// The parameter weekday is True if it is a weekday, and the parameter vacation is True if we are on vacation.
// We sleep in if it is not a weekday or we're on vacation. Return True if we sleep in.
func SleepIn(weekday, vacation bool) bool {
	return (!weekday || vacation)
}

// We have two monkeys, a and b, and the parameters a_smile and b_smile indicate if each is smiling.
// We are in trouble if they are both smiling or if neither of them is smiling.
// Return True if we are in trouble.
func MonkeyTrouble(a_smile, b_smile bool) bool {
	return (a_smile == b_smile)
}

// Given two int values, return their sum.
// Unless the two values are the same, then return double their sum.
func SumDouble(a, b int) int {
	if a == b {
		return (a + b) * 2
	}
	return a + b
}

// Given an int n, return the absolute difference between n and 21,
// except return double the absolute difference if n is over 21.
func Diff21(n int) int {
	if n <= 21 {
		return 21 - n
	} else {
		return (n - 21) * 2
	}
}

// Given 2 ints, a and b, return True if one if them is 10 or if their sum is 10.
func Makes10(a, b int) bool {
	return a == 10 || b == 10 || (a+b) == 10
}

// Given an int n, return True if it is within 10 of 100 or 200.
// Note: abs(num) computes the absolute value of a number.
func NearHundred(n int) bool {
	return (math.Abs(100-float64(n)) <= 10) || (math.Abs(200-float64(n)) <= 10)
}

// Given 2 int values, return True if one is negative and one is positive.
// Except if the parameter "negative" is True, then return True only if both are negative.
func PosNeg(a, b int, negative bool) bool {
	if negative {
		return (a < 0 && b < 0)
	} else {
		return (a < 0 && b >= 0) || (a >= 0 && b < 0)
	}
}

// Given a string, return a new string where "not " has been added to the front.
// However, if the string already begins with "not", return the string unchanged.
func NotString(sentence string) string {
	firstWord := strings.Split(sentence, " ")[0]

	if firstWord == "not" {
		return sentence
	}

	return ("not " + sentence)
}

// Given a non-empty string and an int n,
// return a new string where the char at index n has been removed.
// The value of n will be a valid index of a char in the original string
// (i.e. n will be in the range 0..len(str)-1 inclusive).
func MissingChar(sentence string, n int) string {
	var finalString string

	for i, char := range sentence {
		if i != n {
			finalString += string(char)
		}
	}

	return finalString
}

// Given a string, return a new string where the first and last chars have been exchanged.
func FrontBack(word string) string {
	if len(word) <= 1 {
		return word
	}

	firstChar := word[0]
	lastChar := word[len(word)-1]
	remaining := word[1 : len(word)-1]

	return string(lastChar) + remaining + string(firstChar)
}

// Given a string, we'll say that the front is the first 3 chars of the string.
// If the string length is less than 3, the front is whatever is there.
// Return a new string which is 3 copies of the front.
func Front3(word string) string {
	if len(word) < 3 {
		return strings.Repeat(word, 3)
	}

	return strings.Repeat(word[:3], 3)
}
