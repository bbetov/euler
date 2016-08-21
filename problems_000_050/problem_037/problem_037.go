package main

import "github.com/bbetov/euler/shared/primer"

var smallerPrimes = make(map[uint64]bool)

func isTruncatable(p uint64) bool {
	digs := []uint64{}
	for p > 0 {
		d := p % 10
		digs = append(digs, d)
		p /= 10
		if _, ok := smallerPrimes[p]; p > 0 && !ok {
			return false
		}
	}
	mul := uint64(1)
	current := uint64(0)
	for i := 0; i < len(digs)-1; i++ {
		current += digs[i] * mul
		mul *= 10
		if _, ok := smallerPrimes[current]; !ok {
			return false
		}
	}

	return true
}

func main() {
	p := primer.NewPrimer()
	// Skip up to 10
	prime, _ := p.NextPrime()
	for prime < 10 {
		smallerPrimes[prime] = true
		prime, _ = p.NextPrime()
	}

	cntr := 0
	var sum uint64
	var err error
	for cntr < 11 {
		if isTruncatable(prime) {
			sum += prime
			cntr++
			println(prime)
		}
		smallerPrimes[prime] = true
		prime, err = p.NextPrime()
		if err != nil {
			println("Out of uint64 range!")
			break
		}
	}
	println(sum)
}
