package main

import "testing"

func TestFibonacci(t *testing.T) {
	cases := []struct {
		input int
		expectedOutput string
	} {
		{1, "0 "},
		{2, "0 1 "},
		{3, "0 1 1 "},
		{8, "0 1 1 2 3 5 8 13 "},
	}
	for _, c := range cases {
		actual := Fib(c.input)
		if actual != c.expectedOutput {
			t.Errorf("fib(%d) == %q, but expected %q", c.input, actual, c.expectedOutput)
		}
	}
}