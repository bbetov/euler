package main

func num2Array(p uint64) (rv map[byte]int) {
	rv = make(map[byte]int)
	for p > 0 {
		r := byte(p % 10)
		p /= 10
		rv[byte(r)]++
	}
	return rv
}

func mapsEqual(a, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if v2, ok := b[k]; !ok || v2 != v {
			return false
		}
	}
	return true
}

func findPermutedMult(mult uint64) (pm uint64) {
	for i := uint64(101); i < 1000000000; i++ {
		digs := num2Array(i)
		found := true
		for m := mult; m > 1; m-- {
			digs2 := num2Array(i * m)
			if !mapsEqual(digs, digs2) {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return 0
}

func main() {
	println(findPermutedMult(6))
}
