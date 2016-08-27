package main

import (
	"strconv"

	"github.com/bbetov/euler/shared"
)

// Thou shalt not use big/Int
func getSumDigitsFact(num int) uint64 {
	r := shared.Integer{}
	r.Set(strconv.Itoa(2))
	tmp := shared.Integer{}
	for i := 3; i <= num; i++ {
		tmp.Set(strconv.Itoa(i))
		r.Multiply(&tmp)
	}
	return r.SumDigits()
}

func main() {
	println(getSumDigitsFact(100))
}
