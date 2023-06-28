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

func Test_makeOutWord(t *testing.T) {
	testCases := []struct {
		out, word, expected string
	}{
		{"<<>>", "Yay", "<<Yay>>"},
		{"<<>>", "WooHoo", "<<WooHoo>>"},
		{"[[]]", "word", "[[word]]"},
	}

	for _, test := range testCases {
		result := makeOutWord(test.out, test.word)

		if result != test.expected {
			t.Errorf("makeOutWord(%q, %q) FAILED. Expected %q, got %q", test.out, test.word, test.expected, result)
		}
	}
}

func Test_extraEnd(t *testing.T) {
	testCases := []struct {
		word, expected string
	}{
		{"Hello", "lololo"},
		{"ab", "ababab"},
		{"Hi", "HiHiHi"},
	}

	for _, test := range testCases {
		result := extraEnd(test.word)

		if result != test.expected {
			t.Errorf("extraEnd(%q) FAILED. Expected %q, got %q", test.word, test.expected, result)
		}
	}
}

func Test_firstTwo(t *testing.T) {
	testCases := []struct {
		word, expected string
	}{
		{"Hello", "He"},
		{"abcdefg", "ab"},
		{"ab", "ab"},
	}

	for _, test := range testCases {
		result := firstTwo(test.word)

		if result != test.expected {
			t.Errorf("firstTwo(%q) FAILED. Expected %q, got %q", test.word, test.expected, result)
		}
	}
}

func Test_firstHalf(t *testing.T) {
	testCases := []struct {
		word, expected string
	}{
		{"WooHoo", "Woo"},
		{"HelloThere", "Hello"},
		{"abcdef", "abc"},
	}

	for _, test := range testCases {
		result := firstHalf(test.word)

		if result != test.expected {
			t.Errorf("firstHalf(%q) FAILED. Expected %q, got %q", test.word, test.expected, result)
		}
	}
}

func Test_withoutEnd(t *testing.T) {
	testCases := []struct {
		word, expected string
	}{
		{"Java", "av"},
		{"Hello", "ell"},
		{"coding", "odin"},
	}

	for _, test := range testCases {
		result := withoutEnd(test.word)

		if result != test.expected {
			t.Errorf("withoutEnd(%q) FAILED. Expected %q, got %q", test.word, test.expected, result)
		}
	}
}

func Test_comboString(t *testing.T) {
	testCases := []struct {
		a, b, expected string
	}{
		{"hi", "Hello", "hiHellohi"},
		{"Hello", "hi", "hiHellohi"},
		{"aaa", "b", "baaab"},
	}

	for _, test := range testCases {
		result := comboString(test.a, test.b)

		if result != test.expected {
			t.Errorf("comboString(%q, %q) FAILED. Expected %q, got %q", test.a, test.b, test.expected, result)
		}
	}
}

func Test_nonStart(t *testing.T) {
	testCases := []struct {
		a, b, expected string
	}{
		{"Hello", "There", "ellohere"},
		{"java", "code", "avaode"},
		{"shotl", "java", "hotlava"},
	}

	for _, test := range testCases {
		result := nonStart(test.a, test.b)

		if result != test.expected {
			t.Errorf("nonStart(%q, %q) FAILED. Expected %q, got %q", test.a, test.b, test.expected, result)
		}
	}
}

func Test_left2(t *testing.T) {
	testCases := []struct {
		word, expected string
	}{
		{"Hello", "lloHe"},
		{"java", "vaja"},
		{"Hi", "Hi"},
	}

	for _, test := range testCases {
		result := left2(test.word)

		if result != test.expected {
			t.Errorf("left2(%q) FAILED. Expected %q, got %q", test.word, test.expected, result)
		}
	}
}
