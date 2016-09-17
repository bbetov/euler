package shared

import "testing"

func AAABigIntMultiplication(t *testing.T) {
	b1 := Integer{}
	b1.Set("9987")

	b2 := Integer{}
	b2.Set("99")

	b1.Multiply(&b2)
	b2.Set("33")

	b1.Multiply(&b2)

	println(b1.String())

	b1.Set("1234321")
	println(b1.IsPalindrome())
}

var palindromeTests = []struct {
	i string
	a bool
}{
	{"1234321", true},
	{"12344321", true},
	{"121", true},
	{"123444", false},
	{"12354321", false},
}

func TestIsPalindrome(t *testing.T) {
	in := Integer{}
	for _, tt := range palindromeTests {
		in.Set(tt.i)
		v := in.IsPalindrome()
		if v != tt.a {
			t.Error("Expected ", tt.a, " got ", v)
		}
	}
}
