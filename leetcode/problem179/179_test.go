package problem179

import (
	"slices"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func largestNumber(nums []int) string {
	sort.Slice(nums,func(a,b int)bool{
		aStr:=strconv.Itoa(nums[a])
		bStr:=strconv.Itoa(nums[b])

		return aStr+bStr>bStr+aStr
	})

	if nums[0]==0{
		return "0"
	}

	result:=""
	for _,num:=range nums{
		result+=strconv.Itoa(num)
	}
	return result
}

func Test179(t *testing.T){
	assert.Equal(t,"210",largestNumber([]int{10,2}))
	assert.Equal(t,"9534330",largestNumber([]int{3,30,34,5,9}))
}
