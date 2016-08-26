// Problem 3
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import (
	"fmt"

	"github.com/bbetov/euler/shared"
)

func getLargestPrimeFactor(n uint64) (maxDiv uint64) {
	lp := shared.SqrtUInt64(n) + 1
	primes := shared.GetPrimesUInt64(lp)
	for _, p := range primes {
		if n%p == uint64(0) && maxDiv < p {
			maxDiv = p
		}
	}
	return maxDiv
}

func main() {
	//fmt.Println(getLargestPrimeFactor(600851475143))
	fmt.Println(getLargestPrimeFactor(42))
}
