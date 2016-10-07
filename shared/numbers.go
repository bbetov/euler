package shared

type ByteArray []byte

func (c ByteArray) Len() int {
	return len(c)
}

func (c ByteArray) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c ByteArray) Less(i, j int) bool {
	return c[i] < c[j]
}

func Num2DigitsMap(p uint64) (rv map[byte]int) {
	rv = make(map[byte]int)
	for p > 0 {
		r := byte(p % 10)
		p /= 10
		rv[byte(r)]++
	}
	return rv
}

func IsDigitsMapsEqual(a, b map[byte]int) bool {
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

func Num2Array(p uint64) (rv ByteArray) {
	for p > 0 {
		r := byte(p % 10)
		p /= 10
		rv = append(rv, byte(r))
	}
	return rv
}

func Array2Num(b ByteArray) (rv uint64) {
	scalar := uint64(1)
	for i := 0; i < len(b); i++ {
		rv += scalar * uint64(b[i])
		scalar *= 10
	}
	return rv
}
