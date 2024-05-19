package problem1798

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test1789(t *testing.T) {
	testCase1 := []int{1, 3}

	assert.Equal(t, 2, getMaximumConsecutive(testCase1))

	testCase2 := []int{1, 1, 1, 4}

	assert.Equal(t, 8, getMaximumConsecutive(testCase2))

	testCase3 := []int{1, 4, 10, 3, 1}

	assert.Equal(t, 20, getMaximumConsecutive(testCase3))
}

func getMaximumConsecutive(coins []int) int {
	slices.Sort(coins)

	currentSum := 0

	for i := 0; i < len(coins); i++ {
		if currentSum < coins[i]-1 {
			break
		}

		currentSum += coins[i]
	}

	return currentSum + 1
}
