package shared

import (
	"fmt"
	"math/big"

	"github.com/bbetov/euler/shared/BitSet"
)

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

func SqrtBigInt(i *big.Int) *big.Int {
	var next, prev *big.Int
	one := big.NewInt(1)
	zero := big.NewInt(0)
	prev.Rsh(i, 1) // Start with half the number
	for prev.Cmp(zero) > 0 {
		next.Div(i, prev)
		next.Add(next, prev)
		next.Rsh(next, 1)

		var diff *big.Int
		if prev.Cmp(next) > 0 {
			diff.Set(prev)
			diff.Sub(diff, next)
		} else {
			diff.Set(next)
			diff.Sub(diff, prev)
		}

		if diff.Cmp(one) <= 0 {
			return next
		}
		prev = next
	}
	return zero
}

const (
	sieveSize uint32 = 10000000
)

func GetPrimesInt64(maxPrime uint64) []uint64 {
	segSize := uint64(sieveSize)
	if segSize > maxPrime {
		segSize = maxPrime
	}

	smallPrimes := BitSet.New(uint32(segSize), true)
	for i := uint64(2); i*i < segSize; i++ {
		if smallPrimes.IsSet(uint32(i)) {
			for j := i * i; j <= segSize; j += i {
				smallPrimes.Set(uint32(j), false)
			}
		}
	}

	var primes []uint64
	for i := uint32(2); i < uint32(segSize); i++ {
		if smallPrimes.IsSet(i) {
			primes = append(primes, uint64(i))
		}
	}

	iter := uint64(1)
	numIters := 1 + maxPrime/segSize
	for offset := segSize; offset < maxPrime; offset += segSize {
		fmt.Printf("Progress: %d/%d [%d]\n", iter, numIters, len(primes))
		iter++
		smallPrimes.Reset(true)
		for _, p := range primes {
			r := offset % p
			if r > 0 {
				r = p - r
			}
			for i := r; i < segSize; i += p {
				smallPrimes.Set(uint32(i), false)
			}
		}
		for i := uint64(0); i < segSize; i++ {
			if smallPrimes.IsSet(uint32(i)) {
				primes = append(primes, offset+uint64(i))
			}
		}
	}
	return primes
}
