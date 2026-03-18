package problem3070

func countSubmatrices(grid [][]int, k int) int {
	prefixSum := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		prefixSum[i] = make([]int, len(grid[i]))
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			prefixSum[i][j] = grid[i][j]
			if i > 0 {
				prefixSum[i][j] += prefixSum[i-1][j]
			}
			if j > 0 {
				prefixSum[i][j] += prefixSum[i][j-1]
			}
			if i > 0 && j > 0 {
				prefixSum[i][j] -= prefixSum[i-1][j-1]
			}
			if prefixSum[i][j] <= k {
				result++
			}
		}
	}

	return result
}
