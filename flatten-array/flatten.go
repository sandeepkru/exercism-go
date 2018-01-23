package flatten

// Flatten flattens given nested list.
func Flatten(input interface{}) []interface{} {
	result := []interface{}{}
	for _, cur := range input.([]interface{}) {
		switch cur.(type) {
		case []interface{}:
			sub := Flatten(cur)
			result = append(result, sub...)
		case int:
			result = append(result, cur)
		}
	}
	return result
}
