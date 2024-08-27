package problem1493

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestSubarray(nums []int) int {
    start:=0
	lastZeroIdx:=-1
	oneCount:=0
	result:=0

	for i,v:=range nums{
		if v==1{
			oneCount++
		}else{
			if lastZeroIdx!=-1{
				oneCount-=lastZeroIdx-start
				start=lastZeroIdx+1
			}
			lastZeroIdx=i
		}
		result=max(result,oneCount)
	}

	if len(nums)!=result{
		return result
	}else{
		return result-1
	}
}

func Test1493(t *testing.T) {
	assert.Equal(t,3,longestSubarray([]int{1,1,0,1}))
	assert.Equal(t,5,longestSubarray([]int{0,1,1,1,0,1,1,0,1}))
}