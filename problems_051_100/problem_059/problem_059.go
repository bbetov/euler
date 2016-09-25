package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func loadText() []int {
	var b []int
	fileName := "p059_cipher.txt"
	f, _ := ioutil.ReadFile(fileName)
	s := string(f)
	s = s[:len(s)-1]
	ss := strings.Split(s, ",")
	for _, st := range ss {
		i, _ := strconv.Atoi(st)
		b = append(b, i)
	}
	return b
}

func loadWords() (rv map[string]bool) {
	rv = make(map[string]bool)
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rv[scanner.Text()] = true
	}
	return rv
}

func analyzePotential(in []byte, words map[string]bool) bool {
	s := string(in)
	ss := strings.FieldsFunc(s, func(r rune) bool {
		switch r {
		case ',', ';', ' ', '(', ')', '.', '!', '\n', '\r':
			return true
		}
		return false
	})
	var counter int
	for _, str := range ss {
		match, _ := regexp.MatchString("^[0-9]+$", str)
		if _, ok := words[str]; ok || match {
			counter++
		}
	}
	valid := float32(counter) / float32(len(ss))
	if valid > 0.80 {
		return true
	}

	return false
}

func main() {
	b := loadText()
	words := loadWords()
	pwdchars := []int{32}
	for i := 97; i < 97+26; i++ {
		pwdchars = append(pwdchars, i)
	}
	//fmt.Printf("%v\n", b)
	for _, x := range pwdchars {
		for _, y := range pwdchars {
			for _, z := range pwdchars {
				pwd := []int{x, y, z}
				potential := true
				decoded := make([]byte, len(b))
				for idx, v := range b {
					pc := pwd[idx%3]
					p := v ^ pc
					if (p < 32 || p > 126) && p != 10 && p != 13 {
						potential = false
						break
					}
					decoded[idx] = byte(p)
				}
				if potential {
					if analyzePotential(decoded, words) {
						s := 0
						for _, v := range decoded {
							s += int(v)
						}
						println(s)
						os.Exit(0)
					}
				}
			}
		}
	}
}
