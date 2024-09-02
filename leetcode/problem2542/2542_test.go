package problem2542

import (
	"container/heap"
	"slices"
	"sort"
)

// 定义一个最小堆
type MinHeap struct{  sort.IntSlice }

func (h *MinHeap) Push(x interface{}) {
	(*h).IntSlice = append((*h).IntSlice, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old.IntSlice)
	x := old.IntSlice[n-1]
	(*h).IntSlice = old.IntSlice[0 : n-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	idxArray:=make([]int,len(nums2))
	for i,_:=range idxArray{
		idxArray[i]=i
	}

	slices.SortFunc(idxArray,func(a,b int)int{
		return nums2[b]-nums2[a]
	})

	h:=&MinHeap{}
	heap.Init(h)

	result:=0
	sum1:=0

	for _,v:=range idxArray{
		sum1+=nums1[v]
		heap.Push(h,nums1[v])

		if h.Len()>k{
			sum1-=heap.Pop(h).(int)
		}

		if h.Len()==k{
			result=max(result,nums2[v]*sum1)
		}
	}
	
	return int64(result)
}
