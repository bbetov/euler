package main

func fibEven(br int64) int64 {
	x, y := int64(1), int64(1)
	sum := int64(0)
	for y < br {
		if y%2 == 0 {
			sum += y
		}
		x, y = y, x+y
	}
	return sum
}

func main() {
	print(fibEven(4000000))
}
