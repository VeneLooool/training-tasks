package romannumerals

import "strings"

type romanArabic struct {
	roman  string
	arabic int
}
type decodeAndEncode struct {
	numbers []romanArabic
}

func (decAndEnc *decodeAndEncode) buildup() {
	decAndEnc.numbers = []romanArabic{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}
}

func Encode(n int) (result string, flag bool) {
	if n <= 0 {
		return "", false
	}

	var enc decodeAndEncode
	enc.buildup()

	for _, v := range enc.numbers {
		for n >= v.arabic {
			n -= v.arabic
			result += v.roman
		}
	}
	return result, len(result) > 0
}

func Decode(s string) (result int, flag bool) {
	if len(s) == 0 {
		return 0, false
	}

	var dec decodeAndEncode
	dec.buildup()

	for _, v := range dec.numbers {
		for strings.HasPrefix(s, v.roman) {
			result += v.arabic
			s = s[len(v.roman):]
		}
	}
	if len(s) != 0 {
		return 0, false
	}
	return result, result > 0
}
