package main

import "github.com/bbetov/euler/shared"

const (
	maxK = uint64(10000)
)

func getAmicableSum() (sp uint64) {
	dvsum := make(map[uint64]uint64)
	primes := shared.GetPrimesUInt64(maxK)

	for k := uint64(1); k <= maxK; k++ {
		var s1, s2 uint64
		if val, ok := dvsum[k]; ok {
			s1 = val
		} else {
			s1 = shared.GetDivisorsSumPrimesSet(k, primes)
			dvsum[k] = s1
		}
		if val, ok := dvsum[s1]; ok {
			s2 = val
		} else {
			s2 = shared.GetDivisorsSumPrimesSet(s1, primes)
			dvsum[s1] = s2
		}
		if s2 == k && k != s1 {
			sp += k
		}
	}

	return sp
}

func main() {

	println(getAmicableSum())
}
