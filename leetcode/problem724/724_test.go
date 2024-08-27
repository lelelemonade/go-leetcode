package problem724

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func pivotIndex(nums []int) int {
	if len(nums)==1{
		return 0
	}
	leftSum, rightSum := 0, 0
	for i:=0;i<len(nums);i++{
		rightSum+=nums[i]
	}
	for i:=0;i<len(nums);i++{
		if i==0{
			leftSum+=0
		}else{
			leftSum+=nums[i-1]
		}
		rightSum-=nums[i]
		if leftSum==rightSum{
			return i
		}
	}
	return -1
}

func Test724(t *testing.T) {
	assert.Equal(t, 0, pivotIndex([]int{2, 1, -1}))
	assert.Equal(t, -1, pivotIndex([]int{-1, -1, -1, -1, -1, -1}))
	assert.Equal(t, 2, pivotIndex([]int{-1, -1, -1, -1, -1, 0}))
	assert.Equal(t, 2, pivotIndex([]int{-1, -1, 0, -1, -1, 0}))
	assert.Equal(t, 3, pivotIndex([]int{-1, -1, 0, -1, -1, -1}))
}
