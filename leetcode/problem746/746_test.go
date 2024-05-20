package problem746

func minCostClimbingStairs(cost []int) int {
	stairCount := len(cost)
	minCost := make([]int, stairCount)
	minCost[0] = 0
	minCost[1] = 0
	for i := 2; i < len(cost); i++ {
		minCost[i] = min(minCost[i-2]+cost[i-2], minCost[i-1]+cost[i-1])
	}
	return min(minCost[stairCount-1]+cost[stairCount-1], minCost[stairCount-2]+cost[stairCount-2])
}
