package letter

// FreqMap is simple map of word count.
type FreqMap map[rune]int

// Frequency takes a text and returns it's word count map.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func wordCount(s string, c chan FreqMap) {
	c <- Frequency(s)
}

// ConcurrentFrequency implements word count of list of strings in parallel.
func ConcurrentFrequency(slist []string) FreqMap {
	ch := make(chan FreqMap, len(slist))

	for _, s := range slist {
		go wordCount(s, ch)
	}

	m := FreqMap{}

	for i := 0; i < len(slist); i++ {
		temp := <-ch
		for k, v := range temp {
			m[k] += v
		}
	}

	return m
}
