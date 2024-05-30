package problem2190

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test2190(t *testing.T) {
	assert.Equal(t, 2, mostFrequent([]int{2, 2, 2, 2, 3}, 2))

}

func mostFrequent(nums []int, key int) int {
	count := make(map[int]int)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			count[nums[i+1]]++
		}
	}

	result := 0
	maxCount := 0
	for k, v := range count {
		if v > maxCount {
			result = k
			maxCount = v
		}
	}
	return result
}
