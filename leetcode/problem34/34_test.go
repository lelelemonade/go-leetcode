package problem34

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func searchRange(nums []int, target int) []int {

	searchBound := func(upper bool)int{
		left := 0
		right := len(nums) - 1
	
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] == target {
				if upper{
					if mid == len(nums)-1||nums[mid+1]!=target {
						return mid
					}else {
						left = mid + 1
					}
					continue
				}
				if mid == 0||nums[mid-1]!=target {
					return mid
				}else {
					right = mid - 1
				}
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		return -1
	}
	
	return []int{searchBound(false), searchBound(true)}
}

func Test34(t *testing.T){
	assert.Equal(t,[]int{3,4},searchRange([]int{5,7,7,8,8,10},8))
	assert.Equal(t,[]int{0,1},searchRange([]int{2,2},2))
}
