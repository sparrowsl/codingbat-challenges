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

// Given a string of even length, return the first half.
// So the string "WooHoo" yields "Woo".
func firstHalf(word string) string {
	return word[:len(word)/2]
}

// Given a string, return a version without the first and last char,
// so "Hello" yields "ell". The string length will be at least 2.
func withoutEnd(word string) string {
	return word[1 : len(word)-1]
}

// Given 2 strings, a and b, return a string of the form short+long+short,
// with the shorter string on the outside and the longer string on the inside.
// The strings will not be the same length, but they may be empty (length 0).
func comboString(a, b string) string {
	if len(a) > len(b) {
		return b + a + b
	}
	return a + b + a
}

// Given 2 strings, return their concatenation, except omit the first char of each.
// The strings will be at least length 1.
func nonStart(a, b string) string {
	return a[1:] + b[1:]
}

// Given a string, return a "rotated left 2" version where the first 2 chars are moved to the end.
// The string length will be at least 2.
func left2(word string) string {
	return word[2:] + word[:2]
}
