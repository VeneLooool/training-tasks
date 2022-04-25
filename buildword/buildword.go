package buildword

import "math"

var AllMatches map[int][]int //index length

func BuildWord(word string, fragments []string) int {
	if cap(fragments) == 0 || len(word) == 0 {
		return 0
	}

	var min int
	AllMatches = make(map[int][]int)

	findMatches(word, fragments)
	min = findMinRec(word, 0, 0)

	if min == math.MaxInt {
		return 0
	}
	return min
}

func findMinRec(word string, curIndex int, amountFrag int) int {
	if curIndex == len(word) {
		return amountFrag
	}
	var min int = math.MaxInt

	for i := range AllMatches[curIndex] {
		if curIndex+AllMatches[curIndex][i] <= len(word) {
			min = getMin(min, findMinRec(word, curIndex+AllMatches[curIndex][i], amountFrag+1))
		}
	}

	return min
}

func findMatches(word string, fragments []string) {
	lenWord := len(word)
	for _, frag := range fragments {
		lenCurFrag := len(frag)
		for i := 0; i+lenCurFrag <= lenWord; i++ {
			fragWord := word[i : i+lenCurFrag]
			if fragWord == frag {
				AllMatches[i] = append(AllMatches[i], lenCurFrag)
			}
		}
	}
}

func getMin(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
