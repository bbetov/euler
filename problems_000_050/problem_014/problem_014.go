package main

import "fmt"

func collatzLen(n int64) (l int64) {
	l = 1
	for n != 1 {
		if n&1 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		l++
	}
	return l
}

func getMaxCollatzSeq() (len, start int64) {
	for i := int64(1); i < int64(1000000); i++ {
		l := collatzLen(i)
		if len < l {
			len = l
			start = i
		}
	}
	return len, start
}

func main() {
	maxLen, maxSeed := getMaxCollatzSeq()
	fmt.Printf("Max: %d ==> %d\n", maxSeed, maxLen)
}
