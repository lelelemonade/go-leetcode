package problem122

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test122(t *testing.T) {
	assert.Equal(t, 7, maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	maxOfProfit :=0
	minLeft :=prices[0]
	for i := 1; i < len(prices); i++ {
		if minLeft <= prices[i] {
			maxOfProfit+=prices[i]-minLeft
		}
		minLeft=prices[i]
	}
	return maxOfProfit
}

// func maxProfit(prices []int) int {
// 	maxHoldingDp := make([]int, len(prices))
// 	maxHoldingDp[0] = -prices[0]
// 	maxNotHoldingDp := make([]int, len(prices))
// 	maxNotHoldingDp[0] = 0

// 	for i := 1; i < len(prices); i++ {
// 		maxHoldingDp[i] = max(maxHoldingDp[i-1], maxNotHoldingDp[i-1]-prices[i])
// 		maxNotHoldingDp[i] = max(maxNotHoldingDp[i-1], maxHoldingDp[i-1]+prices[i])

// 	}

// 	return maxNotHoldingDp[len(prices)-1]
// }
