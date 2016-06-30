package main

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

import "github.com/bbetov/euler/shared"

func factorize(num int, primes []uint64) map[int]int {
	rv := make(map[int]int)
	for _, p := range primes {
		for num > 1 && num%int(p) == 0 {
			rv[int(p)]++
			num /= int(p)
		}
		if num == 1 {
			break
		}
	}
	return rv
}

func evenlyDivisible(maxNum int) int {
	primes := shared.GetPrimesInt64(uint64(maxNum))
	nums := make(map[int]int)
	for i := maxNum; i > 0; i-- {
		f := factorize(i, primes)
		for k, v := range f {
			oldv := nums[k]
			if oldv < v {
				nums[k] = v
			}
		}
	}

	rv := 1
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
