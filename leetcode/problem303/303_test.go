package problem303

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type NumArray struct {
    prefixSum []int
}


func Constructor(nums []int) NumArray {
    prefixSum:=make([]int,len(nums)+1)

	for i:=0;i<len(nums);i++{
		prefixSum[i+1]=prefixSum[i]+nums[i]
	}
	return NumArray{prefixSum: prefixSum}
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.prefixSum[right+1]-this.prefixSum[left]
}


func Test303(t *testing.T) {
	numArray:=Constructor([]int{-2,0,3,-5,2,-1})
	assert.Equal(t,1,numArray.SumRange(0,2))
	assert.Equal(t,-1,numArray.SumRange(2,5))
	assert.Equal(t,-3,numArray.SumRange(0,5))
}