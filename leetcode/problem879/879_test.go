package problem879

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	dp := make([][][]int, len(group)+1)

	for j := 0; j < len(group)+1; j++ {
		dp[j] = make([][]int, n+1)
		for k := 0; k < n+1; k++ {
			dp[j][k] = make([]int, minProfit+1)
		}
	}

	return 1

}
