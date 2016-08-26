package main

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

import "github.com/bbetov/euler/shared"

func evenlyDivisible(maxNum int) uint64 {
	primes := shared.GetPrimesUInt64(uint64(maxNum))
	nums := make(map[uint64]int)
	for i := maxNum; i > 0; i-- {
		f := shared.GetDivisorsFreqPrimesSet(uint64(i), primes)
		for k, v := range f {
			oldv := nums[k]
			if oldv < v {
				nums[k] = v
			}
		}
	}

	rv := uint64(1)
	for k, v := range nums {
		for i := 0; i < v; i++ {
			rv *= k
		}
	}
	return rv
}

func main() {
	println(evenlyDivisible(20))
}
