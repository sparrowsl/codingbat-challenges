package string1

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("CodingBat String 1 in Go")
}

// Given a string name, e.g. "Bob", return a greeting of the form "Hello Bob!".
func helloName(name string) string {
	return fmt.Sprintf("Hello %v!", name)
}

// Given two strings, a and b, return the result of putting them together in the order abba,
// e.g. "Hi" and "Bye" returns "HiByeByeHi".
func makeAbba(a, b string) string {
	return a + b + b + a
}

// The web is built with HTML strings like "<i>Yay</i>" which draws Yay as italic text.
// In this example, the "i" tag makes <i> and </i> which surround the word "Yay".
// Given tag and word strings, create the HTML string with tags around the word, e.g. "<i>Yay</i>".
func makeTags(tag, word string) string {
	return fmt.Sprintf("<%v>%v</%[1]v>", tag, word)
}

// Given an "out" string length 4, such as "<<>>", and a word,
// return a new string where the word is in the middle of the out string, e.g. "<<word>>".
func makeOutWord(out, word string) string {
	return fmt.Sprintf("%v%v%v", out[:2], word, out[2:])
}

// Given a string, return a new string made of 3 copies of the last 2 chars of the original string.
// The string length will be at least 2.
func extraEnd(word string) string {
	return strings.Repeat(word[len(word)-2:], 3)
}

// Given a string, return the string made of its first two chars,
// so the String "Hello" yields "He". If the string is shorter than length 2,
// return whatever there is, so "X" yields "X", and the empty string "" yields the empty string "".
func firstTwo(word string) string {
	if len(word) <= 2 {
		return word
	}
	return word[:2]
}
