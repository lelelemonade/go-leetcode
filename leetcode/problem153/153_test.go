package problem153

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findMin(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
    left:=0
	right:=len(nums)-1

	for left <= right {
		mid:=(left+right)/2

		if mid==0&&nums[mid+1]>nums[mid]{
			return nums[mid]
		}else if mid==len(nums)-1&&nums[mid-1]>nums[mid]{
			return nums[mid]
		}else if mid>0 && mid<len(nums)-1 && nums[mid+1]>nums[mid] && nums[mid-1]>nums[mid] {
			return nums[mid]
		}else if nums[mid]<nums[right] {
			right=mid-1
		}else {
			left=mid+1
		}
	}
	return -1
}

func Test153(t *testing.T) {
	assert.Equal(t,0,findMin([]int{4,5,6,7,0,1,2}))
}