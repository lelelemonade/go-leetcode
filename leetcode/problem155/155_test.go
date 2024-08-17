package problem155

type MinStack struct {
	stack []int
	min_stack []int
}


func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min_stack: make([]int, 0),
	}
}


func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack, val)
	minValue := 0
	if len(this.min_stack)==0{
		minValue=val
	}else{
		minValue=min(this.min_stack[len(this.min_stack)-1],val)
	}
	this.min_stack = append(this.min_stack, minValue)
}


func (this *MinStack) Pop()  {
	this.stack=this.stack[:len(this.stack)-1]
	this.min_stack=this.min_stack[:len(this.min_stack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.min_stack[len(this.min_stack)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */