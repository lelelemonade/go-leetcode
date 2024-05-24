package problem714

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test714(t *testing.T) {
	assert.Equal(t, 8, maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}

func maxProfit(prices []int, fee int) int {
	holdDp := make([]int, len(prices))
	notHoldDp := make([]int, len(prices))
	holdDp[0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		holdDp[i] = max(holdDp[i-1], notHoldDp[i-1]-prices[i])
		notHoldDp[i] = max(notHoldDp[i-1], holdDp[i-1]+prices[i]-fee)
	}
	return notHoldDp[len(prices)-1]
}
