package problem435

import (
	"fmt"
	"sort"
	"testing"
)

func Test435(t *testing.T) {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	//intervals := [][]int{{1, 2}, {1, 2}, {1, 2}}
	//intervals := [][]int{{1, 2}, {2, 3}}
	//intervals := [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}

	fmt.Println(eraseOverlapIntervals(intervals))
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0

	for i := 0; i < len(intervals); {
		j := 1
		for i+j < len(intervals) && intervals[i][1] > intervals[i+j][0] {
			count++
			j++
		}
		i += j
	}
	return count
}
