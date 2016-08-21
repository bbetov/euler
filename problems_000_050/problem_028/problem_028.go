package main

import "fmt"

func main() {
	sum := 1
	current := 1
	for a := 3; a <= 1001; a += 2 {
		for i := 0; i < 4; i++ {
			current += a - 1
			sum += current
		}

	}
	fmt.Printf("The sum is: %v", sum)
}
