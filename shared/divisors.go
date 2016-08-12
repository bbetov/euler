package shared

import "github.com/bbetov/euler/shared/primer"

func GetDivisorsFreq(n uint64) (s map[uint64]int) {
	primes := primer.NewPrimer()

	// find the prime divisors
	s = make(map[uint64]int)
	prime, _ := primes.NextPrime()
	for n > prime {
		tmp := n
		for tmp%prime == 0 {
			s[prime]++
			tmp /= prime
		}
		prime, _ = primes.NextPrime()
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

func GetDivisorsFreqPrimesSet(n uint64, primes []uint64) (s map[uint64]int) {
	// find the prime divisors
	s = make(map[uint64]int)
	for _, prime := range primes {
		tmp := n
		for tmp%prime == 0 {
			s[prime]++
			tmp /= prime
		}
	}
	//fmt.Printf("> %d ==> %v\n", n, pds)
	return s
}

func GetDivisorsSumPrimesSet(n uint64, primes []uint64) (s uint64) {
	if n == 1 {
		return 1
	}
	// find the prime divisors
	pds := GetDivisorsFreqPrimesSet(n, primes)

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
