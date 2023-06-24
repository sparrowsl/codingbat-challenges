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
