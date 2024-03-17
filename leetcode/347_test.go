package leetcode

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test347(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
}

func topKFrequent(nums []int, k int) []int {
	myTopKHeap := MyTopKHeap{
		values: []int{},
		k:      k,
	}

	heap.Init(&myTopKHeap)

	for _, num := range nums {
		heap.Push(&myTopKHeap, num)
	}

	var result []int
	for i := 0; i < myTopKHeap.Len(); i++ {
		result = append(result, myTopKHeap.Pop().(int))
	}

	return result
}

type MyTopKHeap struct {
	values []int
	k      int
}

func (myTopKHeap MyTopKHeap) Len() int {
	return len(myTopKHeap.values)
}

func (myTopKHeap MyTopKHeap) Less(i, j int) bool {
	return myTopKHeap.values[i] < myTopKHeap.values[j]
}

func (myTopKHeap MyTopKHeap) Swap(i, j int) {
	myTopKHeap.values[i], myTopKHeap.values[j] = myTopKHeap.values[j], myTopKHeap.values[i]
}

func (myTopKHeap *MyTopKHeap) Push(i interface{}) {
	if myTopKHeap.Len() > myTopKHeap.k {
		myTopKHeap.Pop()
	}

	myTopKHeap.values = append(myTopKHeap.values, i.(int))
}

func (myTopKHeap *MyTopKHeap) Pop() interface{} {
	old := *myTopKHeap
	n := len(old.values)
	x := old.values[n-1]
	myTopKHeap.values = old.values[0 : n-1]
	return x
}
