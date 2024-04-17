package problem519

import "math/rand/v2"

type Solution struct {
	revisedMap  map[int]int
	m, n, total int
}

func Constructor(m int, n int) Solution {
	return Solution{revisedMap: make(map[int]int), m: m, n: n, total: m * n}
}

func (this *Solution) Flip() []int {
	var x, y int
	ran := rand.IntN(this.total)
	this.total--
	if v, exist := this.revisedMap[ran]; exist {
		x, y = v/this.n, v%this.n
	} else {
		x, y = ran/this.n, ran%this.n
	}

	if v, exist := this.revisedMap[this.total]; exist {
		this.revisedMap[ran] = v
	} else {
		this.revisedMap[ran] = this.total
	}

	return []int{x, y}
}

func (this *Solution) Reset() {
	this.total = this.m * this.n
	this.revisedMap = make(map[int]int)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
