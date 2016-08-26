package main

import (
	"bufio"
	"os"
	"strconv"
)

func readInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, c := range scanner.Text() {
			if i, err := strconv.ParseInt(string(c), 10, 8); err == nil {
				nums = append(nums, int(i))
			}
		}
	}
	return nums, scanner.Err()
}

func adjacentProduct(numDigits int) (res int) {
	nums, _ := readInput("input.txt")
	numDigits--
	for h := numDigits; h < len(nums); h++ {
		cs := 1
		for _, i := range nums[h-numDigits : h+1] {
			cs *= i
		}
		if res < cs {
			res = cs
		}
	}
	return res
}

func main() {
	println(adjacentProduct(13))
}
