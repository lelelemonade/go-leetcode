package problem2598

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findSmallestInteger(nums []int, value int) int {
	mex:=0
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[(x%value+value)%value]++
	}
	for cnt[mex%value] > 0 {
		cnt[mex%value]--
		mex++
	}
	return mex
}

func Test2598(t *testing.T) {
	assert.Equal(t,4,findSmallestInteger([]int{1,-10,7,13,6,8},5))
}