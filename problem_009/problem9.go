package main

import "fmt"

type triple struct {
	a, b, c int
}

func pitTripleGen() (ch chan triple) {
	ch = make(chan Triple)

	go func() {
		for a := 2; a < 1000; a++ {
			for b := a; b < 1000; b++ {
				for c := b; c < 1000; c++ {
					if a+b+c == 1000 && a*a+b*b == c*c {
						ch <- triple{a, b, c}
					}
				}
			}
		}
		close(ch)
	}()

	return ch
}

func main() {
	for t := range pitTripleGen() {
		fmt.Printf("%v, Product: %d\n", t, t.a*t.b*t.c)
	}
}
