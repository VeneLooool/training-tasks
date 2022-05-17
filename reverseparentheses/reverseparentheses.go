package reverseparentheses

import (
	"errors"
)

func Reverse(s string) string {
	if isCorrectString, _ := checkStringsFormat(s); !isCorrectString {
		return ""
	}
	return string(reverseString([]byte(s)))
}
func reverseString(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			s = append(s[:i], reverseString(s[i+1:])...)
		}
		if s[i] == ')' {
			revers(s[:i])
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
func revers(s []byte) {
	j := 1
	for i := 0; i < len(s)/2; i++ {
		x := s[i]
		s[i] = s[len(s)-j]
		s[len(s)-j] = x
		j++
	}
}

func checkStringsFormat(input string) (output bool, err error) {
	var amountBrackets int
	for _, value := range input {
		if value == '(' {
			amountBrackets++
		}
		if value == ')' {
			amountBrackets--
		}
	}
	if amountBrackets != 0 {
		return false, errors.New("incorrect input")
	}
	return true, nil
}
