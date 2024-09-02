package problem2462

import (
	"container/heap"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MinHeap struct{ sort.IntSlice }

func (m *MinHeap) Push(a any) {
	(*m).IntSlice = append((*m).IntSlice, a.(int))
}

func (m *MinHeap) Pop() any {
	val := (*m).IntSlice[m.Len()-1]
	(*m).IntSlice = (*m).IntSlice[:m.Len()-1]
	return val
}

func totalCost(costs []int, k int, candidates int) int64 {
	if len(costs) == 1 {
		return int64(costs[0])
	}

	leftIdx, rightIdx := 0, len(costs)-1
	leftHeap, rightHeap := &MinHeap{}, &MinHeap{}
	heap.Init(leftHeap)
	heap.Init(rightHeap)

	for leftIdx < rightIdx && leftHeap.Len() < candidates {
		heap.Push(leftHeap, costs[leftIdx])
		heap.Push(rightHeap, costs[rightIdx])
		leftIdx++
		rightIdx--
	}

	result := 0

	for i := 0; i < k; i++ {
		if leftHeap.Len() == 0 {
			result += heap.Pop(rightHeap).(int)
			if leftIdx <= rightIdx {
				heap.Push(rightHeap, costs[rightIdx])
				rightIdx--
			}
			continue
		} else if rightHeap.Len() == 0 {
			result += heap.Pop(leftHeap).(int)
			if leftIdx <= rightIdx {
				heap.Push(leftHeap, costs[leftIdx])
				leftIdx++
			}
			continue
		}

		leftMin, rightMin := heap.Pop(leftHeap).(int), heap.Pop(rightHeap).(int)
		if leftMin < rightMin || (leftMin == rightMin && costs[leftIdx] <= costs[rightIdx]) {
			result += leftMin
			heap.Push(rightHeap, rightMin)
			if leftIdx <= rightIdx {
				heap.Push(leftHeap, costs[leftIdx])
				leftIdx++
			}
		} else {
			result += rightMin
			heap.Push(leftHeap, leftMin)
			if leftIdx <= rightIdx {
				heap.Push(rightHeap, costs[rightIdx])
				rightIdx--
			}
		}
	}

	return int64(result)
}

func Test2462(t *testing.T) {
	// assert.Equal(t,int64(11),totalCost([]int{17,12,10,2,7,2,11,20,8},3,4))
	// assert.Equal(t,int64(4),totalCost([]int{1,2,4,1},3,3))
	// assert.Equal(t,int64(393),totalCost([]int{57,33,26,76,14,67,24,90,72,37,30},11,2))
	// assert.Equal(t, int64(11), totalCost([]int{10, 1, 11, 10}, 2, 1))
	assert.Equal(t, int64(423), totalCost([]int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}, 11, 2))
}
