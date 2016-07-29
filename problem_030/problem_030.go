package main

import "fmt"

// 9 to the power of 5 is 59049
// Need to find the number of digits where the number is greater than the numberofdigits * 59049
// Well, 6 * 59049 = 354294 and 7*59049 = 413343, so I think looping up to
// 354294 is good enough, even though there is probably a better way...

func main() {
	pw := make([]int, 10)
	for i := 0; i < 10; i++ {
		pw[i] = i * i * i * i * i
	}
	fmt.Printf("%v\n", pw)

	totalSum := uint64(0)
	for i := 10; i < 354294; i++ {
		tmp := i
		rs := 0

		for tmp > 0 {
			r := tmp % 10
			rs += pw[r]
			if rs > i {
				rs = 0
				break
			}
			tmp /= 10
		}
		if rs == i {
			totalSum += uint64(i)
			//fmt.Printf("Found a macting number: %v\n", i)
		}
	}
	fmt.Printf("TotalSum: %v\n", totalSum)
}
