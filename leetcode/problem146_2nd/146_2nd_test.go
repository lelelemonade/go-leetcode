package problem146_2nd

import "testing"

type LRUCache struct {
	capacity    int
	size        int
	lruMap      map[int]*LRUNode
	lruListHead *LRUNode
	lruListTail *LRUNode
}

type LRUNode struct {
	key  int
	val  int
	last *LRUNode
	next *LRUNode
}

func Constructor(capacity int) LRUCache {
	lruListHead := &LRUNode{}
	lruListTail := &LRUNode{}
	lruListHead.next = lruListTail
	lruListTail.last = lruListHead
	return LRUCache{capacity: capacity,
		lruListHead: lruListHead,
		lruListTail: lruListTail,
		lruMap:      make(map[int]*LRUNode),
		size:        0,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, e := this.lruMap[key]; e {
		v.last.next = v.next
		v.next.last = v.last

		this.lruListHead.next.last = v
		v.next = this.lruListHead.next
		this.lruListHead.next = v
		v.last = this.lruListHead

		return v.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		this.lruListHead.next.val = value
	} else {
		if this.capacity == this.size {
			// evict
			evictedkey := this.lruListTail.last.key
			this.lruListTail.last = this.lruListTail.last.last
			this.lruListTail.last.next = this.lruListTail
			this.size--
			delete(this.lruMap, evictedkey)
		}
		newNode := &LRUNode{key: key, val: value}

		newNode.next = this.lruListHead.next
		newNode.last = this.lruListHead
		this.lruListHead.next.last = newNode
		this.lruListHead.next = newNode

		this.lruMap[key] = newNode
		this.size++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func Test146(t *testing.T) {
	lruCache := Constructor(2)

	// lruCache.Put(2, 1)
	// lruCache.Put(2, 2)
	// lruCache.Get(2)
	// lruCache.Put(1, 1)
	// lruCache.Put(4, 1)
	// lruCache.Get(2)

	lruCache.Put(1, 0)
	lruCache.Put(2, 2)
	lruCache.Get(1)
	lruCache.Put(3, 3)
	lruCache.Get(2)
	lruCache.Put(4, 4)
	lruCache.Get(1)
	lruCache.Get(3)
	lruCache.Get(4)
}
