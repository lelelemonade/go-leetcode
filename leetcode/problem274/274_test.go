package problem274

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func hIndex(citations []int) int {
	slices.Sort(citations)
	citationCount:=len(citations)
	h := 0

	for i:=citationCount-1;i>=0;i--{
		if citations[i]>=citationCount-i{
			h++
		}
	}

	return h
}

func Test274(t *testing.T) {
	assert.Equal(t,3,hIndex([]int{3,0,6,1,5}))
	assert.Equal(t,1,hIndex([]int{1,3,1}))
}