package snowflakes

func OverlaidTriangles(n, m int) int {
	if m&1 == 0 || n == 0 {
		return 0
	}

	result, supporting := make([]int, n+1), make([]int, n+1)
	generator := [2][4]int{
		{6, 6, 0},
		{-1, 1, 1},
	}

	result[0] = 1
	for k := 0; k < n-1; k++ {
		for j := 0; j <= k; j++ {
			off := j & 0x7ffffffffffffffe
			j1 := j & 1
			for i := 0; i < 3; i++ {
				supporting[off+i] += result[j] * generator[j1][i]
			}
		}
		result, supporting = supporting, result
		for i := range supporting {
			supporting[i] = 0
		}
	}
	return result[m-1]
}
