package problem164

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maximumGap(nums []int) int {
	maxNum := slices.Max(nums)

	n := len(nums)
	aux := make([]int, n)

	for exp := 1; maxNum/exp > 0; exp *= 10 {
		count := make([]int, 10)

		for i := 0; i < n; i++ {
			count[(nums[i]/exp)%10]++
		}

		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}

		for i := n - 1; i >= 0; i-- {
			idx := (nums[i] / exp) % 10
			aux[count[idx]-1] = nums[i]
			count[idx]--
		}

		for i := 0; i < n; i++ {
			nums[i] = aux[i]
		}
	}

	result := 0
	for i := 1; i < len(nums); i++ {
		result = max(result, nums[i]-nums[i-1])
	}
	return result
}

func Test164(t *testing.T) {
	assert.Equal(t, 9999999, maximumGap([]int{1, 10000000}))
}
