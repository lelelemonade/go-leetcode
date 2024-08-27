package problem334

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func increasingTriplet(nums []int) bool {
	n:=len(nums)
	if n<3{
		return false
	}
	first,second:=0,math.MaxInt
	for i:=1;i<n;i++{
		if nums[i]>second{
			return true
		}else if nums[i]>first{
			second=nums[i]
		}else{
			first=nums[i]
		}
	}

	return false
}

func Test334(t *testing.T) {
	assert.Equal(t,true,increasingTriplet([]int{20,100,10,12,5,13}))
}