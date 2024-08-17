package problem57

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals)==0{
		return [][]int{newInterval}
	}
    for i:=0;i<len(intervals);{
		for i<len(intervals) && intervals[i][1]<newInterval[0]{
			i++
		}
		start:=i
		for i<len(intervals) && intervals[i][0]<=newInterval[1]{
			i++
		}
		end:=i
		if start==end {
			// non overlapping
			intervals=append(intervals,[]int{})
			copy(intervals[start+1:],intervals[start:len(intervals)-1])
			intervals[start]=newInterval
		}else{
			// overlapping
			intervals[start]=[]int{min(intervals[start][0],newInterval[0]),max(intervals[end-1][1],newInterval[1])}
			intervals=append(intervals[:start+1],intervals[end:]...)
		}

		break
	}

	return intervals
}

func Test57(t *testing.T) {
	assert.Equal(t,[][]int{{6,8}},insert([][]int{},[]int{6,8}))
	assert.Equal(t,[][]int{{1,5},{6,8}},insert([][]int{{1,5}},[]int{6,8}))
	// assert.Equal(t,[][]int{{1,5},{6,9}},insert([][]int{{1,3},{6,9}},[]int{2,5}))
}