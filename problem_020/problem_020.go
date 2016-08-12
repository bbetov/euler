package main

import (
	"strconv"

	"github.com/bbetov/euler/shared"
)

// Thou shalt not use big/Int
func main() {
	r := shared.Integer{}
	r.Set("2")
	tmp := shared.Integer{}
	for i := 3; i <= 100; i++ {
		tmp.Set(strconv.Itoa(i))
		r.Multiply(&tmp)
	}
	println(r.SumDigits())
	//println(r.String())
}
