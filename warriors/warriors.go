package warriors

import "strings"

func Count(image string) int {
	if image == "" {
		return 0
	}

	counter := 1
	newImage := parseImage(image)

	for i, line := range newImage {
		for j, column := range line {
			if column > 1 {
				paintRegion(column, i, j, newImage)
			}
			if column == 1 {
				counter++
				paintRegion(counter, i, j, newImage)
			}
		}
	}
	return counter - 1
}

func parseImage(image string) (parsedResult [][]int) {

	rows := strings.Split(image, "\n")
	parsedResult = make([][]int, len(rows)+1)
	for curRow, line := range rows {
		parsedResult[curRow] = make([]int, len(line))
		for column, thisRune := range line {
			parsedResult[curRow][column] = int(thisRune - '0')
		}
	}
	return parsedResult

}

func paintRegion(color, pointI, pointJ int, image [][]int) {
	for i := 0; i <= 1; i++ {
		for j := -i; j <= 1; j++ {
			toPaintI := pointI + i
			toPaintJ := pointJ + j

			isInBorder := false
			if toPaintJ >= 0 && toPaintI < len(image) && toPaintJ < len(image[toPaintI]) {
				isInBorder = true
			}

			if isInBorder && image[toPaintI][toPaintJ] == 1 {
				image[toPaintI][toPaintJ] = color
			}
		}
	}
}
