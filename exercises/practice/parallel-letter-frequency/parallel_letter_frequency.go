package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(words []string) FreqMap {
	c := make(chan FreqMap)
	for _, word := range words {
		go func(word string) {
			c <- Frequency(word)
		}(word)
	}

	result := FreqMap{}
	for range words {
		for letter, count := range <-c {
			result[letter] += count
		}
	}

	return result
	//panic("Implement the ConcurrentFrequency function")
}
