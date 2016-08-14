package main

type rangeInfo struct {
	scalar         int
	digFrom, digTo int
	from, to       int
}

func getDigit(n int, r []rangeInfo) int {
	// Find range
	var ri rangeInfo
	for i := 0; i < len(r); i++ {
		if n >= r[i].digFrom && n <= r[i].digTo {
			ri = r[i]
			break
		}
	}
	offsetFromBegRange := n - ri.digFrom
	numberOffset := offsetFromBegRange / ri.scalar
	num := numberOffset + ri.from
	digitWithinNum := offsetFromBegRange % ri.scalar

	digits := []int{}
	for num > 0 {
		rem := num % 10
		num /= 10
		digits = append(digits, rem)
	}

	return digits[len(digits)-digitWithinNum-1]
}

func main() {
	// Prepare data
	var data []rangeInfo
	var r rangeInfo
	r.from = 1
	r.to = 9
	r.scalar = 1
	r.digFrom = 1
	r.digTo = 9

	for r.digFrom < 100000000 {
		data = append(data, r)
		r.from = r.from * 10
		r.to = r.from*10 - 1
		r.scalar++
		r.digFrom = r.digTo + 1
		r.digTo = r.digFrom + (r.to+1-r.from)*r.scalar - 1
		//fmt.Println(data)
	}

	prod := 1
	digit := 1
	for p := 0; p <= 6; p++ {
		prod *= getDigit(digit, data)
		digit *= 10
	}
	println(prod)
}
