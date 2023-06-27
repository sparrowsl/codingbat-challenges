package string1

import "testing"

func Test_helloName(t *testing.T) {
	testCases := []struct {
		name, expected string
	}{
		{"Bob", "Hello Bob!"},
		{"Alice", "Hello Alice!"},
		{"X", "Hello X!"},
	}

	for _, test := range testCases {
		result := helloName(test.name)

		if result != test.expected {
			t.Errorf("helloName(%q) FAILED. Expected %q, got %q", test.name, test.expected, result)
		}
	}
}

func Test_makeAbba(t *testing.T) {
	testCases := []struct {
		a, b, expected string
	}{
		{"Hi", "Bye", "HiByeByeHi"},
		{"Yo", "Alice", "YoAliceAliceYo"},
		{"What", "Up", "WhatUpUpWhat"},
	}

	for _, test := range testCases {
		result := makeAbba(test.a, test.b)

		if result != test.expected {
			t.Errorf("makeAbba(%q, %q) FAILED. Expected %q, got %q", test.a, test.b, test.expected, result)
		}
	}
}

func Test_makeTags(t *testing.T) {
	testCases := []struct {
		tag, word, expected string
	}{
		{"i", "Yay", "<i>Yay</i>"},
		{"i", "Hello", "<i>Hello</i>"},
		{"cite", "Yay", "<cite>Yay</cite>"},
	}

	for _, test := range testCases {
		result := makeTags(test.tag, test.word)

		if result != test.expected {
			t.Errorf("makeAbba(%q, %q) FAILED. Expected %q, got %q", test.tag, test.word, test.expected, result)
		}
	}
}
