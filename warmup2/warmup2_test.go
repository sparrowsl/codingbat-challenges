package warmup2

import "testing"

func TestStringTimes(t *testing.T) {
	testCases := []struct {
		word     string
		n        uint
		expected string
	}{
		{"Hi", 2, "HiHi"},
		{"Hi", 3, "HiHiHi"},
		{"Hi", 1, "Hi"},
	}

	for _, test := range testCases {
		result := StringTimes(test.word, test.n)

		if result != test.expected {
			t.Errorf("StringTimes(%q, %v) FAILED. Expected %v, got %v", test.word, test.n, test.expected, result)
		}
	}
}

func TestFrontTimes(t *testing.T){
	testCases := []struct {
		word     string
		n        uint
		expected string
	}{
		{"Chocolate", 2, "ChoCho"},
		{"Chocolate", 3, "ChoChoCho"},
		{"Abc", 3, "AbcAbcAbc"},
	}

	for _,test := range testCases{
		result := FrontTimes(test.word,test.n)

				if result != test.expected {
			t.Errorf("FrontTimes(%q, %v) FAILED. Expected %v, got %v", test.word, test.n, test.expected, result)
		}
	}
}
