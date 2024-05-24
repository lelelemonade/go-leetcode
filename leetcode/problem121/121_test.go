package problem121

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test121(t *testing.T) {
	assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	leftMin := prices[0]
	maxOfProfit := 0

	for i := 1; i < len(prices); i++ {
		maxOfProfit = max(maxOfProfit, prices[i]-leftMin)
		leftMin = min(leftMin, prices[i])
	}

	return maxOfProfit
}
