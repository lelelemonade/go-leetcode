package problem373

import (
	"container/heap"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SumPair struct{
	x int
	y int
}

type MaxHeap []SumPair

func (m *MaxHeap) Len()int{
	return len(*m)
}

func (m *MaxHeap) Less(a,b int)bool{
	return (*m)[a].x+(*m)[a].y>(*m)[b].x+(*m)[b].y
}

func (m *MaxHeap) Swap(a,b int){
	(*m)[a],(*m)[b]=(*m)[b],(*m)[a]
}

func (m *MaxHeap) Push(a any){
	(*m)=append((*m), a.(SumPair))
}

func (m *MaxHeap) Pop()any{
	value :=(*m)[m.Len()-1]

	(*m)=(*m)[:m.Len()-1]

	return value
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	slices.Sort(nums1)
	slices.Sort(nums2)

	maxHeap:=&MaxHeap{}
	heap.Init(maxHeap)
    for _,v1:=range nums1{
		for _,v2:=range nums2{
			heap.Push(maxHeap,SumPair{v1,v2})
			if maxHeap.Len()>k{
				val:=heap.Pop(maxHeap).(SumPair)
				if val.x+val.y==v1+v2{
					break
				}
			}
		}
	}
	result:=make([][]int,0)
	for _,v:=range *maxHeap{
		result = append(result, []int{v.x,v.y})
	}
	return result
}

func Test373(t *testing.T) {
	assert.Equal(t,[][]int{},kSmallestPairs([]int{1,7,11},[]int{2,4,6},3))
}