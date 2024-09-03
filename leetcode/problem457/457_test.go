package problem457

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func circularArrayLoop(nums []int) bool {
	moveOneStep := func(start int) int {
		return (nums[start] + start + 1001*len(nums)) % len(nums)
	}

	var checkCircular func(fast, slow int, positive bool) bool

	checkCircular = func(fast, slow int, positive bool) bool {
		if fast == slow {
			//circular
			return true
		}
		if nums[fast] > 0 != positive {
			nums[slow] = 0
			// not all positive or negative
			return false
		}
		defer func() {
			nums[fast] = 0
			nums[slow] = 0
		}()
		nextFast := moveOneStep(fast)
		if fast == nextFast||nums[nextFast] > 0 != positive {
			//self circular
			nums[slow] = 0
			return false
		}
		nextNextFast := moveOneStep(nextFast)

		nextSlow := moveOneStep(slow)

		return checkCircular(nextNextFast, nextSlow, positive)
	}

	for i := 0; i < len(nums); {
		fast:=moveOneStep(i)
		slow:=i
		if fast!=slow{
			if checkCircular(fast, slow, nums[i] > 0) {
				return true
			}
		}else{
			nums[i]=0
		}

		for i < len(nums) && nums[i] == 0 {
			i++
		}
	}

	return false
}

func Test457(t *testing.T) {
	assert.Equal(t,true, circularArrayLoop([]int{2,-1,1,2,2}))
	assert.Equal(t,false, circularArrayLoop([]int{-1,-2,-3,-4,-5,6}))
	assert.Equal(t,true, circularArrayLoop([]int{1,-1,5,1,4}))
	assert.Equal(t,false, circularArrayLoop([]int{-2,1,-1,-2,-2}))
}