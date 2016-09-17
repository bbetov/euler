package main

import (
	"strconv"

	"github.com/bbetov/euler/shared"
)

func isLychrelNumber(num int) bool {
	i1 := shared.Integer{}
	i2 := shared.Integer{}
	i1.Set(strconv.Itoa(num))
	for iter := 1; iter <= 50; iter++ {
		i2.Set(i1.String())
		i2.Reverse()
		i1.Add(&i2)
		if i1.IsPalindrome() {
			return false
		}
	}
	return true
}

func getNumLychrelNumbers(upper int) (numLychrel int) {
	for i := 1; i < upper; i++ {
		if isLychrelNumber(i) {
			numLychrel++
		}
	}
	return numLychrel
}

func main() {
	println(getNumLychrelNumbers(10000))
}
