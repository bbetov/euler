package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func loadData(path string) ([][]int, error) {
	var rv [][]int
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		rv = append(rv, []int{})
		for _, c := range tokens {
			if i, err := strconv.ParseInt(string(c), 10, 16); err == nil {
				rv[row] = append(rv[row], int(i))
			}
		}
		row++
	}

	return rv, scanner.Err()
}

func maxProduct(i, j int, arr [][]int) int {
	mp := 0
	a, b, c, d := 1, 1, 1, 1
	for o := 0; o < 4; o++ {
		if i < len(arr)-4 {
			b *= arr[i+o][j]
		}
		if j < len(arr[i])-4 {
			a *= arr[i][j+o]
		}
		if j < len(arr[i])-4 && i < len(arr)-4 {
			c *= arr[i+o][j+o]
		}
		if j < len(arr[i])-4 && i >= 4 {
			d *= arr[i-o][j+o]
		}
	}
	mp = a
	if mp < b {
		mp = b
	}
	if mp < c {
		mp = c
	}
	if mp < d {
		mp = d
	}
	return mp
}

func getMaxProduct() (m int) {
	arr, _ := loadData("input.txt")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			mp := maxProduct(i, j, arr)
			if m < mp {
				m = mp
			}
		}
	}
	return m
}

func main() {
	println(getMaxProduct())
}
