package problem283

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func moveZeroes(nums []int)  {
	j:=0
	for i:=0;i<len(nums);i++{
		if nums[i]!=0{
			nums[i],nums[j]=nums[j],nums[i]
			j++
		}
	}
}

func Test283(t *testing.T) {
	testCase1:=[]int{0,1,0,3,12}
	moveZeroes(testCase1)
	assert.Equal(t,[]int{1,3,12,0,0},testCase1)
	testCase2:=[]int{1,0}
	moveZeroes(testCase2)
	assert.Equal(t,[]int{1,0},testCase2)
}