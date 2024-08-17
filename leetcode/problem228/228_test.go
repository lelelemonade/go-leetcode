package problem228

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	results:= make([]string,0)

    for i:=0;i<len(nums);{
		start:=nums[i]
		for i<len(nums)-1&&nums[i]+1==nums[i+1]{
			i++
		}

		if nums[i]==start{
			results = append(results, strconv.Itoa(nums[i]))
		}else{
			results = append(results, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i]))
		}
		i++
	}
	return results
}

func Test228(t *testing.T) {
	assert.Equal(t,[]string{"0->2","4->5","7"},summaryRanges([]int{0,1,2,4,5,7}))
	assert.Equal(t,[]string{"0","2->4","6","8->9"},summaryRanges([]int{0,2,3,4,6,8,9}))
}