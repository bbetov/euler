package main

import "fmt"

var fs = make([]uint64, 10)

func findUpperBound() uint64 {
	maxNum := uint64(9)
	f := fs[9]
	for maxNum < f {
		f = f + fs[9]
		maxNum = maxNum*10 + 9
	}
	return f + 1
}

func main() {
	prod := uint64(1)
	fs[0] = uint64(1)
	for i := uint64(1); i < 10; i++ {
		prod *= i
		fs[i] = prod
	}
	ub := findUpperBound()
	answer := uint64(0)
	for i := uint64(3); i < ub; i++ {
		num := i
		sum := uint64(0)
		for num > 0 {
			d := num % 10
			sum += fs[d]
			if sum > i {
				break
			}
			num /= 10
		}
		if sum == i {
			answer += i
		}
	}

	fmt.Println(answer)
}
