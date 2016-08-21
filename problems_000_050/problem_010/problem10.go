package main

import "github.com/bbetov/euler/shared"

func main() {
	primes := shared.GetPrimesUInt64(2000000)
	sum := uint64(0)
	for _, p := range primes {
		sum += p
	}
	println(sum)
}
