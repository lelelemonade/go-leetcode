package problem120

func minimumTotal(triangle [][]int) int {

	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle))
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	min := 10 ^ 4
	for i := 0; i < len(triangle); i++ {
		if dp[len(triangle)-1][i] < min {
			min = dp[len(triangle)-1][i]
		}
	}
	return min

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
