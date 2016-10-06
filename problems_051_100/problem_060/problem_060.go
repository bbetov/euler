package main

import (
	"fmt"

	"github.com/bbetov/euler/shared/primer"
)

var powers []uint64
var primesMap = make(map[uint64]bool)

func init() {
	tmp := uint64(1)
	for i := uint64(0); i < 19; i++ {
		powers = append(powers, tmp)
		tmp *= 10
	}
}

func getNumDigits(n uint64) int {
	prev := 0
	for i := 0; i < len(powers); i++ {
		if n > powers[i] {
			continue
		}
		prev = i
		break
	}
	return prev
}

func compatiblePrimes(a, b uint64) bool {
	if a == b {
		return false
	}
	nda := getNumDigits(a)
	ndb := getNumDigits(b)
	ab := powers[nda]*b + a
	ba := powers[ndb]*a + b
	if _, ok := primesMap[ab]; !ok {
		return false
	}
	if _, ok := primesMap[ba]; !ok {
		return false
	}
	return true
}

func checkForSolution(st []map[uint64]uint32, i uint32, primes []uint64) (rv uint64, foundSolution bool) {
	sl := []uint64{}
	sli := []uint32{}
	for k, v := range st[i] {
		sl = append(sl, k)
		sli = append(sli, v)
	}
	sllen := len(sl)
	var ok bool
	for a := 0; a < sllen; a++ {
		for b := a + 1; b < sllen; b++ {
			if _, ok = st[sli[a]][sl[b]]; !ok {
				continue
			}
			for c := b + 1; c < sllen; c++ {
				// ac, ab, bc
				if _, ok = st[sli[a]][sl[c]]; !ok {
					continue
				}
				if _, ok = st[sli[b]][sl[c]]; !ok {
					continue
				}
				for d := c + 1; d < sllen; d++ {
					if _, ok = st[sli[a]][sl[d]]; !ok {
						continue
					}
					if _, ok = st[sli[b]][sl[d]]; !ok {
						continue
					}
					if _, ok = st[sli[c]][sl[d]]; !ok {
						continue
					}
					//fmt.Printf("Primes: %v, %v, %v, %v, %v\n", primes[i], sl[a], sl[b], sl[c], sl[d])
					return primes[i] + sl[a] + sl[b] + sl[c] + sl[d], true
				}
			}
		}
	}
	return 0, false
}

func main() {
	var st []map[uint64]uint32
	pr := primer.NewPrimer()
	var primes []uint64
	var prIndex uint32
	var p uint64
	for i := 0; i < 2; i++ { // seed
		p, _ = pr.NextPrime()
		primes = append(primes, p)
		primesMap[p] = true
	}
	for {
		foundSolution := false
		nd := getNumDigits(primes[prIndex+1])
		for p < powers[nd]*primes[prIndex]+primes[prIndex+1] {
			p, _ = pr.NextPrime()
			primes = append(primes, p)
			primesMap[p] = true
		}
		st = append(st, make(map[uint64]uint32))
		currentPrime := primes[prIndex]
		for i := uint32(0); i < prIndex; i++ {
			if compatiblePrimes(primes[i], currentPrime) {
				st[i][currentPrime] = prIndex
				st[prIndex][primes[i]] = i
				if len(st[i]) >= 4 {
					if sol, ok := checkForSolution(st, i, primes); ok {
						foundSolution = true
						fmt.Printf("The sum is %v.", sol)
						break
					}
				}
			}
		}
		if foundSolution {
			break
		}
		prIndex++
	}

}
