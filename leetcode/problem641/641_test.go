package problem641

type MyDequeElement struct {
	next  *MyDequeElement
	prev  *MyDequeElement
	value int
}

type MyCircularDeque struct {
	head *MyDequeElement
	tail *MyDequeElement
	cap  int
	len  int
}

func Constructor(k int) MyCircularDeque {
	head := &MyDequeElement{}
	tail := &MyDequeElement{}
	head.next = tail
	tail.prev = head
	return MyCircularDeque{
		head: head,
		tail: tail,
		cap:  k,
		len:  0,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.len == this.cap {
		return false
	}
	newElement := &MyDequeElement{value: value}
	newElement.next = this.head.next
	this.head.next.prev = newElement
	this.head.next = newElement
	newElement.prev = this.head
	this.len++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.len == this.cap {
		return false
	}
	newElement := &MyDequeElement{value: value}
	newElement.next = this.tail
	newElement.prev = this.tail.prev
	this.tail.prev.next = newElement
	this.tail.prev = newElement

	this.len++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.len == 0 {
		return false
	}

	this.head.next = this.head.next.next
	this.head.next.prev = this.head

	this.len--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.len == 0 {
		return false
	}

	this.tail.prev = this.tail.prev.prev
	this.tail.prev.next = this.tail

	this.len--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.len == 0 {
		return -1
	}
	return this.head.next.value
}

func (this *MyCircularDeque) GetRear() int {
	if this.len == 0 {
		return -1
	}
	return this.tail.prev.value
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.len == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.len == this.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
