package problem23

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap []*ListNode

func (h *MinHeap) Len()int{return len(*h)}

func (h *MinHeap) Less(a,b int) bool{
	return (*h)[a].Val<(*h)[b].Val
}

func (h *MinHeap) Swap(i,j int){
	(*h)[i],(*h)[j]=(*h)[j],(*h)[i]
}

func (h *MinHeap) Push(a any){
	*h=append(*h, a.(*ListNode))
}

func (h *MinHeap) Pop()any{
	value:=(*h)[h.Len()-1]
	*h=(*h)[:h.Len()-1]
	return value
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy:=&ListNode{}
	iter:=dummy
	h:=&MinHeap{}
	heap.Init(h)

	//init
	for _,v:=range lists{
		if v != nil {
			heap.Push(h,v)
		}
	}

	for h.Len()!=0{
		pop:=heap.Pop(h).(*ListNode)
		iter.Next=pop
		iter=iter.Next

		for h.Len()!=0 && pop.Next==nil {
			pop=heap.Pop(h).(*ListNode)
			iter.Next=pop
			iter=iter.Next
		}

		if pop!=nil&&pop.Next != nil {
			heap.Push(h,pop.Next)
		}
	}

	return dummy.Next
}
