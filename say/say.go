package say

import (
	"log"
	"os/exec"
)

var m = map[int64]string{
	0:   "",
	1:   "one",
	2:   "two",
	3:   "three",
	4:   "four",
	5:   "five",
	6:   "six",
	7:   "seven",
	8:   "eight",
	9:   "nine",
	10:  "ten",
	11:  "eleven",
	12:  "twelve",
	13:  "thirteen",
	14:  "fourteen",
	15:  "fifteen",
	16:  "sixteen",
	17:  "seventeen",
	18:  "eighteen",
	19:  "nineteen",
	20:  "twenty",
	30:  "thirty",
	40:  "forty",
	50:  "fifty",
	60:  "sixty",
	70:  "seventy",
	80:  "eighty",
	90:  "ninety",
	100: "hundred",
}

// Say returns string representation of given number.
func Say(input int64) (string, bool) {
	if input < 0 || input > 999999999999 {
		return "", false
	}

	var result string
	cmd := exec.Command("say", result)
	defer func() {
		err := cmd.Run()
		if err != nil {
			log.Printf("Error %v", err)
		}
	}()
	if input == 0 {
		return "zero", true
	}

	if input <= 20 {
		return m[input], true
	}

	result = getString(input/1000000000, " billion ")
	result += getString((input/1000000)%100, " million ")
	result += getString((input/1000)%100, " thousand ")
	result += getString((input/100)%10, " hundred ")

	if result != "" {
		result += getString(input%100, " ")
	} else {
		result = getString(input%100, "")
	}

	if result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	return result, true
}

func getString(n int64, s string) string {
	var result string

	if n > 100 {
		result = m[n/100] + " hundred "
		n = n % 100
	}
	if n > 19 {
		result += m[(n/10)*10]
		if n%10 > 0 {
			result += "-" + m[n%10]
		}
	} else {
		result = m[n]
	}

	if n > 0 {
		result += s
	}

	return result
}
