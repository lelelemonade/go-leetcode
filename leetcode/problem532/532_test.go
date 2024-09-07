package problem532

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

type pair struct{
	x int
	y int
}

func findPairs(nums []int, k int) int {
	slices.Sort(nums)

	accepted:=make(map[int]int)
	result:=make(map[pair]struct{})
	for _,v:=range nums{
		if accepted[v]>0{
			accepted[v]--
			result[pair{x:v-k,y:v}]=struct{}{}
		}
		accepted[v+k]++
	}
	return len(result)
}

func Test532(t *testing.T) {
	assert.Equal(t,1,findPairs([]int{1,1,1,1,1},0))
}