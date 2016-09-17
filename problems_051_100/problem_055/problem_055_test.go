package main

import "testing"

var tests = []struct {
	input           int
	isLychrelNumber bool
}{
	{121, false},
	{349, false},
	{4994, true},
}

func TestIsLychrelNumber(t *testing.T) {
	for _, tt := range tests {
		v := isLychrelNumber(tt.input)
		if v != tt.isLychrelNumber {
			if tt.isLychrelNumber {
				t.Error("Expected ", tt.input, " to be flagged as Lychrel number, but it was not.")
			} else {
				t.Error("Expected ", tt.input, " to not be flagged as Lychrel number, but it was.")
			}
		}
	}
}

func TestNumLBelow10k(t *testing.T) {
	v := getNumLychrelNumbers(10000)
	if v != 249 {
		t.Error("Number of Lychrel numbers below 10,000 shoud be ", 249)
	}
}
