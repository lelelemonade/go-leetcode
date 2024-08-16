package problem128

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestConsecutive(nums []int) int {
    numSet := make(map[int]struct{})

	for _,v:=range nums {
		numSet[v]=struct{}{}
	}

	maxConsec:=0
	for _,v:=range nums {
		if _,e:=numSet[v-1];e{
			continue
		}
		consecCount:=1
		larger:=v+1
		_,e:=numSet[larger]
		for e{
			consecCount++
			larger++
			_,e=numSet[larger]
		}
		maxConsec=max(maxConsec,consecCount)
		
	}
	return maxConsec
}

func Test128(t *testing.T) {
	assert.Equal(t,4,longestConsecutive([]int{100,4,200,1,3,2}))
}