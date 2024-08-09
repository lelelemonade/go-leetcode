package problem45

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func jump(nums []int) int {
	dpJump := make([]int, len(nums))
	dpJump[0] = 0

	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			if dpJump[i+j] == 0 {
				dpJump[i+j]=dpJump[i]+1
			}else {
				dpJump[i+j] = min(dpJump[i+j], dpJump[i]+1)
			}
		}
	}
	return dpJump[len(dpJump)-1]
}

func Test45(t *testing.T) {
	testCase1 := []int{2,3,1,1,4}

	assert.Equal(t,2,jump(testCase1))
}
