package main

import "testing"

func TestNewton(t *testing.T) {
	cases := []struct {
		input float64
		expectedOutput float64
		expectedError bool
	} {
		{25.0, 5.0, false},
		{-1.0, 1.0, true},
	}
	for _, c := range cases {
		actual, err := Sqrt(c.input)
		actualError := err != nil
		if actual != c.expectedOutput || actualError != c.expectedError {
			t.Errorf("Sqrt(%v) == %v, but expected %v", c.input, actual, c.expectedOutput)
		}
	}
}