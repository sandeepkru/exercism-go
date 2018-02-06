package diamond

import (
	"errors"
)

// Gen returns diamond string of b.
func Gen(b byte) (string, error) {
	var result string
	if b < 'A' || b > 'Z' {
		return result, errors.New("Invalid Byte")
	}
	diff := int((b - 'A' + 1))
	lineCount := (diff * 2) - 1 // this serves as width as well
	lines := make([]string, diff)
	mid := int(lineCount / 2)
	for i := 0; i < diff; i++ {
		cur := 'A' + i
		curLine := make([]byte, lineCount)
		for j := 0; j < lineCount; j++ {
			curLine[j] = ' '
		}
		// Handle first and last indices seperately.
		if i == 0 {
			curLine[mid] = 'A'
		} else if i == diff-1 {
			curLine[0] = byte(cur)
			curLine[lineCount-1] = byte(cur)
		} else {
			curLine[mid-i] = byte(cur)
			curLine[mid+i] = byte(cur)
		}
		lines[i] = string(curLine[:lineCount])
	}

	// create result string by adding all lines
	for i := 0; i < diff; i++ {
		result += lines[i] + "\n"
	}
	for i := diff - 2; i >= 0; i-- {
		result += lines[i] + "\n"
	}
	return result, nil
}
