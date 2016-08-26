package main

func sum(upper int) int {
	var retval int
	for i := 3; i < upper; i++ {
		if i%3 == 0 || i%5 == 0 {
			retval += i
		}
	}
	return retval
}

func main() {
	const upper int = 1000
	println(sum(upper))
}
