package main

import "github.com/bbetov/euler/shared"

func main() {
	i1 := shared.Integer{}
	i1.Set("1")
	i2 := shared.Integer{}
	i2.Set("1")
	it := 2
	for i2.NumDigits() < 1000 {
		i1.Add(&i2)
		it++
		i1.Swap(&i2)
	}
	println(it)
}
