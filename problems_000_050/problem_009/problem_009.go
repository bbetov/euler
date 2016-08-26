package main

type triple struct {
	a, b, c int
}

func pitTripleGen() (t *triple) {

	for a := 2; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			for c := b; c < 1000; c++ {
				if a+b+c == 1000 && a*a+b*b == c*c {
					return &triple{a, b, c}
				}
			}
		}
	}
	return nil
}

func main() {
	t := pitTripleGen()
	println(t.a * t.b * t.c)
}
