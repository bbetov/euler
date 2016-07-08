package main

import "github.com/bbetov/euler/shared"

func main() {
	dvsum := make(map[uint64]uint64)
	var sp uint64

	for k := uint64(1); k <= 10000; k++ {
		var s1, s2 uint64
		if val, ok := dvsum[k]; ok {
			s1 = val
		} else {
			s1 = shared.GetDivisorsSum(k)
			dvsum[k] = s1
		}
		if val, ok := dvsum[s1]; ok {
			s2 = val
		} else {
			s2 = shared.GetDivisorsSum(s1)
			dvsum[s1] = s2
		}
		if s2 == k && k != s1 {
			sp += k
		}
	}

	println(sp)
}
