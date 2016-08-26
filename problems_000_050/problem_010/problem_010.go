package main

import "github.com/bbetov/euler/shared"

func getPrimesSum(limit uint64) (sum uint64) {
	primes := shared.GetPrimesUInt64(limit)
	for _, p := range primes {
		sum += p
	}
	return sum
}

func main() {
	println(getPrimesSum(2000000))
}
