package main

import "testing"

var tests = []struct {
	n       []int
	numPerm int
	e       []int
}{
	{[]int{0, 1, 2}, 4, []int{1, 2, 0}},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 1000000, []int{2, 7, 8, 3, 9, 1, 5, 4, 6, 0}},
}

func testEqIntSlice(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestNextLexPerm(t *testing.T) {
	for _, tt := range tests {
		var v []int
		var err error
		for i := 1; i < tt.numPerm; i++ {
			v, err = nextLexPerm(tt.n)
			if err != nil {
				t.Error("Expected permutation, but returned error.")
			}
		}
		if !testEqIntSlice(v, tt.e) {
			t.Errorf("Expected %v, got %v", tt.e, v)
		}
	}
}
