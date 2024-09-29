package problem1283

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func smallestDivisor(nums []int, threshold int) int {
	left, right := 1, slices.Max(nums)
	sumDivident := func(divisor int) (result int) {
		for _, v := range nums {
			result += v/divisor
			if v%divisor>0{
				result+=1
			}
		}
		return
	}
	for left < right {
		mid := left + (right-left)/2
		midSum := sumDivident(mid)
		if midSum > threshold {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func Test1283(t *testing.T){
	assert.Equal(t,5,smallestDivisor([]int{1,2,5,9},6))
}
