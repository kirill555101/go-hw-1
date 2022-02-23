package main

import "testing"

var testsCalc = []struct {
	polishStr string
	expected float64
}{
	{"1 2 +", 3},
	{"10 5 -", 5},
	{"2 2 *", 4},
	{"6 3 /", 2},
	{"69 0 /", 0},
}

func TestCalculate(t *testing.T) {
	for _, e := range testsCalc {
		result, err := Calculate(e.polishStr)

		if result != e.expected && err != nil {
			t.Fatalf("Test has been failed")
		}
	}
}

var testsPolishNotation = []struct {
	str string
	expected string
}{
	{"(8+2*5)/(1+3*2-4)", "8 2 5 * + 1 3 2 * + 4 - /"},
	{"(8PLUS2*5)/(1+3*2-4)", ""},
}

func TestPolishNotation(t *testing.T) {
	for _, e := range testsPolishNotation {
		result, err := GetPostfixNotation(e.str)

		if result != e.expected && err != nil {
			t.Fatalf("Test has been failed")
		}
	}
}
