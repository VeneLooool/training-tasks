package chess

import (
	"errors"
	"strconv"
)

type Knight struct {
	X int
	Y int
}

type MovesKnight struct {
	OfsX int
	OfsY int
}

func CanKnightAttack(white, black string) (bool, error) {
	if !isStrValid(white) || !isStrValid(black) || white == black {
		return false, errors.New("not valid strings")
	}
	wKnight := getPosKnight(white)
	bKnight := getPosKnight(black)

	allMov := []MovesKnight{
		{+1, +2}, {+1, -2},
		{-1, +2}, {-1, -2},
		{+2, +1}, {+2, -1},
		{-2, +1}, {-2, -1},
	}

	for i := range allMov {
		if wKnight.X+allMov[i].OfsX == bKnight.X && wKnight.Y+allMov[i].OfsY == bKnight.Y {
			return true, nil
		}
	}

	return false, nil
}

func isStrValid(posKnight string) bool {
	if len(posKnight) != 2 {
		return false
	}
	if posKnight[0] < 'a' || posKnight[0] > 'h' || posKnight[1] < '1' || posKnight[1] > '8' {
		return false
	}
	return true
}

func getPosKnight(posKnight string) Knight {
	const firstLetterInAlphabet int = 97
	y, _ := strconv.Atoi(string(posKnight[1]))
	return Knight{
		int(posKnight[0]) - firstLetterInAlphabet,
		y,
	}
}
