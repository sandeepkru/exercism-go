package accumulate

type converter func(string) string

// Accumulate takes slice of strings and converts as per converter and returns new slice.
func Accumulate(input []string, cnfunc converter) []string {
	var output []string
	for _, word := range input {
		output = append(output, cnfunc(word))
	}
	return output
}
