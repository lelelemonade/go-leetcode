package _3_02

import "math"

type MinStack struct {
	body []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{min: []int{math.MaxInt64}, body: []int{}}
}

func (this *MinStack) Push(x int) {
	if len(this.body) == 0 {
		this.min = append(this.min, x)
		this.body = append(this.body, x)
		return
	}
	if x < this.min[len(this.min)-1] {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1])
	}
	this.body = append(this.body, x)
}

func (this *MinStack) Pop() {
	this.body = this.body[:len(this.body)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.body[len(this.body)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
