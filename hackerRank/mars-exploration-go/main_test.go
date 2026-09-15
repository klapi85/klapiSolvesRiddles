package main

import "testing"

func TestCheckLetter(t *testing.T) {

	tests := []struct {
		name     string
		a, b     string
		expected int
	}{
		{name: "Comparing 2xS", a: "S", b: "S", expected: 0},
		{name: "Comparing 2xO", a: "O", b: "O", expected: 0},
		{name: "Comparing S&O", a: "S", b: "O", expected: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkLetter(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("checkLetter(\"%s\", \"%s\") = %d; expected %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestMarsExploration(t *testing.T) {

	tests := []struct {
		name     string
		a        string
		expected int
	}{
		{name: "Checking SOS", a: "SOS", expected: 0},
		{name: "Checking SXS", a: "SXS", expected: 1},
		{name: "Checking XXX", a: "XXX", expected: 3},
		{name: "Checking SOSSOS", a: "SOSSOS", expected: 0},
		{name: "Checking XXXYYY", a: "XXXYYY", expected: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := marsExploration(tt.a)
			if got != tt.expected {
				t.Errorf("marsExploration(\"%s\") = %d; expected %d", tt.a, got, tt.expected)
			}
		})
	}
}
