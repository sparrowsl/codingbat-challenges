package warmup1

import "testing"

func TestSleepIn(t *testing.T) {
	sleepCases := []struct {
		a, b, c bool
	}{
		{false, false, true},
		{true, false, false},
		{false, true, true},
		{true, true, true},
	}

	for _, v := range sleepCases {
		ans := SleepIn(v.a, v.b)

		if ans != v.c {
			t.Error("For", v, "Expected", v.c, "got", ans)
		}
	}
}

func TestMonkeyTrouble(t *testing.T) {
	monkeyCases := []struct {
		a, b, c bool
	}{
		{true, true, true},
		{false, false, true},
		{true, false, false},
		{false, true, false},
	}

	for _, v := range monkeyCases {
		ans := MonkeyTrouble(v.a, v.b)

		if ans != v.c {
			t.Error("For", v, "Expected", v.c, "got", ans)
		}
	}
}

func TestSumDouble(t *testing.T) {
	doubleCases := []struct {
		a, b, ans int
	}{
		{1, 2, 3},
		{3, 2, 5},
		{2, 2, 8},
		{-1, 0, -1},
		{3, 3, 12},
		{0, 0, 0},
	}

	for _, v := range doubleCases {
		sum := SumDouble(v.a, v.b)

		if sum != v.ans {
			t.Error("For", v, "Expected", v.ans, "got", sum)
		}
	}
}

func TestDiff21(t *testing.T) {
	diffCases := []struct {
		n, ans int
	}{
		{19, 2},
		{10, 11},
		{21, 0},
		{22, 2},
		{25, 8},
		{0, 21},
	}

	for _, diff := range diffCases {
		ans := Diff21(diff.n)

		if ans != diff.ans {
			t.Error("For", diff.n, "Expected", diff.ans, "got", ans)
		}
	}
}

func TestMakes10(t *testing.T) {
	numbers := []struct {
		a, b int
		ans  bool
	}{
		{9, 10, true},
		{9, 9, false},
		{1, 9, true},
	}

	for _, v := range numbers {
		ans := Makes10(v.a, v.b)

		if ans != v.ans {
			t.Error("For", v, "Expected", v.ans, "got", ans)
		}
	}
}

func TestNearHundred(t *testing.T) {
	nums := []struct {
		n   int
		ans bool
	}{
		{93, true},
		{90, true},
		{89, false},
	}

	for _, num := range nums {
		v := NearHundred(num.n)

		if v != num.ans {
			t.Error("For", num, "Expected", num.ans, "got", v)
		}
	}
}

func TestPosNeg(t *testing.T) {
	posNegs := []struct {
		a, b          int
		negative, ans bool
	}{
		{1, -1, false, true},
		{-1, 1, false, true},
		{-4, -5, true, true},
	}

	for _, n := range posNegs {
		v := PosNeg(n.a, n.b, n.negative)

		if v != n.ans {
			t.Error("For", n, "Expected", n.ans, "got", v)
		}
	}
}

func TestNotString(t *testing.T) {
	testCases := []struct {
		value, expected string
	}{
		{"candy", "not candy"},
		{"x", "not x"},
		{"not bad", "not bad"},
	}

	for _, test := range testCases {
		result := NotString(test.value)

		if result != test.expected {
			t.Errorf(`NotString("%v") FAILED. Expected %v, got %v`, test.value, test.expected, result)
		}
	}
}

func TestMissingChar(t *testing.T) {
	testCases := []struct {
		value    string
		index    int
		expected string
	}{
		{"kitten", 1, "ktten"},
		{"kitten", 0, "itten"},
		{"kitten", 4, "kittn"},
	}

	for _, test := range testCases {
		result := MissingChar(test.value, test.index)

		if result != test.expected {
			t.Errorf("MissingChar(%v, %v) FAILED. Expected %v, got %v", test.value, test.index, test.expected, result)
		}
	}
}

func TestFrontBack(t *testing.T) {
	testCases := []struct {
		word, expected string
	}{
		{"code", "eodc"},
		{"a", "a"},
		{"ab", "ba"},
	}

	for _, test := range testCases {
		result := FrontBack(test.word)

		if result != test.expected {
			t.Errorf("FrontBack(%v) FAILED. Expected %v, got %v", test.word, test.expected, result)
		}
	}
}
