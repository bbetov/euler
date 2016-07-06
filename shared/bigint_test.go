package shared

import "testing"

func TestBigIntMultiplication(t *testing.T) {
	b1 := Integer{}
	b1.Set("9987")

	b2 := Integer{}
	b2.Set("99")

	b1.Multiply(&b2)
	b2.Set("33")

	b1.Multiply(&b2)

	println(b1.String())
}
