package primer

import (
	"errors"

	"github.com/bbetov/euler/shared/BitSet"
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
		if err := p.fillNext(); err != nil {
			return 0, err
		}
	}
	return p.primes[p.primeIndex], nil
}

func (p *primer) fillNext() error {
	b := BitSet.New(p.segSize+1, true)
	low := uint64(p.segIndex)*uint64(p.segSize) + 1
	if maxUint64-uint64(p.segSize) < low {
		return errors.New("At the end of the uint64 range")
	}
	p.segIndex++
	high := uint64(p.segIndex) * uint64(p.segSize)
	for i := uint32(0); i < uint32(len(p.primes)); i++ {
		start := (low / p.primes[i]) * p.primes[i]
		if start < low {
			start += p.primes[i]
		}

		for j := start; j < high; j += p.primes[i] {
			b.Set(uint32(j-low), false)
		}
	}

	for i := low; i < high; i++ {
		if b.IsSet(uint32(i - low)) {
			p.primes = append(p.primes, i)
		}
	}
	return nil
}

func (p *primer) initialize(segSize uint32) {
	p.segSize = segSize
	p.segIndex = 1
	p.primeIndex = -1
	// Generate small primes
	b := BitSet.New(p.segSize+1, true)
	for i := uint32(2); i*i <= p.segSize; i++ {
		if b.IsSet(i) {
			for j := i * i; j <= p.segSize; j += i {
				b.Set(j, false)
			}
		}
	}
	for i := uint32(2); i <= p.segSize; i++ {
		if b.IsSet(i) {
			p.primes = append(p.primes, uint64(i))
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
