package problem643

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findMaxAverage(nums []int, k int) float64 {
	var avg float64 = 0
	for i := 0; i < k; i++ {
		avg += float64(nums[i])
	}
	avg /= float64(k)
	result := avg
	for i := k; i < len(nums); i++ {
		avg = avg + float64(nums[i]-nums[i-k])/float64(k)
		result = max(avg, result)
	}
	return result
}

func Test643(t *testing.T) {
	assert.Equal(t, float64(4), findMaxAverage([]int{0, 4, 0, 3, 2}, 1))
}
