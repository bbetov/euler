package main

import "github.com/bbetov/euler/shared"

func getDivisorsSum(n uint64) (s uint64) {
	if n == 1 {
		return 1
	}
	primes := shared.GetPrimesInt64(n)

	// find the prime divisors
	pds := make(map[uint64]int)
	for _, v := range primes {
		tmp := n
		for tmp%v == 0 {
			pds[v]++
			tmp /= v
		}
	}
	//fmt.Printf("> %d ==> %v\n", n, pds)

	dvs := []uint64{1}
	for k, v := range pds {
		tdvs := []uint64{}
		tv := k
		for i := 1; i <= v; i++ {
			for _, q := range dvs {
				tdvs = append(tdvs, tv*q)
			}
			tv *= k
		}
		dvs = append(dvs, tdvs...)
	}
	//fmt.Printf("> %d ==> %v\n", n, dvs)
	for _, q := range dvs {
		s += q
	}
	return s - n
}

func main() {
	dvsum := make(map[uint64]uint64)
	var sp uint64

	for k := uint64(1); k <= 10000; k++ {
		var s1, s2 uint64
		if val, ok := dvsum[k]; ok {
			s1 = val
		} else {
			s1 = getDivisorsSum(k)
			dvsum[k] = s1
		}
		if val, ok := dvsum[s1]; ok {
			s2 = val
		} else {
			s2 = getDivisorsSum(s1)
			dvsum[s1] = s2
		}
		if s2 == k && k != s1 {
			sp += k
		}
	}

	println(sp)
}
