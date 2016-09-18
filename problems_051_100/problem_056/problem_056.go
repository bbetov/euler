package main

import (
	"strconv"

	"github.com/bbetov/euler/shared"
)

func getMaxSumDigits() (max uint64) {
	for a := 2; a <= 100; a++ {
		for b := 1; b <= 100; b++ {
			res := shared.Integer{}
			res.Set("1")
			num := shared.Integer{}
			num.Set(strconv.Itoa(a))
			e := b
			for e > 0 {
				if e&1 == 1 {
					res.Multiply(&num)
				}
				e >>= 1
				num2 := shared.Integer{}
				num2.Set(num.String())
				num.Multiply(&num2)
			}
			sumDig := res.SumDigits()
			if sumDig > max {
				max = sumDig
			}
		}
	}
	return max
}

func main() {
	println(getMaxSumDigits())
}
