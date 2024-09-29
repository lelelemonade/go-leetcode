package problem1475

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func finalPrices(prices []int) []int {
	maxStack := make([]int, 0)
	result := make([]int, len(prices))

	for i, price := range prices {
		n := len(maxStack)
		for n > 0 && prices[maxStack[n-1]] >= price {
			result[maxStack[n-1]] = prices[maxStack[n-1]] - price
			maxStack = maxStack[:n-1]
		}
		maxStack = append(maxStack, i)
	}
	for _, v := range maxStack {
		result[v] = prices[v]
	}

	return result
}

func Test1475(t *testing.T) {
	assert.Equal(t, []int{4, 2, 4, 2, 3}, finalPrices([]int{8, 4, 6, 2, 3}))
}
