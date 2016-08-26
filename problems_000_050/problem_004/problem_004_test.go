package main

import "testing"

var tests = []struct {
	n        int  // input
	expected bool // expected result
}{
	{1234, false},
	{1234321, true},
	{1987891, true},
	{87698, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range tests {
		v := isPalindrome(tt.n)
		if v != tt.expected {
			t.Error("Expected ", tt.expected, " got ", v)
		}
	}
}

var tests2 = []struct {
	n        int
	expected int
}{
	{100, 9009},
	{1000, 906609},
}

func TestGetMaxPalindrome(t *testing.T) {
	for _, tt := range tests2 {
		v := getMaxPalindrome(tt.n)
		if v != tt.expected {
			t.Error("Expected ", tt.expected, " got ", v)
		}
	}
}
