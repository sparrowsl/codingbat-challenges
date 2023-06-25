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

func TestFrontTimes(t *testing.T) {
	testCases := []struct {
		word     string
		n        uint
		expected string
	}{
		{"Chocolate", 2, "ChoCho"},
		{"Chocolate", 3, "ChoChoCho"},
		{"Abc", 3, "AbcAbcAbc"},
	}

	for _, test := range testCases {
		result := FrontTimes(test.word, test.n)

		if result != test.expected {
			t.Errorf("FrontTimes(%q, %v) FAILED. Expected %v, got %v", test.word, test.n, test.expected, result)
		}
	}
}

func TestStringBits(t *testing.T) {
	testCases := []struct {
		word, expected string
	}{
		{"Hello", "Hlo"},
		{"Hi", "H"},
		{"Heeololeo", "Hello"},
	}

	for _, test := range testCases {
		result := StringBits(test.word)

		if result != test.expected {
			t.Errorf("FrontTimes(%q) FAILED. Expected %q, got %q", test.word, test.expected, result)
		}
	}
}

func Test_arrayCount(t *testing.T) {
	testCases := []struct {
		array    []int
		expected int
	}{
		{[]int{1, 2, 9}, 1},
		{[]int{1, 9, 9}, 2},
		{[]int{1, 9, 9, 3, 9}, 3},
		{[]int{}, 0},
	}

	for _, test := range testCases {
		result := arrayCount(test.array)

		if result != test.expected {
			t.Errorf("arrayCount(%v) FAILED. Expected %v, got %v", test.array, test.expected, result)
		}
	}
}

func Test_arrayFront(t *testing.T) {
	testCases := []struct {
		array    []int
		expected bool
	}{
		{[]int{1, 2, 9, 3, 4}, true},
		{[]int{1, 2, 3, 4, 9}, false},
		{[]int{1, 2, 3, 4, 5}, false},
	}

	for _, test := range testCases {
		result := arrayFront(test.array)

		if result != test.expected {
			t.Errorf("arrayFront(%v) FAILED. Expected %v, got %v", test.array, test.expected, result)
		}
	}
}
