package problem3212

func numberOfSubmatrices(grid [][]byte) int {
	prefixSum := make([][]int, len(grid))
	containX := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		prefixSum[i] = make([]int, len(grid[i]))
		containX[i] = make([]bool, len(grid[i]))
	}

	result := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch grid[i][j] {
			case 'X':
				containX[i][j] = true
				prefixSum[i][j]++
			case 'Y':
				prefixSum[i][j]--
			}

			if i > 0 {
				containX[i][j] = containX[i][j] || containX[i-1][j]
				prefixSum[i][j] += prefixSum[i-1][j]
			}
			if j > 0 {
				containX[i][j] = containX[i][j] || containX[i][j-1]
				prefixSum[i][j] += prefixSum[i][j-1]
			}
			if i > 0 && j > 0 {
				prefixSum[i][j] -= prefixSum[i-1][j-1]
			}
			if containX[i][j] && prefixSum[i][j] == 0 {
				result++
			}
		}
	}

	return result
}
