package secretmessage

import "sort"

// Decode func
func Decode(encoded string) string {
	counter := make(map[rune]int, 27)
	for _, v := range encoded {
		counter[v]++
	}
	runeCounter := make([]rune, 0, 26)

	for key, value := range counter {
		if key != '_' && value >= counter['_'] {
			runeCounter = append(runeCounter, key)
		}
	}

	sort.Slice(runeCounter, func(i, j int) bool {
		return counter[runeCounter[i]] > counter[runeCounter[j]]
	})

	return string(runeCounter)
}
