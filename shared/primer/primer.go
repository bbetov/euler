package primer

import (
	"errors"

	"github.com/bbetov/euler/shared/bitset"
)

type primer struct {
	primes     []uint64
	primeIndex int
	segSize    uint32
	segIndex   int
}

type Primer interface {
	NextPrime() (uint64, error)
}

const maxUint64 = ^uint64(0)

func (p *primer) NextPrime() (n uint64, err error) {
	p.primeIndex++
	if p.primeIndex >= len(p.primes) {
		for {
			n, err := p.fillNext()
			if err != nil {
				return 0, err
			}
			if n > 0 {
				break
			}
		}
	}
	return p.primes[p.primeIndex], nil
}

func (p *primer) fillNext() (int, error) {
	var cnt int
	b := bitset.New(p.segSize+1, true)
	low := uint64(p.segIndex)*uint64(p.segSize)*2 + 3
	if maxUint64-uint64(p.segSize)*2 < low {
		return 0, errors.New("At the end of the uint64 range")
	}
	p.segIndex++
	high := uint64(p.segIndex)*uint64(p.segSize)*2 + 3
	for i := uint32(1); /*skip 2*/ i < uint32(len(p.primes)); i++ {
		start := (low / p.primes[i]) * p.primes[i]
		if start < low {
			start += p.primes[i]
		}

		for j := start; j < high; j += p.primes[i] {
			if j%2 == 0 {
				continue
			}
			idx := uint32((j - low + 1) / 2)
			b.Set(idx, false)
		}
	}

	for i := uint32(0); i < p.segSize; i++ {
		if b.IsSet(i) {
			p.primes = append(p.primes, low+uint64(i)*2)
			cnt++
		}
	}

	return cnt, nil
}

func (p *primer) initialize(segSize uint32) {
	p.segSize = segSize
	p.segIndex = 1
	p.primeIndex = -1
	p.primes = append(p.primes, 2)
	// Generate small primes
	b := bitset.New(p.segSize+1, true)
	maxNum := p.segSize*2 + 3
	for i := uint32(3); i*i <= maxNum; i += 2 {
		idx := (i - 3) / 2
		if b.IsSet(idx) {
			for j := i * i; j <= maxNum; j += i {
				if j%2 == 0 {
					continue
				}
				idx := (j - 3) / 2
				b.Set(idx, false)
			}
		}
	}

	for i := uint32(0); i < p.segSize; i++ {
		if b.IsSet(i) {
			p.primes = append(p.primes, uint64(i)*2+3)
		}
	}
}

func NewPrimer() Primer {
	p := &primer{}
	p.initialize(32768)

	return p
}

func NewPrimerWithSegSize(segSize uint32) Primer {
	p := &primer{}
	p.initialize(segSize)

	return p
}
