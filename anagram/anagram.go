package anagram

import (
	"strings"
)

func runeToZero(input int) (output int) {
	const FirstRuneOfAlphabet = 97
	return input - FirstRuneOfAlphabet
}

func calculateAmountOfRune(input string) (output [26]int) {
	input = strings.ToLower(input)
	for r := range input {
		if input[r] >= 'a' && input[r] <= 'z' {
			output[runeToZero(int(input[r]))]++
		}
	}
	return output
}

func deleteSymbol(in string) (output string) {
	if len(in) > 0 {
		if int(in[len(in)-1]) == 13 {
			in = in[0 : len(in)-1]
		}
	}
	return in
}

func FindAnagrams(dictionary []string, word string) (result []string) {

	runeOfWord := calculateAmountOfRune(word)
	word = strings.ToLower(word)

	for i := range dictionary {
		s := strings.ToLower(dictionary[i])
		s = deleteSymbol(s)

		if s == word || s == "" {
			continue
		}

		runeOfCurrent := calculateAmountOfRune(dictionary[i])

		if runeOfWord == runeOfCurrent {
			dictionary[i] = deleteSymbol(dictionary[i])
			result = append(result, dictionary[i])
		}
	}
	return result
}
