package problem238

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func productExceptSelf(nums []int) []int {
    answer:=make([]int,len(nums))
	answer[0]=1

	leftProduct:=1
	for i:=0;i<len(nums)-1;i++ {
		leftProduct*=nums[i]
		answer[i+1]=leftProduct
	}

	rightProduct:=1
	for i:=len(nums)-1;i>=0;i-- {
		answer[i]=answer[i]*rightProduct
		rightProduct*=nums[i]
	}

	return answer
}

func Test238(t *testing.T) {
	assert.Equal(t, []int{24,12,8,6},productExceptSelf([]int{1,2,3,4}))
	assert.Equal(t, []int{0,0,9,0,0},productExceptSelf([]int{-1,1,0,-3,3}))
}