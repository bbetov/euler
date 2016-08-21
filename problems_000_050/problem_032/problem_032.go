package main

func fillBits(n int, fl uint16) (uint16, bool) {
	for n > 0 {
		r := uint(n % 10)
		bit := uint16(1 << r)
		if r == 0 || fl&bit > 0 { // already found or zero
			return fl, false
		}
		fl |= bit
		n /= 10
	}
	return fl, true
}

func isPandigital(m, m2, prod int) bool {
	if fl, ok := fillBits(m, 0); ok {
		if fl, ok = fillBits(m2, fl); ok {
			prod := m * m2
			if fl, ok = fillBits(prod, fl); ok {
				if fl == 0x03FE { // 1111111110
					return true
				}
			}
		}
	}

	return false
}

func sumProd() (s int) {
	unique := make(map[int]bool)
	for m := 1; m < 9999; m++ {
		for m2 := m + 1; m2 < 9999; m2++ {

			prod := m * m2
			if prod > 10000 {
				break
			}

			if _, ok := unique[prod]; ok {
				continue
			}

			if ok := isPandigital(m, m2, prod); ok {
				unique[prod] = true
				s += prod
				println(prod)
			}
		}
	}
	return s
}

func main() {
	println(sumProd())
}
