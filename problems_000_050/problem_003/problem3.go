// Problem 3
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import (
	"fmt"

	"github.com/bbetov/euler/shared"
)

func main() {
	limit := uint64(600851475143)
	lp := shared.SqrtUInt64(limit)
	primes := shared.GetPrimesUInt64(lp)
	maxDiv := uint64(0)
	for _, p := range primes {
		if limit%p == uint64(0) && maxDiv < p {
			maxDiv = p
		}
	}
	fmt.Println(maxDiv)
}
