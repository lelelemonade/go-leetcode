package problem2382

import (
	"fmt"
	"testing"
)

func Test2383(t *testing.T) {
	//fmt.Println(maximumSegmentSum([]int{1, 2, 5, 6, 1}, []int{0, 3, 2, 4, 1}))
	fmt.Println(maximumSegmentSum([]int{3, 2, 11, 1}, []int{3, 2, 1, 0}))
}

type UnionFindSet struct {
	parent        []int
	rank          []int
	segmentSum    map[int]int64
	MaxSegmentSum int64
}

func NewUnionFindSet(n int) UnionFindSet {
	parent := make([]int, n)
	rank := make([]int, n)
	segmentSum := make(map[int]int64)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return UnionFindSet{parent, rank, segmentSum, 0}
}

func (ufs *UnionFindSet) find(x int) int {
	if x != ufs.parent[x] {
		ufs.parent[x] = ufs.find(ufs.parent[x])
	}
	if ufs.segmentSum[x] != ufs.segmentSum[ufs.parent[x]] {
		ufs.segmentSum[x] = ufs.segmentSum[ufs.parent[x]]
	}
	return ufs.parent[x]
}

func (ufs *UnionFindSet) union(x, y int) {
	rootx, rooty := ufs.find(x), ufs.find(y)
	if rootx != rooty {
		if ufs.rank[rootx] < ufs.rank[rooty] {
			ufs.parent[rootx] = rooty
			ufs.segmentSum[rooty] += ufs.segmentSum[rootx]
			if ufs.segmentSum[rooty] > ufs.MaxSegmentSum {
				ufs.MaxSegmentSum = ufs.segmentSum[rooty]
			}
			delete(ufs.segmentSum, rootx)
		} else if ufs.rank[rootx] > ufs.rank[rooty] {
			ufs.parent[rooty] = rootx
			ufs.segmentSum[rootx] += ufs.segmentSum[rooty]
			if ufs.segmentSum[rootx] > ufs.MaxSegmentSum {
				ufs.MaxSegmentSum = ufs.segmentSum[rootx]
			}
			delete(ufs.segmentSum, rooty)
		} else {
			ufs.parent[rootx] = rooty
			ufs.rank[rooty]++
			ufs.segmentSum[rooty] += ufs.segmentSum[rootx]
			if ufs.segmentSum[rooty] > ufs.MaxSegmentSum {
				ufs.MaxSegmentSum = ufs.segmentSum[rooty]
			}
			delete(ufs.segmentSum, rootx)
		}
	}
}

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	result := make([]int64, len(nums))
	ufs := NewUnionFindSet(len(removeQueries))
	processed := make([]bool, len(removeQueries))

	for i := len(removeQueries) - 1; i > 0; i-- {
		idxToAdd := removeQueries[i]
		ufs.segmentSum[idxToAdd] = int64(nums[idxToAdd])
		if ufs.segmentSum[idxToAdd] > ufs.MaxSegmentSum {
			ufs.MaxSegmentSum = ufs.segmentSum[idxToAdd]
		}
		if idxToAdd < len(removeQueries)-1 && processed[idxToAdd+1] {
			ufs.union(idxToAdd, idxToAdd+1)
		}
		if idxToAdd > 0 && processed[idxToAdd-1] {
			ufs.union(idxToAdd, idxToAdd-1)
		}
		processed[idxToAdd] = true
		result[i-1] = ufs.MaxSegmentSum
	}

	return result
}
