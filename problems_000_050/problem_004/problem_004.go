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

func getMaxPalindrome(lim int) (mp int) {
	for i := 1; i < lim; i++ {
		for j := 1; j < lim; j++ {
			n := i * j
			if isPalindrome(n) && n > mp {
				mp = n
			}
		}
	}
	return mp
}

func main() {
	println(getMaxPalindrome(1000))
}
