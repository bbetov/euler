package main

// Bruteforced it ... I will worry about efficiency when I get to 67...

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type node struct {
	l *node
	r *node
	v int
}

func (n *node) maxPath() int {
	if n.l == nil {
		return n.v
	}

	l := n.l.maxPath()
	r := n.r.maxPath()
	if l > r {
		return n.v + l
	}
	return n.v + r
}

func loadData(path string) (*node, error) {
	var prev []*node
	var root *node
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstRow := true
	for scanner.Scan() {
		line := scanner.Text()
		toks := strings.Split(line, " ")
		var cur []*node
		for _, t := range toks {
			i, _ := strconv.Atoi(t)
			cur = append(cur, &node{v: i})
		}
		if !firstRow {
			for i, v := range prev {
				v.l = cur[i]
				v.r = cur[i+1]
			}
		} else {
			root = cur[0]
		}
		prev = cur
		firstRow = false
	}
	return root, scanner.Err()
}

func main() {
	rn, _ := loadData("input.txt")
	println(rn.maxPath())
}
