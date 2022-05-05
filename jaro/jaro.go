package jaro

import (
	"strings"
)

func Distance(word1 string, word2 string) (ans float64) {
	word1, word2 = strings.ToLower(word1), strings.ToLower(word2)
	firstWordLen, secondWordLen := len(word1), len(word2)

	switch {
	case firstWordLen == 0 && secondWordLen == 0:
		return 1
	case firstWordLen == 0 || secondWordLen == 0:
		return 0
	case word1 == word2:
		return 1
	}

	maxDistBetweenLetter := (Max(firstWordLen, secondWordLen) / 2) - 1

	amountMatchSim, runeWord1, runeWord2 := countMatchAndNoteThem(word1, word2, maxDistBetweenLetter)

	if amountMatchSim == 0 {
		return 0
	}

	transposition := countTransposition(word1, word2, runeWord1, runeWord2, amountMatchSim) / 2

	return float64((amountMatchSim-transposition)*firstWordLen*secondWordLen+
		amountMatchSim*amountMatchSim*firstWordLen+
		amountMatchSim*amountMatchSim*secondWordLen) / float64(firstWordLen*secondWordLen*amountMatchSim*3)
}

func countMatchAndNoteThem(word1, word2 string, maxDistBetweenLetter int) (amountMatchSim int, runeWord1, runeWord2 []rune) {
	runeWord1, runeWord2 = []rune(word1), []rune(word2)
	l2 := len(runeWord2)
	for i, LTFirstW := range word1 {
		leftBorder, rightBorder := i-maxDistBetweenLetter, i+maxDistBetweenLetter+1

		if leftBorder >= l2 {
			return amountMatchSim, runeWord1, runeWord2
		}
		if leftBorder < 0 {
			leftBorder = 0
		}
		if rightBorder > l2 {
			rightBorder = l2
		}

		for j := leftBorder; j < rightBorder; j++ {
			if LTFirstW == runeWord2[j] {
				amountMatchSim++
				runeWord1[i], runeWord2[j] = 0, 0
				break
			}
		}
	}
	return amountMatchSim, runeWord1, runeWord2
}

func countTransposition(word1, word2 string, runeWord1, runeWord2 []rune, amountMatchSim int) (transposition int) {
	suppRuneWord1, suppRuneWord2 := []rune(word1), []rune(word2)
	var j, curr int
	for i := 0; curr < amountMatchSim; i++ {
		if runeWord1[i] != 0 {
			continue
		}
		//Я выделил в отдельную функцию, так как очень сложно понять что тут происходит.
		//Вопрос: правильно ли так выделять(ведь мы передаем массив(это время и память))?
		j = searchForOccurrenceInSecondWord(j, runeWord2)

		if suppRuneWord1[i] != suppRuneWord2[j] {
			transposition++
		}
		curr++
		j++
	}
	return transposition
}

func searchForOccurrenceInSecondWord(curJ int, runeWord2 []rune) (outJ int) {
	for runeWord2[curJ] != 0 {
		curJ++
	}
	return curJ
}

func Max(a int, b int) (max int) {
	if a > b {
		return a
	}
	return b
}
