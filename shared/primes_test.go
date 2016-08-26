package shared

import "testing"

func TestSqrtUInt64(t *testing.T) {
	var tests = []struct {
		in  uint64
		out uint64
	}{
		{81, 9},
		{625, 25},
		{60, 7},
		{12345, 111},
		{987654321, 31426},
		{2, 1},
	}

	for _, c := range tests {
		v := SqrtUInt64(c.in)
		if v != c.out {
			t.Errorf("Expected sqrt(%v) to equal %v, but got %v", c.in, c.out, v)
		}
	}
}
