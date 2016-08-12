package shared

import "github.com/bbetov/euler/shared/primer"

// SqrtUInt64 computes the integer square root using https://en.wikipedia.org/wiki/Integer_square_root
func SqrtUInt64(i uint64) uint64 {
	var next, prev uint64
	prev = i >> 1 // Start with half the number
	for prev > 0 {
		next = (i/prev + prev)
		next = next >> 1

		var diff uint64
		if prev > next {
			diff = prev - next
		} else {
			diff = next - prev
		}

		if diff <= 1 {
			return next
		}
		prev = next
	}
	return 0
}

//GetPrimesUInt64 returns a list of prime numbers smaller than maxPrime
func GetPrimesUInt64(maxPrime uint64) []uint64 {
	if maxPrime == 2 {
		return []uint64{2}
	}
	if maxPrime == 3 {
		return []uint64{2, 3}
	}

	pr := primer.NewPrimer()
	var primes []uint64

	for prime, _ := pr.NextPrime(); prime <= maxPrime; prime, _ = pr.NextPrime() {
		primes = append(primes, prime)
	}

	return primes
}

//GetPrimeUInt64 returns the N-th prime number
func GetPrimeUInt64(N int) uint64 {
	N--

	pr := primer.NewPrimer()
	prime, _ := pr.NextPrime()
	for N > 0 {
		N--
		prime, _ = pr.NextPrime()
	}
	return prime
}
