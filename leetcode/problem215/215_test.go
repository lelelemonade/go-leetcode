package problem215

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test215(t *testing.T) {
	//assert.Equal(t, 4, findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	//assert.Equal(t, 2, findKthLargest([]int{2, 1}, 1))
	//assert.Equal(t, 1, findKthLargest([]int{1}, 1))
	//assert.Equal(t, 5, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	assert.Equal(t, 6, findKthLargest([]int{7, 6, 5, 4, 3, 2, 1}, 2))
}

func findKthLargest(nums []int, k int) int {
	if len(nums) <= 1 {
		return nums[0]
	} else if len(nums) == 2 && k == 2 {
		return min(nums[0], nums[1])
	} else if len(nums) == 2 && k == 1 {
		return max(nums[0], nums[1])
	}
	pivot := nums[0]

	i, j := 1, len(nums)-1

	for {
		for i < j && nums[i] < pivot {
			i++
		}
		for i < j && nums[j] > pivot {
			j--
		}
		if i == j {
			if nums[i] > pivot {
				nums[i-1], nums[0] = nums[0], nums[i-1]
			} else {
				nums[i], nums[0] = nums[0], nums[i]
			}
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	if i-1 < len(nums)-k {
		return findKthLargest(nums[i:], k)
	} else if i-1 > len(nums)-k {
		return findKthLargest(nums[:i], k-i)
	} else {
		return nums[i-1]
	}
}
