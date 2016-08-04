package main

import "github.com/bbetov/euler/shared"

func checkCircular(p int, m map[int]bool) bool {
	if p < 10 {
		return true
	}
	digits := []int{}
	for p > 0 {
		d := p % 10
		digits = append(digits, d)
		p /= 10
	}
	for i := 0; i < len(digits)-1; i++ {
		// rotate
		fd := digits[0]
		for d := 0; d < len(digits)-1; d++ {
			digits[d] = digits[d+1]
		}
		digits[len(digits)-1] = fd

		// check
		num, ten := 0, 1
		for d := 0; d < len(digits); d++ {
			num += digits[d] * ten
			ten *= 10
		}
		if _, ok := m[num]; !ok {
			return false
		}
	}
	return true
}

func main() {
	// get all primes up to 1M
	p := shared.GetPrimesInt64(uint64(1000001))
	m := make(map[int]bool)
	var cnt int
	for _, v := range p {
		m[int(v)] = true
	}
	for _, v := range p {
		if checkCircular(int(v), m) {
			cnt++
		}
	}
	println(cnt)
}
