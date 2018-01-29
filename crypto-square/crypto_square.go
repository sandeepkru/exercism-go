package cryptosquare

import (
	"bytes"
	"log"
	"regexp"
	"strings"
)

// Encode encodes given input string to based on crypto square behavior.
func Encode(s string) string {
	s = strings.ToLower(s)
	reg, err := regexp.Compile("[^a-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	s = reg.ReplaceAllString(s, "")
	if len(s) <= 1 {
		return s
	}

	i := 1
	j := i + 1
	var c, r int
	for {
		if i*i >= len(s) {
			c = i
			r = i
			break
		} else if i*j >= len(s) {
			c = j
			r = i
			break
		}

		i++
		j++
	}
	slist := make([][]rune, r)
	for i := 0; i < r; i++ {
		slist[i] = make([]rune, c)
	}
	curRow := 0
	curCol := 0
	for i := 0; i < len(s); i++ {
		slist[curRow][curCol] = rune(s[i])
		curCol++
		if curCol >= c {
			curCol = 0
			curRow++
		}
	}
	var result []string
	for i := 0; i < c; i++ {
		column := make([]rune, r)
		for j, row := range slist {
			if strings.Contains(s, string(row[i])) {
				column[j] = row[i]
			}
		}
		b := []byte(string(column))
		result = append(result, string(bytes.Trim(b, "\x00")))
	}

	var final string
	final = strings.Join(result, " ")
	return final
}
