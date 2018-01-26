package etl

import "strings"

// Transform trnasforms map of int to string list to map of string to int.
func Transform(input map[int][]string) map[string]int {
	result := make(map[string]int)
	for key, val := range input {
		for _, cur := range val {
			result[strings.ToLower(cur)] = key
		}
	}
	return result
}
