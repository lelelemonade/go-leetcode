package problem238

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func productExceptSelf(nums []int) []int {
    total:=1
	zeroCount:=0

	for _,v :=range nums {
		if v == 0 {
			zeroCount++
		}
		total*=v
	}

	if zeroCount == 0 {
		for i:=0;i<len(nums);i++{
			nums[i]=total/nums[i]
		}
	}else if zeroCount == 1 {
		for i:=0;i<len(nums);i++{
			if nums[i] == 0 {
				nums[i]=1
				for j:=0;j<len(nums);j++{
					if i!=j {
						nums[i]*=nums[j]
						nums[j]=0
					}
				}
				break
			}
		}
	}else {
		for i:=0;i<len(nums);i++{
			nums[i]=0
		}
	}
	return nums
}

func Test238(t *testing.T) {
	// assert.Equal(t, []int{24,12,8,6},productExceptSelf([]int{1,2,3,4}))
	assert.Equal(t, []int{0,0,9,0,0},productExceptSelf([]int{-1,1,0,-3,3}))
}