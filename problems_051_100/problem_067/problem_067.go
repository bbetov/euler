package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func loadData(path string) ([][]int, error) {
	var all [][]int
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		toks := strings.Split(line, " ")
		cur := []int{}
		for _, t := range toks {
			i, _ := strconv.Atoi(t)
			cur = append(cur, i)
		}
		all = append(all, cur)
	}
	return all, scanner.Err()
}

func main() {
	ii, _ := loadData("p067_triangle.txt")
	for i := len(ii) - 1; i > 0; i-- {
		c := ii[i-1]
		n := ii[i]
		for k := 0; k < len(c); k++ {
			if n[k] > n[k+1] {
				c[k] += n[k]
			} else {
				c[k] += n[k+1]
			}
		}
	}
	println(ii[0][0])

}
