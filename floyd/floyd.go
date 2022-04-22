package floyd

// Triangle makes a Floyd's triangle matrix with rows count.
func Triangle(rows int) [][]int {
	if rows <= 0 {
		return [][]int{}
	}

	var res [][]int
	counter := 1

	for i := 1; i <= rows; i++ {
		var curLevel []int
		for j := 0; j < i; j++ {
			curLevel = append(curLevel, counter)
			counter++
		}
		res = append(res, curLevel)
	}

	return res
}
