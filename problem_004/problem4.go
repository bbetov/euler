package main

// A palindromic number reads the same both ways. The largest palindrome made
// from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}

	number := num
	var n int
	for number > 0 {
		n = 10*n + number%10
		number /= 10
	}

	return n == num
}

func main() {
	maxPali := 0
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			n := i * j
			if isPalindrome(n) && n > maxPali {
				maxPali = n
			}
		}
	}
	println(maxPali)
}
