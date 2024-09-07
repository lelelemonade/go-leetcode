package problem565

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func arrayNesting(nums []int) int {
    visited:=make(map[int]struct{})
	maxSetSize:=0

	for i:=0;i<len(nums);i++{
		if _,e:=visited[nums[i]];e{
			continue
		}
		set:=make(map[int]int)
		j:=i
		for {
			set[j]=nums[j]
			visited[j]=struct{}{}
			j=nums[j]
			if _,e:=set[j];e{
				break
			}
		}
		maxSetSize=max(maxSetSize,len(set))
	}

	return maxSetSize
}

func Test565(t *testing.T){
	assert.Equal(t,4,arrayNesting([]int{5,4,0,3,1,6,2}))
	assert.Equal(t,1,arrayNesting([]int{0,1,2}))
}
