package main

import (
	"errors"
	"fmt"
)

func nextLexPerm(a []int) ([]int, error) {
	k := -1
	var l int
	for i := len(a) - 2; i >= 0; i-- {
		if a[i] < a[i+1] {
			k = i
			break
		}
	}
	if k < 0 {
		return nil, errors.New("No lex perm")
	}

	for i := len(a) - 1; i > k; i-- {
		if a[k] < a[i] {
			l = i
			break
		}
	}
	a[k], a[l] = a[l], a[k]
	q := len(a) - 1
	for i := k + 1; i < q; i++ {
		a[i], a[q] = a[q], a[i]
		q--
	}

	return a, nil
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//a := []int{1, 2, 3, 4}
	var err error
	for i := 1; i < 1000000; i++ {
		a, err = nextLexPerm(a)
		if err != nil {
			fmt.Printf("Failed getting perm %d: %s", i, err.Error())
		}
	}
	fmt.Printf("%v", a)
}
