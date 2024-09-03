package problem295

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MinHeap []int

func (m *MinHeap) Len() int{return len(*m)}
func (m *MinHeap) Less(a,b int) bool{return (*m)[a]<(*m)[b]}
func (m *MinHeap) Swap(a,b int) {(*m)[a],(*m)[b]=(*m)[b],(*m)[a]}
func (m *MinHeap) Push(a any) {
	*m=append(*m,a.(int))
}
func (m *MinHeap) Pop() any{
	value := (*m)[m.Len()-1]
	*m=(*m)[:m.Len()-1]
	return value
}

type MaxHeap []int
func (m *MaxHeap) Len() int{return len(*m)}
func (m *MaxHeap) Less(a,b int) bool{return (*m)[a]>(*m)[b]}
func (m *MaxHeap) Swap(a,b int) {(*m)[a],(*m)[b]=(*m)[b],(*m)[a]}
func (m *MaxHeap) Push(a any) {
	*m=append(*m,a.(int))
}
func (m *MaxHeap) Pop() any{
	value := (*m)[m.Len()-1]
	*m=(*m)[:m.Len()-1]
	return value
}

type MedianFinder struct {
    maxHeap *MaxHeap
    minHeap *MinHeap
}


func Constructor() MedianFinder {
    minHeap:=&MinHeap{}
    maxHeap:=&MaxHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)
	return MedianFinder{minHeap: minHeap,maxHeap: maxHeap}
}


func (this *MedianFinder) AddNum(num int)  {
	if this.minHeap.Len()==0 && this.maxHeap.Len()==0{
		heap.Push(this.maxHeap,num)
		return
	}
    if this.minHeap.Len()==this.maxHeap.Len(){
		minOfMinHeap:=heap.Pop(this.minHeap).(int)
		if num <= minOfMinHeap{
			heap.Push(this.maxHeap,num)
		}else{
			heap.Push(this.minHeap,num)
		}

		heap.Push(this.minHeap,minOfMinHeap)
	}else if this.minHeap.Len()>this.maxHeap.Len(){
		minOfMinHeap:=heap.Pop(this.minHeap).(int)
		if num <= minOfMinHeap{
			heap.Push(this.maxHeap,num)
			heap.Push(this.minHeap,minOfMinHeap)
		}else{
			heap.Push(this.minHeap,num)
			heap.Push(this.minHeap,minOfMinHeap)
			heap.Push(this.maxHeap,heap.Pop(this.minHeap))
		}
	}else if this.minHeap.Len()<this.maxHeap.Len(){
		maxOfMaxHeap:=heap.Pop(this.maxHeap).(int)
		if num >= maxOfMaxHeap{
			heap.Push(this.minHeap,num)
			heap.Push(this.maxHeap,maxOfMaxHeap)
		}else{
			heap.Push(this.maxHeap,num)
			heap.Push(this.maxHeap,maxOfMaxHeap)
			heap.Push(this.minHeap,heap.Pop(this.maxHeap))
		}

	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len()==0 && this.maxHeap.Len()==0{
		return 0
	}

    if this.minHeap.Len()==this.maxHeap.Len(){
		maxOfMaxHeap:=heap.Pop(this.maxHeap).(int)
		minOfMinHeap:=heap.Pop(this.minHeap).(int)

		result := (float64(maxOfMaxHeap)+float64(minOfMinHeap))/float64(2)

		heap.Push(this.maxHeap,maxOfMaxHeap)
		heap.Push(this.minHeap,minOfMinHeap)
		return result
	}else if this.minHeap.Len()>this.maxHeap.Len(){
		minOfMinHeap:=heap.Pop(this.minHeap).(int)

		result := float64(minOfMinHeap)

		heap.Push(this.minHeap,minOfMinHeap)

		return result
	}else{
		maxOfMaxHeap:=heap.Pop(this.maxHeap).(int)

		result := float64(maxOfMaxHeap)

		heap.Push(this.maxHeap,maxOfMaxHeap)

		return result
	}
}

func Test295(t *testing.T) {
	// medianFinder:=Constructor()
	// medianFinder.AddNum(1)
	// medianFinder.AddNum(2)
	// assert.Equal(t,1.5,medianFinder.FindMedian())
	// medianFinder.AddNum(3)
	// assert.Equal(t,2.0,medianFinder.FindMedian())

	medianFinder:=Constructor()
	medianFinder.AddNum(-1)
	assert.Equal(t,-1.0,medianFinder.FindMedian())
	medianFinder.AddNum(-2)
	assert.Equal(t,-1.5,medianFinder.FindMedian())
	medianFinder.AddNum(-3)
	assert.Equal(t,-2.0,medianFinder.FindMedian())
	medianFinder.AddNum(-4)
	assert.Equal(t,-2.5,medianFinder.FindMedian())
}