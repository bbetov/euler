package main

import "github.com/bbetov/euler/shared"

func copy(a []byte) []byte {
	rv := []byte{}
	for _, v := range a {
		rv = append(rv, v)
	}
	return rv
}

func cmpDig(a, b []byte) bool {
	// make a copy of b
	b = copy(b)
	for _, va := range a {
		found := false
		for i, vb := range b {
			if vb == va {
				found = true
				b[i] = 255
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func decompose(n uint64) []byte {
	rv := []byte{}
	for n > 0 {
		r := n % 10
		n /= 10
		rv = append(rv, byte(r))
	}
	return rv
}

func main() {
	tmpp := shared.GetPrimesUInt64(uint64(10000))
	primes := make(map[uint64][]byte)
	plist := []uint64{}
	// Filter all 4 digits primes
	for _, t := range tmpp {
		if t > 1000 {
			plist = append(plist, t)
			primes[t] = decompose(t)
		}
	}

	for _, p := range plist {
		// find the primes with the same digits
		var cdig []byte
		var ok bool
		if cdig, ok = primes[p]; !ok {
			continue // Already processed
		}
		cur := []uint64{p}
		prev := p
		for _, pr := range plist {
			if pr == p {
				continue
			}
			if dig, ok := primes[pr]; ok {
				if cmpDig(dig, cdig) && pr-prev == 3330 {
					cur = append(cur, pr)
					prev = pr
				}
			}
		}
		for _, c := range cur {
			delete(primes, c)
		}
		if len(cur) == 3 && cur[0] != 1487 {
			println(cur[0]*100000000 + cur[1]*10000 + cur[2])
		}
	}

}
