package armstrong

import (
	"math"
	"strconv"
)

// IsNumber checks if given number is armstrong number or not.
func IsNumber(n int) bool {
	nstr := strconv.Itoa(n)
	exp := len(nstr)
	var curNum int
	for i := 0; i < len(nstr); i++ {
		c, err := strconv.Atoi(string(nstr[i]))
		if err != nil {
			panic("string to int conversion failed")
		}

		curNum += int(math.Pow(float64(c), float64(exp)))
	}

	return curNum == n
}
