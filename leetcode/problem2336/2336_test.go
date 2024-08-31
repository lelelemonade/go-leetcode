package problem2336

import (
	"container/heap"
	"sort"
)

type IntHeap struct{ sort.IntSlice }

func (i *IntHeap) Push(a any) { i.IntSlice = append(i.IntSlice, a.(int)) }
func (i *IntHeap) Pop() any {
	old := i.IntSlice
	n := len(old)
	x := old[n-1]
	i.IntSlice = old[0 : n-1]
	return x
}

type SmallestInfiniteSet struct {
	infiniteStart int
	added         *IntHeap
}

func Constructor() SmallestInfiniteSet {
	h := &IntHeap{make([]int, 0)}
	heap.Init(h)
	return SmallestInfiniteSet{infiniteStart: 1, added: h}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.added.Len() > 0 && this.added.IntSlice[0] <= this.infiniteStart {
		return heap.Pop(this.added).(int)
	} else {
		this.infiniteStart++
		return this.infiniteStart - 1
	}
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.infiniteStart {
		return
	}
	for i := 0; i < this.added.Len(); i++ {
		if this.added.IntSlice[i] == num {
			return
		}
	}
	heap.Push(this.added, num)
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
