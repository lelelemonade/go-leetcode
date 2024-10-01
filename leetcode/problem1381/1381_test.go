package problem1381

type CustomStackNode struct {
	value     int
	increment int
}

type CustomStack struct {
	maxSize int
	nodes   []CustomStackNode
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{maxSize: maxSize, nodes: make([]CustomStackNode, 0)}
}

func (this *CustomStack) Push(x int) {
	if len(this.nodes) >= this.maxSize {
		return
	}
	this.nodes = append(this.nodes, CustomStackNode{value: x})
}

func (this *CustomStack) Pop() int {
	if len(this.nodes) == 0 {
		return -1
	}
	toPop := this.nodes[len(this.nodes)-1]
	this.nodes = this.nodes[:len(this.nodes)-1]
	if len(this.nodes) != 0 {
		this.nodes[len(this.nodes)-1].increment += toPop.increment
	}
	return toPop.value + toPop.increment
}

func (this *CustomStack) Increment(k int, val int) {
	if len(this.nodes) == 0 {
		return
	}
	if len(this.nodes) < k {
		this.nodes[len(this.nodes)-1].increment += val
	} else {
		this.nodes[k-1].increment += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
