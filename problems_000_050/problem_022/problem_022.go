package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("p022_names.txt")
	if err != nil {
		return
	}
	defer file.Close()

	var names []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n := strings.Split(line, ",")
		for _, nn := range n {
			names = append(names, strings.Replace(nn, "\"", "", -1))
		}
	}

	sort.Strings(names)

	var total uint64
	for i, n := range names {
		ws := uint64(0)
		for _, c := range n {
			ws += uint64(c) - 'A' + 1
		}
		total += ws * uint64(i+1)
	}
	println(total)
}
