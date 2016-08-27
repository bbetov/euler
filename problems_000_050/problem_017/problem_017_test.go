package main

import (
	"strings"
	"testing"
)

var tests = []struct {
	n        int
	expected int
}{
	{342, 23},
	{115, 20},
}

func TestNumberInWords(t *testing.T) {
	for _, tt := range tests {
		v := numberInWords(tt.n)
		v = strings.Replace(v, " ", "", -1)
		if len(v) != tt.expected {
			t.Error("Expected ", tt.expected, " got ", v)
		}
	}
}

func TestOneToThousand(t *testing.T) {
	v := onetoThousandNumLetters()
	if v != 21124 {
		t.Error("Expected ", 21124, " got ", v)
	}
}
