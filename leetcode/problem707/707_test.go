package problem707

import (
	"fmt"
	"testing"
)

type MyLinkedNode struct {
	val  int
	prev *MyLinkedNode
	next *MyLinkedNode
}

type MyLinkedList struct {
	head *MyLinkedNode
	tail *MyLinkedNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: &MyLinkedNode{
			val:  -1,
			prev: nil,
			next: nil,
		},
		tail: &MyLinkedNode{
			val:  -1,
			prev: nil,
			next: nil,
		},
	}
}

func (this *MyLinkedList) Get(index int) int {
	iter := this.head
	i := -1
	for iter != nil && i != index {
		i++
		iter = iter.next
	}
	if iter != nil {
		return iter.val
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	myLinkedNode := &MyLinkedNode{
		val:  val,
		prev: nil,
		next: nil,
	}

	if this.head.next == nil {
		this.head.next = myLinkedNode
		myLinkedNode.prev = this.head
		this.tail.prev = myLinkedNode
		myLinkedNode.next = this.tail
	} else {
		myLinkedNode.next = this.head.next
		myLinkedNode.prev = this.head
		this.head.next.prev = myLinkedNode
		this.head.next = myLinkedNode
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	myLinkedNode := &MyLinkedNode{
		val:  val,
		prev: nil,
		next: nil,
	}

	if this.tail.prev == nil {
		this.head.next = myLinkedNode
		myLinkedNode.prev = this.head
		this.tail.prev = myLinkedNode
		myLinkedNode.next = this.tail
	} else {
		myLinkedNode.next = this.tail
		myLinkedNode.prev = this.tail.prev
		this.tail.prev.next = myLinkedNode
		this.tail.prev = myLinkedNode
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.head.next == nil && index == 0 {
		myLinkedNode := &MyLinkedNode{
			val:  val,
			prev: nil,
			next: nil,
		}

		this.head.next = myLinkedNode
		myLinkedNode.prev = this.head
		this.tail.prev = myLinkedNode
		myLinkedNode.next = this.tail
		return
	}

	iter := this.head
	i := -1

	for iter != nil && i != index {
		i++
		iter = iter.next
	}

	if iter != nil {
		myLinkedNode := &MyLinkedNode{
			val:  val,
			prev: nil,
			next: nil,
		}

		myLinkedNode.prev = iter.prev
		myLinkedNode.next = iter
		iter.prev.next = myLinkedNode
		iter.prev = myLinkedNode
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	iter := this.head
	i := -1

	for iter != nil && iter.next != nil && i != index {
		i++
		iter = iter.next
	}

	if iter != nil && iter.next != nil {
		iter.prev.next = iter.next
		iter.next.prev = iter.prev
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func Test707(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtIndex(0, 10)
	myLinkedList.AddAtIndex(0, 20)
	myLinkedList.AddAtIndex(1, 30)

	fmt.Println(myLinkedList.Get(0))
}
