package problem502

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MaxHeap struct{
	profits []int
	capital []int
	idx []int
}

func (m *MaxHeap) Len()int{
	return len((*m).idx)
}

func (m *MaxHeap) Less(a,b int)bool{
	return (*m).profits[(*m).idx[a]]>(*m).profits[(*m).idx[b]]
}

func (m *MaxHeap) Swap(a,b int){
	(*m).idx[a],(*m).idx[b]=(*m).idx[b],(*m).idx[a]
}

func (m *MaxHeap) Push(a any){
	(*m).idx=append((*m).idx, a.(int))
}

func (m *MaxHeap) Pop()any{
	value:=(*m).idx[len((*m).idx)-1]
	(*m).idx=(*m).idx[:len((*m).idx)-1]
	return value
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
    maxHeap:=&MaxHeap{profits: profits,capital: capital,idx: make([]int, 0)}
	heap.Init(maxHeap)

	for i:=0;i<len(profits);i++{
		heap.Push(maxHeap,i)
	}

	for i:=0;i<k;i++{
		notChosen:=make([]int,0)

		if maxHeap.Len()==0{
			return w
		}
		maxProfitIdx:=heap.Pop(maxHeap).(int)

		for capital[maxProfitIdx]>w{

			notChosen = append(notChosen, maxProfitIdx)
			if maxHeap.Len()==0{
				return w
			}

			maxProfitIdx=heap.Pop(maxHeap).(int)
		}

		w+=profits[maxProfitIdx]

		for _,v:=range notChosen{
			heap.Push(maxHeap,v)
		}
	}
	return w
}

func Test503(t *testing.T) {
	assert.Equal(t,6,findMaximizedCapital(3,0,[]int{1,2,3},[]int{0,1,2}))
	assert.Equal(t,4,findMaximizedCapital(2,0,[]int{1,2,3},[]int{0,1,1}))
}