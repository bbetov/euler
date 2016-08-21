package main

import (
	"bufio"
	"os"
	"strings"
	"unicode/utf8"
)

func wordScore(w string) (score int) {
	tmp := make([]byte, 4)
	utf8.EncodeRune(tmp, 'A')
	aval := tmp[0]
	for _, c := range w {
		if c >= 'A' && c <= 'Z' {
			utf8.EncodeRune(tmp, c)
			score += int((tmp[0] - aval) + byte(1))
		}
	}
	return score
}

func main() {
	// Read file work by word
	f, _ := os.Open("p042_words.txt")
	defer f.Close()
	trNum := make(map[int]bool)
	reader := bufio.NewReader(f)
	// TODO: should do a smarter way to fill this on demand, but for now will do
	for i := 0; i < 2600; i++ { // 100 letter word
		ii := i * (i + 1)
		ii /= 2
		if ii > 2600 {
			break
		}
		trNum[ii] = true
	}

	// Scan all words from the file.
	var cntr int
	for {
		word, err := reader.ReadString(',')
		if err == nil {
			word = strings.ToUpper(strings.Trim(word, "\","))
			score := wordScore(word)
			if _, ok := trNum[score]; ok {
				cntr++
				//fmt.Printf("%v ==> %v\n", word, score)
			}
		} else {
			break
		}
	}
	println(cntr)
}
