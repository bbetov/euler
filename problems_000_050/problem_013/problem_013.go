package main

// Could have used big/Int, but I thought that the point is to implement the actual numeric addition.

import (
	"bufio"
	"os"

	"github.com/bbetov/euler/shared"
)

func loadData(path string) ([]string, error) {
	var rv []string
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rv = append(rv, line)
	}

	return rv, scanner.Err()
}

func getLast10LargeSum() string {
	lines, _ := loadData("input.txt")

	i1 := shared.Integer{}
	for _, val := range lines {
		i := shared.Integer{}
		i.Set(val)
		i1.Add(&i)
	}
	return i1.String()[:10]
}

func main() {
	println(getLast10LargeSum())
}
