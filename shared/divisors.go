package shared

func GetDivisorsFreq(n uint64) (s map[uint64]int) {
	primes := GetPrimesInt64(n)

	// find the prime divisors
	s = make(map[uint64]int)
	for _, v := range primes {
		tmp := n
		for tmp%v == 0 {
			s[v]++
			tmp /= v
		}
	}
	//fmt.Printf("> %d ==> %v\n", n, pds)
	return s
}

func GetDivisorsSum(n uint64) (s uint64) {
	if n == 1 {
		return 1
	}
	// find the prime divisors
	pds := GetDivisorsFreq(n)

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
