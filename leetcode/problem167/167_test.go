package problem167

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(numbers []int, target int) []int {
    left,right:=0,len(numbers)-1

	for left<right{
		if target - numbers[left]< numbers[right]{
			right--
		} else if target - numbers[left]> numbers[right] {
			left++
		}else{
			return []int{left+1,right+1}
		}
	}

	return []int{-1,-1}
}

func Test167(t *testing.T) {
	assert.Equal(t,[]int{1,2},twoSum([]int{2,7,11,15},9))
	assert.Equal(t,[]int{1,3},twoSum([]int{2,3,4},6))
}