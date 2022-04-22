package coins

var P = make(map[int]int, 100)

func Piles(n int) int {
	if n < 0 {
		return -1
	}
	if value, ok := P[n]; ok {
		return value
	}

	P[0] = 1
	P[1] = 1

	for i := 2; i <= n; i++ {
		var (
			orgDelta, blueDelta, generalDelta int
			blueOrOrange, sum                 int
		)
		orgDelta = 1
		blueDelta = 1
		blueOrOrange = 1

		for {
			sum = sum + blueOrOrange*P[i-generalDelta-orgDelta]
			generalDelta += orgDelta
			sum = sum + blueOrOrange*P[i-generalDelta-blueDelta]
			generalDelta += blueDelta
			blueOrOrange *= -1
			blueDelta++
			orgDelta += 2
			if i-orgDelta < 0 || i-blueDelta < 0 {
				break
			}
		}
		P[i] = sum
	}

	return P[n]
}
