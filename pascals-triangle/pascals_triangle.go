package pascal

// Triangle implements pascal triangle generation.
func Triangle(n int) [][]int {
	var output [][]int
	firstRow := []int{1}
	secondRow := []int{1, 1}
	if n == 1 {
		output = append(output, firstRow)
		return output
	}

	if n == 2 {
		output = append(output, firstRow)
		output = append(output, secondRow)
		return output
	}

	output = append(output, firstRow)
	output = append(output, secondRow)
	n -= 2
	for n > 0 {
		prev := len(output) - 1
		prevRow := output[prev]
		curRow := make([]int, len(prevRow)+1)
		for i := 0; i < len(curRow); i++ {
			if i == 0 || i == len(curRow)-1 {
				curRow[i] = 1
			} else {
				curRow[i] = prevRow[i-1] + prevRow[i]
			}
		}
		output = append(output, curRow)
		n--
	}
	return output
}
