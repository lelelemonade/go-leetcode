package problem309

func maxProfit(prices []int) int {
	holdDp := make([]int, len(prices))
	keepSoldDp := make([]int, len(prices))
	soldTodayDp := make([]int, len(prices))
	coolDownTodayDp := make([]int, len(prices))
	holdDp[0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		holdDp[i] = max(holdDp[i-1], max(coolDownTodayDp[i-1], keepSoldDp[i-1])-prices[i])
		keepSoldDp[i] = max(keepSoldDp[i-1], coolDownTodayDp[i-1])
		soldTodayDp[i] = holdDp[i-1] + prices[i]
		coolDownTodayDp[i] = soldTodayDp[i-1]
	}

	return max(keepSoldDp[len(prices)-1], soldTodayDp[len(prices)-1], coolDownTodayDp[len(prices)-1])
}
