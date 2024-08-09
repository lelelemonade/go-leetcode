package problem146

import (
	"testing"
)

func Test146(t *testing.T) {
	lruCache := Constructor(2)

	lruCache.Put(2, 1)
	lruCache.Put(2, 2)
	lruCache.Get(2)
	lruCache.Put(1, 1)
	lruCache.Put(4, 1)
	lruCache.Get(2)

	//lruCache.Put(1, 0)
	//lruCache.Put(2, 2)
	//lruCache.Get(1)
	//lruCache.Put(3, 3)
	//lruCache.Get(2)
	//lruCache.Put(4, 4)
	//lruCache.Get(1)
	//lruCache.Get(3)
	//lruCache.Get(4)
}

type LRUCache struct {
	capacity  int
	dummyHead *LRUCell
	dummyTail *LRUCell
	container map[int]*LRUCell
}

type LRUCell struct {
	key   int
	value int
	last  *LRUCell
	next  *LRUCell
}

func Constructor(capacity int) LRUCache {
	container := make(map[int]*LRUCell, capacity)
	dummyHead := LRUCell{}
	dummyTail := LRUCell{}
	dummyTail.last = &dummyHead
	dummyHead.next = &dummyTail
	return LRUCache{
		capacity:  capacity,
		dummyHead: &dummyHead,
		dummyTail: &dummyTail,
		container: container,
	}
}

func (this *LRUCache) Get(key int) int {
	result := this.container[key]

	if nil == result {
		return -1
	}

	result.last.next = result.next
	result.next.last = result.last

	result.next = this.dummyHead.next
	this.dummyHead.next.last = result
	this.dummyHead.next = result
	result.last = this.dummyHead

	return result.value
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		this.container[key].value = value
		return
	}

	if len(this.container)+1 > this.capacity {
		beforeDummy := this.dummyTail.last
		beforeDummy.last.next = this.dummyTail
		this.dummyTail.last = beforeDummy.last
		delete(this.container, beforeDummy.key)
		beforeDummy.next = nil
		beforeDummy.last = nil
	}

	cellToPut := LRUCell{
		key:   key,
		value: value,
		last:  this.dummyHead,
		next:  this.dummyHead.next,
	}

	this.dummyHead.next.last = &cellToPut
	this.dummyHead.next = &cellToPut

	this.container[key] = &cellToPut
}
