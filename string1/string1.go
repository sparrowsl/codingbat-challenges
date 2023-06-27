package string1

import "fmt"

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
